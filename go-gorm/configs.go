package main 

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)


func LoadConfig(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDbConfig() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName :=os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	return dsn
}
