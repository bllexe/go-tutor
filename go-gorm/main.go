package main

import (
	"go-tutor/go-gorm/handler"
	"go-tutor/go-gorm/service"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	LoadConfig()
    // config.LoadConfig()
    dsn := GetDbConfig()

    // Open the MySQL database connection
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database", err)
        return nil, err
    }
    // Get the underlying *sql.DB instance for connection pooling
    sqlDb, err := db.DB()
    if err != nil {
        log.Fatal("failed to get database object", err)
        return nil, err
    }
    // Ensure that the database connection is alive
    if err := sqlDb.Ping(); err != nil {
        log.Fatal("failed to ping database", err)
        return nil, err
    }

    return db, nil
}

func main(){

	_,err := InitDB() 
	if err != nil {
		log.Fatal(err)
	}

	var svc service.UserService
	handleRequest(svc)
}

func handleRequest(svc service.UserService){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/getAllUsers",handler.AllUsersHandler(svc)).Methods("GET")
	myRouter.HandleFunc("/addUser",handler.AddUserHandler(svc)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
