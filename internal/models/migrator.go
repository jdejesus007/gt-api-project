package models

import (
	"log"

	"github.com/jdejesus007/gt-api-project/internal/db"
)

func AutoMigrate(DBExecutor db.DBExecutor) {
	ifaces := []interface{}{
		&Book{},
	}

	DBExecutor.GetConn().Debug().AutoMigrate(ifaces...)

	log.Println("Migrations Complete")
}