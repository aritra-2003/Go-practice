package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	
)


func main(){

	connn:="host=localhost port=5432 user=postgres password=ARga12@@ dbname=Demo sslmode=disable"
    
	db,err:=sql.Open("postgres",connn) 
	
	if err!=nil{
		log.Fatal(err)
	}

	defer db.Close()
	if err=db.Ping();err!=nil{
		log.Fatal(err)
	}
     createTables(db)
}
func createTables(db *sql.DB){
	query:= `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW())`

		_,err:=db.Exec(query)
		if err!=nil{
			log.Fatal(err)

		}
}