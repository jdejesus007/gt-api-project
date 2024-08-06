package models

import (
	"log"

	"github.com/jdejesus007/gt-api-project/internal/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/google/uuid"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func AutoSeed(dbExecutor db.DBExecutor) {
	dbConn := dbExecutor.GetConn().Debug().Clauses(clause.OnConflict{DoNothing: true})

	log.Println("uuid", uuid.New().String())

	// Insert test data for local postgres testing
	dbConn.Exec(`INSERT INTO "books" ("created_at","updated_at","deleted_at","base_status","uuid", "name","author") VALUES ('2024-08-05 12:03:41.937337','2024-08-05 12:03:41.937337',NULL,0,'ba398055-8df8-497a-af1d-bc2fcf20b03d','Design Patterns','Gangs of 4')`)
	dbConn.Exec(`INSERT INTO "books" ("created_at","updated_at","deleted_at","base_status","uuid", "name","author") VALUES ('2024-08-05 15:03:41.937337','2024-08-05 15:03:41.937337',NULL,0,'662e342d-0929-4cd5-bd9a-6b9913d61b71','Go Programming Language','Donovan/Kernighan')`)

	log.Println("Completed seeding data")
}
