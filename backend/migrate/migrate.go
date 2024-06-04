package main

import (
	"fmt"

	"github.com/tigaweb/reversi-app/backend/db"
	"github.com/tigaweb/reversi-app/backend/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}
