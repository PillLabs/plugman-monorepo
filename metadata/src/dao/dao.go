package dao

import (
	"errors"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/database"
	"github.com/PillLabs/plugman-monorepo/metadata/src/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sort"
	"strconv"
	"time"
)

type PlugmanDao struct {
	DB *gorm.DB
}

func NewPlugmanDao() *PlugmanDao {
	return &PlugmanDao{
		DB: database.DB,
	}
}

func (d *PlugmanDao) GetNonce(address string, nonce *model.Nonce) error {
	return d.DB.Model(model.Nonce{}).Where("address = ?", address).First(nonce).Error
}

func (d *PlugmanDao) AllocateNonce(address string, mintType, count uint, nonce *model.Nonce) error {
	tx := d.DB.Begin()
	if tx.Error != nil {
		// transaction not begin
		return tx.Error
	}
	commit := false
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error("[Dao - AllocateNonce] Recovered in AllocateNonce: %v", r)
		}
		if !commit {
			tx.Rollback()
		}
	}()

	err := tx.Set("gorm:query_option", "FOR UPDATE").Model(model.Nonce{}).
		Where("address = ?", address).First(nonce).Error
	if err == gorm.ErrRecordNotFound {
		if mintType == common.MINT_TYPE_PUBLIC_SALE {
			nonce = new(model.Nonce)
			nonce.Address = address
			nonce.PublicNonce = 1
			err = tx.Create(nonce).Error
			if err != nil {
				return err
			}
			commit = true
		} else {
			return common.ErrInsufficientPresaleQuota
		}
	} else if err == nil {
		switch mintType {
		case common.MINT_TYPE_OG:
			if nonce.OGCount-nonce.OGMinted < count {
				return common.ErrInsufficientPresaleQuota
			}
			nonce.OGMinted = nonce.OGMinted + count
			nonce.OGNonce = nonce.OGNonce + 1
		case common.MINT_TYPE_WL:
			if nonce.WLCount-nonce.WLMinted < count {
				return common.ErrInsufficientPresaleQuota
			}
			nonce.WLMinted = nonce.WLMinted + count
			nonce.WLNonce = nonce.WLNonce + 1
		case common.MINT_TYPE_PUBLIC_SALE:
			nonce.PublicNonce = nonce.PublicNonce + 1
		default:
			return common.ErrInvalidMintType
		}
		err = tx.Save(nonce).Error
		if err != nil {
			return err
		}
		commit = true
	}
	if commit {
		return tx.Commit().Error
	}
	return err
}

func (d *PlugmanDao) AllocateMetadata(address string, mintType, count uint, nonce uint) (*[]*model.Metadata, error) {
	var (
		oneOfOneMeta   []*model.Metadata
		sequentialMeta []*model.Metadata
		unusedMeta     []*model.Metadata
		finalMeta      []*model.Metadata
	)

	tx := d.DB.Begin().Set("gorm:query_option", "FOR UPDATE").Session(&gorm.Session{})
	if tx.Error != nil {
		// transaction not begin
		return nil, tx.Error
	}
	commit := false
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error("[Dao - AllocateMetadata] Recovered in AllocateMetadata: %v", r)
		}
		if !commit {
			tx.Rollback()
		}
	}()

	now := uint64(time.Now().Unix())

	// 1. get allocated but un-minted one-of-ones metadata
	// 2. if remnant > 0, get un-allocated metadata
	// 3. if remnant > 0, get allocated un-minted generic metadata
	// 4. if remnant > 0, return common.ErrNoUnusedMetadata
	err := tx.Model(&model.Metadata{}).Not("address = ?", "").
		Where("one_of_one = ? AND tx_hash = ?", true, "").Order("id").Find(&oneOfOneMeta).Error
	if err != nil {
		return nil, err
	}
	for _, meta := range oneOfOneMeta {
		if now >= meta.Timestamp {
			finalMeta = append(finalMeta, meta)
		}
	}
	remnant := int(count) - len(finalMeta)
	if remnant < 0 {
		finalMeta = finalMeta[:count]
	} else if remnant > 0 {
		err = tx.Model(&model.Metadata{}).Where("address = ?", "").
			Order("id").Limit(remnant).Find(&sequentialMeta).Error
		if err != nil {
			return nil, err
		}
		remnant = remnant - len(sequentialMeta)
		if remnant > 0 {
			var m []*model.Metadata
			err = tx.Model(&model.Metadata{}).Not("address = ?", "").
				Where("one_of_one = ? AND tx_hash = ?", false, "").
				Order("id").Limit(remnant).Find(&m).Error
			if err != nil {
				return nil, err
			}
			for _, meta := range m {
				if now >= meta.Timestamp {
					unusedMeta = append(unusedMeta, meta)
				}
			}
			remnant = remnant - len(unusedMeta)
			if remnant > 0 {
				return nil, common.ErrNoUnusedMetadata
			}
			finalMeta = append(finalMeta, unusedMeta...)
		}
		finalMeta = append(finalMeta, sequentialMeta...)
	}
	sort.Sort(model.SortMetadata(finalMeta))
	for _, meta := range finalMeta {
		meta.Address = address
		meta.MintType = mintType
		meta.Nonce = nonce
		meta.Timestamp = now + 1800
	}
	err = tx.Save(finalMeta).Error
	if err != nil {
		return nil, err
	}
	commit = true
	return &finalMeta, tx.Commit().Error
}

