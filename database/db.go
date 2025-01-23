package database

import(
    "database/sql"
    "fmt"
    "log"
    _"github.com/lib/pq"
)

var DB * sql.DB

func Connect(){
    connStr := "host=localhost port=5432 user=Tobi password='' dbname=task_db sslmode=disable"

    var err error

    DB, err = sql.Open("postgres", connStr)
    if err != nil{
        log.Fatalf("unable to connect to database: %v", err)
    }

    err = DB.Ping()
    if(err != nil){
        log.Fatalf("unable to connect to database: %v", err)

    }
    
    fmt.Println("connected to the db successfully")


}
