package database

import (
	"log"
	"video-enconder/domain"

	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "Test"
	dbInstance.DbType = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatal("Test db error %w", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	// var err error

	if d.Env != "Test" {
		// d.Db, err = gorm.Open(d.DbType, d.Dsn)
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		d.Db = db
		if err != nil {
			return nil, err
		}
	} else {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		d.Db = db
		// d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
		if err != nil {
			return nil, err
		}
	}

	// if err != nil {
	// 	return nil, err
	// }

	if d.Debug {
		d.Db.Config.Logger.LogMode(logger.Info)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return d.Db, nil
}
