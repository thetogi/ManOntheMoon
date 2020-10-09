package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", connectionString())
	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}
}
func connectionString() string {
	user := "root"
	password := "Password1!"
	address := "localhost"
	port := "3306"
	database := "ManOnTheMoon"
	params := "parseTime=true"
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", user, password, address, port, database, params)

}
