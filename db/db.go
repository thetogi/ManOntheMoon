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
		log.Println(connectionString())
		log.Panic("MYSQL: " + err.Error())
	} else {
		log.Println("Connected to MySQL")
	}
}

//Builds the connection string for the MySQL DB
func connectionString() string {

	//Build connection string from environments variables.
	//Development provides them from the Goland IDE otherwise the docker-compose file will provide them from the db.dev.env file.

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	address := os.Getenv("MYSQL_DB_ADDRESS")
	port := os.Getenv("MYSQL_DB_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	params := os.Getenv("MYSQL_DB_PARAMS")

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", user, password, address, port, database, params)
}

//Known bug, use this connection string for local builds and debugging if not storing env variables in IDE.
//Development provides them from the Goland IDE otherwise the docker-compose file will provide them from the db.dev.env file.
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
