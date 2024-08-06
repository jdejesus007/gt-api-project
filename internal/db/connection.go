package db

import (
	"fmt"
	"log"

	"github.com/jdejesus007/gt-api-project/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBExecutor interface {
	GetConn() *gorm.DB
}

type GtDB struct {
	connection *gorm.DB
}

// NewGtDB creates a new gt database connection
func NewGtDB(connection *gorm.DB) *GtDB {
	database := &GtDB{
		connection: connection,
	}
	return database
}

func (dbImpl *GtDB) GetConn() *gorm.DB {
	// If connection already exists, use it
	if dbImpl.connection != nil {
		return dbImpl.connection
	}

	conf := &gorm.Config{
		PrepareStmt: false,
	}

	var err error
	var db *gorm.DB

	dsn := fmt.Sprintf(config.GtConfig.DatabaseDSN,
		config.GtConfig.DatabaseHost,
		config.GtConfig.DatabaseUser,
		config.GtConfig.DatabasePassword,
		config.GtConfig.DatabaseName,
		config.GtConfig.DatabasePort)

	db, err = gorm.Open(postgres.Open(dsn), conf)
	if err != nil {
		log.Fatalln("failed to open conn", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("failed to acquire db", err)
		return nil
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalln("failed to verify live conn with ping", err)
		return nil
	}

	log.Println("Successfully connected to DB")

	dbImpl.connection = db

	return dbImpl.connection
}
