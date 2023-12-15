package processing

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Connector struct {
	User   string
	Passwd string
	Addr   string
	DBName string
}

func ConnectToDatabase(connector Connector) {

	conf := mysql.Config{
		User:   connector.User,
		Passwd: connector.Passwd,
		Net:    "tcp",
		Addr:   connector.Addr,
		DBName: connector.DBName,
	}

	var err error
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
func connectToFile()          {}
func connectToDataStreaming() {}
func dataToObject()           {}
