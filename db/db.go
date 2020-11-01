package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

//Global database object to make calls to SQL database
var Db *sql.DB

//Initialize DB global object to open connection for DB calls
func init() {
	var err error
	Db, err = sql.Open("mysql", connectionString())

	//Test connection to MySQL Server
	if err = Db.Ping(); err != nil {
		log.Panic(err)
	} else {
		log.Println("Connected to MySQL")
	}
}

//Builds the connection string for the MySQL DB
func connectionString() string {

	//Build connection string from environments variables.
	//Development provides them from the Goland IDE otherwise the docker-compose file will provide them from the .env file.

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")
	params := os.Getenv("PARAMS")

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", user, password, address, port, database, params)
}

//Known bug, use this connection string for local builds and debugging if not storing env variables in IDE.
//Development provides them from the Goland IDE otherwise the docker-compose file will provide them from the .env file.
//func connectionString() string {
//
//	user := "root"
//	password := "Password1!"
//	address := "localhost"
//	port := "3306"
//	database := "ManOnTheMoon"
//	params := "parseTime=true"
//
//	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", user, password, address, port, database, params)
//}