func (d *PlugmanDao) GetBlockByChain(chainName string) (*model.Block, error) {
	var block model.Block
	err := d.DB.Model(model.Block{}).Where("chain_name = ?", chainName).First(&block).Error
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func (d *PlugmanDao) SetBlockByChain(block *model.Block) error {
	return d.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "chain_name"}},
		DoUpdates: clause.Assignments(
			map[string]interface{}{"number": block.Number}),
	}).Create(block).Error
}

func (d *PlugmanDao) UpdateMetadataTx(address, txHash string, nonce, firstTokenId, mintType uint) error {
	var (
		err          error
		metadataList []*model.Metadata
	)

	tx := d.DB.Begin().Set("gorm:query_option", "FOR UPDATE")
	if tx.Error != nil {
		// transaction not begin
		return tx.Error
	}
	commit := false
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error("[Dao - UpdateMetadataTx] Recovered in UpdateMetadataTx: %v", r)
		}
		if !commit {
			tx.Rollback()
		}
	}()

	err = tx.Model(model.Metadata{}).
		Where("address = ? AND nonce = ? AND mint_type = ? AND tx_hash = ?", address, nonce, mintType, "").
		Order("id").Find(&metadataList).Error
	if err != nil {
		return err
	}
	if len(metadataList) == 0 {
		return errors.New("not found such metadata")
	}
	for i, m := range metadataList {
		m.TxHash = txHash
		tokenId := strconv.Itoa(int(firstTokenId) + i)
		m.TokenId = tokenId
		if err = tx.Save(m).Error; err != nil {
			return err
		}
	}
	commit = true
	return tx.Commit().Error
}

func (d *PlugmanDao) SetWhitelist(addressWLs *[]model.AddressWL) error {
	tx := d.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	commit := false
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error("[Dao - SetWhitelist] Recovered in SetWhitelist: %v", r)
		}
		if !commit {
			tx.Rollback()
		}
	}()

	for _, wl := range *addressWLs {
		nonce := model.Nonce{
			Address: wl.Address,
			OGCount: wl.OGCount,
			WLCount: wl.WLCount,
		}
		err := tx.Model(model.Nonce{}).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&nonce).Error
		if err != nil {
			return err
		}
	}
	commit = true
	return tx.Commit().Error
}

func (d *PlugmanDao) SetRarity(rarity *model.Rarity) error {
	return d.DB.Model(model.Rarity{}).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(rarity).Error
}

func (d *PlugmanDao) CreateMetadata(metadataList *[]model.Metadata) error {
	return d.DB.Model(model.Metadata{}).Create(metadataList).Error
}

func (d *PlugmanDao) GetMetadataByAddressAndMintTypeAndNonce(address string, nonce uint, mintType uint, metadataList *[]model.Metadata) error {
	return d.DB.Model(model.Metadata{}).Where("address = ? AND nonce = ? AND mint_type = ?", address, nonce, mintType).Find(metadataList).Error
}

func (d *PlugmanDao) CreateTssTransfer(transfer *model.TssTransfer) error {
	return d.DB.Model(model.TssTransfer{}).Clauses(clause.Insert{Modifier: "IGNORE"}).Create(transfer).Error
}

func (d *PlugmanDao) GetTssTransfer(address string, transfers *[]model.TssTransfer) error {
	return d.DB.Model(model.TssTransfer{}).Where("from_address = ?", address).Find(transfers).Error
}

func (d *PlugmanDao) GetUnfinishedTssTransfer(transfers *[]model.TssTransfer) error {
	return d.DB.Model(model.TssTransfer{}).Where("order_status = ?", common.ORDER_INIT).Find(transfers).Error
}

func (d *PlugmanDao) UpdateTssTransferStatus(txHash string, status int) error {
	return d.DB.Model(model.TssTransfer{}).Where("tx_hash = ?", txHash).Update("order_status", status).Error
}
