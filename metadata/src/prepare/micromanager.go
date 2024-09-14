package prepare

import (
	"encoding/json"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"github.com/PillLabs/plugman-monorepo/metadata/src/dao"
	"github.com/PillLabs/plugman-monorepo/metadata/src/model"
	"github.com/PillLabs/plugman-monorepo/metadata/src/premise"
	"github.com/gocarina/gocsv"
	"os"
	"strconv"
	"strings"
)

var MicromanagerInstance *Micromanager

type Micromanager struct {
	metadataJson string
	whitelistCsv string
	dao          *dao.PlugmanDao
}

func (m *Micromanager) Run() {
	m.prepareWhitelist()
	m.prepareMetadata()
}

func (m *Micromanager) prepareWhitelist() {
	var (
		whitelist  []WhitelistCSV
		addressWLs []model.AddressWL
	)

	log.Info("[Micromanager - PrepareWhitelist] start preparing whitelist to database")

	csvFile, err := os.Open(m.whitelistCsv)
	if err != nil {
		log.Fatal("[Micromanager - PrepareWhitelist] While opening whitelist, error: ", err)
	}
	defer func() {
		_ = csvFile.Close()
	}()

	err = gocsv.UnmarshalFile(csvFile, &whitelist)
	if err != nil {
		log.Fatal("[Micromanager - PrepareWhitelist] While unmarshalling whitelist: ", err)
	}

	for _, w := range whitelist {
		var addressWL model.AddressWL
		err = m.renderWhitelist(w, &addressWL)
		if err != nil {
			log.Fatal("[Micromanager - PrepareWhitelist] While rendering whitelist: ", err)
		}
		addressWLs = append(addressWLs, addressWL)
	}

	log.Info("[Micromanager - PrepareWhitelist] found %d entries of whitelist from csv file", len(addressWLs))

	if len(addressWLs) != 0 {
		err = m.dao.SetWhitelist(&addressWLs)
		if err != nil {
			log.Fatal("[Micromanager - PrepareWhitelist] While saving whitelist: ", err)
		}
		log.Info("[Micromanager - PrepareWhitelist] successfully saved whitelist")
	}
}

func (m *Micromanager) renderWhitelist(whitelist WhitelistCSV, addressWL *model.AddressWL) error {
	og, err := strconv.ParseUint(whitelist.OG, 10, 64)
	if err != nil {
		return err
	}
	wl, err := strconv.ParseUint(whitelist.Whitelist, 10, 64)
	if err != nil {
		return err
	}
	addressWL.Address = whitelist.Address
	addressWL.OGCount = uint(og)
	addressWL.WLCount = uint(wl)
	return nil
}

func (m *Micromanager) prepareMetadata() {
	var (
		metadataJson []SourceJson
		metadataList []model.Metadata
	)
	log.Debug("[Micromanager - PrepareMetadata] start preparing metadata to database")
	metadataBytes, err := os.ReadFile(m.metadataJson)
	if err != nil {
		log.Fatal("[Micromanager - PrepareMetadata] Could not read _metadata.json file")
	}

	err = json.Unmarshal(metadataBytes, &metadataJson)
	if err != nil {
		log.Fatal("[Micromanager - PrepareMetadata] Could not parse _metadata.json file")
	}
	for i, j := range metadataJson {
		metadata := model.Metadata{
			OriginalId: uint(i + 1),
		}
		m.renderMetadata(j, &metadata)
		metadataList = append(metadataList, metadata)
	}
	log.Info("[Micromanager - PrepareMetadata] found %d entries of metadata from json file", len(metadataList))
	err = m.dao.CreateMetadata(&metadataList)
	if err != nil {
		log.Fatal("[Micromanager - PrepareMetadata] Could not create metadata in database, error: %v", err)
	}
	log.Info("[Micromanager - PrepareMetadata] created %d entries of metadata in database", len(metadataList))
}

// TODO: rarity tag
func (m *Micromanager) renderMetadata(j SourceJson, metadata *model.Metadata) {
	var oneOfOne bool
	for _, attribute := range j.Attributes {
		switch attribute.TraitType {
		case common.TRAIT_TYPE_BACKGROUND:
			metadata.Background = strings.ToLower(attribute.Value)
		case common.TRAIT_TYPE_BODY_COLOR:
			metadata.BodyColor = strings.ToLower(attribute.Value)
		case common.TRAIT_TYPE_BODY:
			metadata.Body = strings.ToLower(attribute.Value)
		case common.TRAIT_TYPE_FRONT:
			metadata.Front = strings.ToLower(m.RaritySpliter(attribute.TraitType, attribute.Value, &oneOfOne))
		case common.TRAIT_TYPE_SIDE:
			metadata.Side = strings.ToLower(m.RaritySpliter(attribute.TraitType, attribute.Value, &oneOfOne))
		case common.TRAIT_TYPE_TOP:
			metadata.Top = strings.ToLower(m.RaritySpliter(attribute.TraitType, attribute.Value, &oneOfOne))
		default:
			log.Fatal("[Micromanager - renderMetadata] Unknown trait type: %s in name %s",
				attribute.TraitType, j.Name)
		}
	}
	metadata.OneOfOne = oneOfOne
}

func (m *Micromanager) RaritySpliter(traitKey, originalTraitValue string, oneOfOne *bool) string {
	traitSlice := strings.Split(originalTraitValue, "#")
	if len(traitSlice) == 2 {
		r := traitSlice[1]
		if r == "UR" {
			*oneOfOne = true
		}
		if traitElements := strings.Split(traitSlice[0], " "); len(traitElements) == 2 {
			rarity := model.Rarity{
				Name:     traitSlice[0],
				TraitKey: traitKey,
				Rarity:   r,
			}
			err := m.dao.SetRarity(&rarity)
			if err != nil {
				log.Fatal("[Micromanager - RaritySpliter] Could not save rarity in database, error: %v]")
			}
			return traitElements[0]
		} else {
			log.Error("[Micromanager - RaritySpliter] Unknown trait type: %s", originalTraitValue)
			return originalTraitValue
		}
	} else {
		log.Error("[Micromanager - RaritySpliter] Invalid trait value: %s", originalTraitValue)
		return originalTraitValue
	}
}

func init() {
	MicromanagerInstance = &Micromanager{
		metadataJson: conf.GConfig.GetString("micromanager.metadataJson"),
		whitelistCsv: conf.GConfig.GetString("micromanager.whitelistCsv"),
		dao:          premise.RuntimeVars.Dao,
	}
}
