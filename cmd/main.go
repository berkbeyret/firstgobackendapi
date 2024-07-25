package main

import (
	"log"

	"github.com/berkbeyret/firstgobackendapi/cmd/api"
	"github.com/berkbeyret/firstgobackendapi/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "pass",
		Addr:                 "localhost:3306",
		DBName:               "firstgobackendapi",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
