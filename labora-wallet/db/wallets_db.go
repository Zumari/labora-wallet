package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// PostgresDbHandler contains a pointer to the SQL database.
type PostgresDbHandler struct {
	*sql.DB
}

// GetEnvCredentials obtain the .env file credentials, transforms them for their correct use and returns them.
func GetEnvCredentials() (string, string, string, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	dbName := os.Getenv("dbName")
	rolName := os.Getenv("rolName")
	rolPassword := os.Getenv("rolPassword")

	return host, port, dbName, rolName, rolPassword

}

var DBHandler *PostgresDbHandler

// PingOrDie indicate whether or not there is connection with the database server.
func (db *PostgresDbHandler) PingOrDie() {
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot reach database, error: %v", err)
	}
}

// Connect_DB establish the connection with the database.
func Connect_DB() (*PostgresDbHandler, error) {
	host, port, dbName, rolName, rolPassword := GetEnvCredentials()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	fmt.Println("Successful connection to the database:", dbConn)
	DBHandler = &PostgresDbHandler{dbConn}
	DBHandler.PingOrDie()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The request to the database is active")

	return DBHandler, nil
}
