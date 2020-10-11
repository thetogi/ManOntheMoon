package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	}
}

//Builds the connection string for the MySQL db
func connectionString() string {
	mydir, _ := os.Getwd()
	fmt.Println(mydir)
	//Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Build connection string from .env config
	envFile, err := godotenv.Read()
	user := envFile["USER"]
	password := envFile["PASSWORD"]
	address := envFile["ADDRESS"]
	port := envFile["PORT"]
	database := envFile["DATABASE"]
	params := envFile["PARAMS"]

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", user, password, address, port, database, params)

}

//func connectionString() string {
//	user := "root"
//	password := "Password1!"
//	address := "localhost"
//	port := "3306"
//	database := "ManOnTheMoon"
//	params := "parseTime=true"
//	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", user, password, address, port, database, params)
//
//}
