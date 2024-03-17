package database

import (
	"GoDeveloperVK-testTask/models"
	"GoDeveloperVK-testTask/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitDatabase(host string, port string, username string, password string, dbname string, sslmode string) (*sql.DB, error) {
	db, err := initPostgres(host, port, username, password, dbname, sslmode)
	if err != nil {
		return nil, err
	}

	utils.DB = db

	if err = models.CreateUserTable(); err != nil {
		return nil, err
	}

	if err = models.CreateActorTable(); err != nil {
		return nil, err
	}

	if err = models.CreateMovieTable(); err != nil {
		return nil, err
	}

	if err = models.CreateActorMovieTable(); err != nil {
		return nil, err
	}

	return db, nil
}

func initPostgres(host string, port string, username string, password string, dbname string, sslmode string) (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		host,
		port,
		dbname)

	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
