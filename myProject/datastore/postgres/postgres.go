package postgres

import (
	"database/sql"
	// "fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	
		postgres_host = "dpg-ci2upie7avj2t35v6sk0-a.singapore-postgres.render.com"
		postgres_port = 5432
		postgres_user = "root"
		postgres_password = "qWRMPbevJb8TZ3MgqqPNeeLRfm4zcLq9"
		postgres_dbname = "chatbot_db_6mbu"
		
)

var Db *sql.DB

func init() { //to setup the database together with code

	// db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",postgres_host,postgres_port,postgres_user,postgres_password,postgres_dbname)
	connStr := "postgres://root:qWRMPbevJb8TZ3MgqqPNeeLRfm4zcLq9@dpg-ci2upie7avj2t35v6sk0-a.singapore-postgres.render.com:5432/chatbot_db_6mbu?sslmode=require"
	var err error
	Db, err = sql.Open("postgres", connStr)

	if err != nil {
		// log.Printf(err)
		panic(err)
	} else {
		log.Println("Database successfully connected")
	}
}
