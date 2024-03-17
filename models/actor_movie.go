package models

import (
	"GoDeveloperVK-testTask/utils"
)

func CreateActorMovieTable() error {
	_, err := utils.DB.Exec(`
        CREATE TABLE IF NOT EXISTS actors_movies (
            actor_id INT,
            movie_id INT
        )
    `)
	if err != nil {
		return err
	}

	return nil
}
