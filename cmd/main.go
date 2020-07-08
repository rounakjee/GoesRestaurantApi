package main

import (
	"GoResturantAPI/DatabseProvider/Psql"
	"GoResturantAPI/servers"
	"fmt"
)



func main() {
	fmt.Println("API HAHAHA....")
	db := Psql.NewPSqlDB("bond","password","localhost","api")
	defer db.Close()
	ser := servers.NewServer(db)
	servers.InitialMigration(db.DB)
	ser.Start("8080")
}
