package go_basic_orm

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connection(database string) (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error al obtener configuración del servidor")
	}
	server := os.Getenv("ENV_DDBB_SERVER")
	user := os.Getenv("ENV_DDBB_USER")
	password := os.Getenv("ENV_DDBB_PASSWORD")
	port := os.Getenv("ENV_DDBB_PORT")
	connection_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, database)
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		errs := errors.New(fmt.Sprintf("Error connection: %s ", err.Error()))
		return nil, errs
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		errs := errors.New(fmt.Sprintf("Error creating connection: %s", err.Error()))
		return nil, errs
	}
	return db, nil
}

func ConnectionCloud() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error al obtener configuración del servidor cloud")
	}
	server := os.Getenv("ENV_DDBB_SERVER")
	user := os.Getenv("ENV_DDBB_USER")
	password := os.Getenv("ENV_DDBB_PASSWORD")
	database := os.Getenv("ENV_DDBB_DATABASE")
	port := os.Getenv("ENV_DDBB_PORT")

	connection_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, database)

	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		errs := errors.New(fmt.Sprintf("Error connection: %s ", err.Error()))
		return nil, errs
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		errs := errors.New(fmt.Sprintf("Error creating connection: %s", err.Error()))
		return nil, errs
	}
	return db, nil
}
