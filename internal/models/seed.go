package models

import (
	"log"

	"github.com/jdejesus007/gt-api-project/internal/db"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func AutoSeed(dbExecutor db.DBExecutor) {
	// dbConn := dbExecutor.GetConn().Debug().Clauses(clause.OnConflict{DoNothing: true})
	// for _, seed := range all() {
	//   if err := seed.Run(dbConn); err != nil {
	//     log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
	//   }
	// }

	// Insert test data for local postgres testing
	// dbConn.Exec(`INSERT INTO "books" ("created_at","updated_at","deleted_at","base_status","name") VALUES ('2024-08-05 09:09:12.388','2024-08-05 19:09:12.388',NULL,0,'Go Rocks')`)

	log.Println("Completed seeding data")
}

// func all() []Seed {
//   return []Seed{
//     Seed{
//       Name: "CreateTestData",
//       Run: func(db *gorm.DB) error {
//         return createTestData(
//           db,
//         )
//       },
//     },
//   }
// }
//
// func createTestData(db *gorm.DB) error {
//   return db.Create(&Book{
//     Name:  "Go Rocks",
//   }).Error
// }
