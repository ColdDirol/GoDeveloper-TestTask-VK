package repository

import (
	"GoDeveloperVK-testTask/models"
	"GoDeveloperVK-testTask/utils"
	"errors"
	"log/slog"
)

func AddActor(actor models.Actor) error {
	_, err := utils.DB.Exec(`
		INSERT INTO actors (first_name, last_name, sex, date_of_birth) 
		VALUES ($1, $2, $3, $4)
	`, actor.FirstName, actor.LastName, actor.Sex, actor.DateOfBirth)
	if err != nil {
		return err
	}

	utils.LOG.Info("New actor: ", slog.String("first name", actor.FirstName), slog.String("last name", actor.LastName))

	return nil
}

func GetAllActors() (*[]models.Actor, error) {
	rows, err := utils.DB.Query(`
		SELECT id, first_name, last_name, sex, date_of_birth 
		FROM actors
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []models.Actor

	for rows.Next() {
		var actor models.Actor
		err := rows.Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.Sex, &actor.DateOfBirth)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &actors, nil
}

func GetActorByMovieID(movieID int) (*[]models.Actor, error) {
	rows, err := utils.DB.Query(`
		SELECT a.id, a.first_name, a.last_name, a.sex, a.date_of_birth
		FROM actors a
		JOIN actors_movies am ON a.id = am.actor_id
		WHERE am.movie_id = $1
	`, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []models.Actor

	for rows.Next() {
		var actor models.Actor
		err := rows.Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.Sex, &actor.DateOfBirth)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &actors, nil
}

func GetActorByID(id int) (*models.ActorWithMovies, error) {
	var actor models.Actor
	err := utils.DB.QueryRow(`
        SELECT id, first_name, last_name, sex, date_of_birth 
        FROM actors
        WHERE id = $1
    `, id).Scan(&actor.ID, &actor.FirstName, &actor.LastName, &actor.Sex, &actor.DateOfBirth)
	if err != nil {
		return nil, err
	}

	var movies []models.Movie
	rows, err := utils.DB.Query(`
        SELECT m.id, m.name, m.rating, m.release_date
        FROM movies m
        JOIN actors_movies am ON m.id = am.movie_id
        WHERE am.actor_id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Rating, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &models.ActorWithMovies{Actor: actor, Movies: movies}, nil
}

func UpdateActor(id int, newActor models.Actor) error {
	_, err := utils.DB.Exec(`
		UPDATE actors
		SET first_name = $1, last_name = $2, sex = $3, date_of_birth = $4
		WHERE id = $5
	`, newActor.FirstName, newActor.LastName, newActor.Sex, newActor.DateOfBirth, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteActorByID(id int) error {
	_, err := utils.DB.Exec(`
		DELETE FROM actors 
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllActors() error {
	_, err := utils.DB.Exec(`
		DELETE FROM actors
	`)
	if err != nil {
		return errors.New("failed to delete actors: " + err.Error())
	}
	DeleteAllActorsMovies()
	return nil
}
