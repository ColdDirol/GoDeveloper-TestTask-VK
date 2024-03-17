package models

import (
	"GoDeveloperVK-testTask/utils"
)

type MoviesCreate struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Rating      int    `json:"rating" db:"rating"`
	ReleaseDate string `json:"release_date" db:"release_date"`
	ActorsID    []int  `json:"actors_id"`
}

type Movie struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Rating      int    `json:"rating" db:"rating"`
	ReleaseDate string `json:"release_date" db:"release_date"`
}

type MoviesByName struct {
	Name string `json:"name"`
}

type MovieWithActors struct {
	Movie  Movie   `json:"movie"`
	Actors []Actor `json:"actors"`
}

func CreateMovieTable() error {
	_, err := utils.DB.Exec(`
        CREATE TABLE IF NOT EXISTS movies (
            id SERIAL PRIMARY KEY,
            name VARCHAR(150),
            rating INTEGER,
            release_date TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}

	return nil
}

func MoviesEquals(movie1 Movie, movie2 Movie) bool {
	if movie1.Name != movie2.Name {
		return false
	} else if movie1.Rating != movie2.Rating {
		return false
	} else if movie1.ReleaseDate != movie2.ReleaseDate {
		return false
	} else {
		return true
	}
}
