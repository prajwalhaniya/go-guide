package main

import ( 
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:rootadmin@(127.0.0.1:3306)/appdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Create a new table
        query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

        if _, err := db.Exec(query); err != nil {
            log.Fatal(err)
        }
    }
}