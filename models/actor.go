package models

import (
	"GoDeveloperVK-testTask/utils"
)

type Actor struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Sex         string `json:"sex" db:"sex"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
}

type ActorWithMovies struct {
	Actor  Actor   `json:"actor"`
	Movies []Movie `json:"movies"`
}

func CreateActorTable() error {
	_, err := utils.DB.Exec(`
        CREATE TABLE IF NOT EXISTS actors (
            id SERIAL PRIMARY KEY,
            first_name VARCHAR(100),
            last_name VARCHAR(100),
            sex VARCHAR(10),
            date_of_birth TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	return nil
}

func ActorsEquals(actor1 Actor, actor2 Actor) bool {
	if actor1.FirstName != actor2.FirstName {
		return false
	} else if actor1.LastName != actor2.LastName {
		return false
	} else if actor1.Sex != actor2.Sex {
		return false
	} else if actor1.DateOfBirth != actor2.DateOfBirth {
		return false
	} else {
		return true
	}
}
