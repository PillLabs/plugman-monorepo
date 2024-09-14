package database

import (
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"github.com/PillLabs/plugman-monorepo/metadata/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := conf.GConfig.GetString("mysql.db.dsn")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("[PANIC][Database - Init] Open dsn error: %s", err)
	}
	mysqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("[PANIC][Database - Init] Get mysqlDB error: %s", err)
	}
	mysqlDB.SetMaxIdleConns(conf.GConfig.GetInt("mysql.db.idle"))
	mysqlDB.SetMaxOpenConns(conf.GConfig.GetInt("mysql.db.active"))
	mysqlDB.SetConnMaxLifetime(time.Second * time.Duration(conf.GConfig.GetInt("mysql.db.connMaxLifeTime")))
	log.Info("[Database - Init] database client setting done")
	err = DB.AutoMigrate(
		&model.Metadata{},
		&model.Nonce{},
		&model.Block{},
		&model.Rarity{},
		&model.TssTransfer{})
	if err != nil {
		log.Fatal("[PANIC][Database -Init] Auto migrate table schema error: %s", err)
	}
	log.Info("[Database - Init] database schema auto migrated successfully")
}
