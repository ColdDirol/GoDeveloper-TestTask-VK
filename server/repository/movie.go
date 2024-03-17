package repository

import (
	"GoDeveloperVK-testTask/models"
	"GoDeveloperVK-testTask/utils"
	"errors"
	"log/slog"
)

func AddMovie(movie models.MoviesCreate) error {
	if movie.Rating < 0 && movie.Rating > 10 {
		return errors.New("invalid rating number")
	}

	tx, err := utils.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Вставляем информацию о фильме в таблицу movies и получаем ID
	var movieID int
	err = tx.QueryRow(`
		INSERT INTO movies (name, rating, release_date)
		VALUES ($1, $2, $3)
		RETURNING id
	`, movie.Name, movie.Rating, movie.ReleaseDate).Scan(&movieID)
	if err != nil {
		return err
	}

	// Вставляем связи с актерами в таблицу actors_movies, используя полученный movieID
	for _, actorID := range movie.ActorsID {
		_, err := tx.Exec(`
			INSERT INTO actors_movies (actor_id, movie_id)
			VALUES ($1, $2)
		`, actorID, movieID)
		if err != nil {
			return err
		}
	}

	utils.LOG.Info("New movie: ", slog.String("name", movie.Name))

	return nil
}

func GetAllMovies() (*[]models.MovieWithActors, error) {
	rows, err := utils.DB.Query(`
		SELECT id, name, rating, release_date
		FROM movies AS m
		ORDER BY m.rating DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moviesWithActors []models.MovieWithActors

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Rating, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}

		actors, err := GetActorByMovieID(movie.ID)
		if err != nil {
			return nil, err
		}

		movieWithActors := models.MovieWithActors{
			Movie: models.Movie{
				ID:          movie.ID,
				Name:        movie.Name,
				Rating:      movie.Rating,
				ReleaseDate: movie.ReleaseDate,
			},
			Actors: *actors,
		}

		moviesWithActors = append(moviesWithActors, movieWithActors)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &moviesWithActors, nil
}

func GetMovieByID(id int) (*models.MovieWithActors, error) {
	var movie models.Movie
	err := utils.DB.QueryRow(`
        SELECT id, name, rating, release_date
        FROM movies
        WHERE id = $1
    `, id).Scan(&movie.ID, &movie.Name, &movie.Rating, &movie.ReleaseDate)
	if err != nil {
		return nil, err
	}

	actors, err := GetActorByMovieID(id)
	if err != nil {
		return nil, err
	}

	movieWithActors := models.MovieWithActors{
		Movie: models.Movie{
			ID:          movie.ID,
			Name:        movie.Name,
			Rating:      movie.Rating,
			ReleaseDate: movie.ReleaseDate,
		},
		Actors: *actors,
	}

	return &movieWithActors, nil
}

func GetMoviesSortedByRating() (*[]models.MovieWithActors, error) {
	rows, err := utils.DB.Query(`
		SELECT id, name, rating, release_date
		FROM movies
		ORDER BY rating DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moviesWithActors []models.MovieWithActors

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Rating, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}

		actors, err := GetActorByMovieID(movie.ID)
		if err != nil {
			return nil, err
		}

		movieWithActors := models.MovieWithActors{
			Movie: models.Movie{
				ID:          movie.ID,
				Name:        movie.Name,
				Rating:      movie.Rating,
				ReleaseDate: movie.ReleaseDate,
			},
			Actors: *actors,
		}

		moviesWithActors = append(moviesWithActors, movieWithActors)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &moviesWithActors, nil
}

func GetMoviesSortedByName() (*[]models.MovieWithActors, error) {
	rows, err := utils.DB.Query(`
		SELECT id, name, rating, release_date
		FROM movies
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moviesWithActors []models.MovieWithActors

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Rating, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}

		actors, err := GetActorByMovieID(movie.ID)
		if err != nil {
			return nil, err
		}

		movieWithActors := models.MovieWithActors{
			Movie: models.Movie{
				ID:          movie.ID,
				Name:        movie.Name,
				Rating:      movie.Rating,
				ReleaseDate: movie.ReleaseDate,
			},
			Actors: *actors,
		}

		moviesWithActors = append(moviesWithActors, movieWithActors)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &moviesWithActors, nil
}

func GetMoviesSortedByReleaseDate() (*[]models.MovieWithActors, error) {
	rows, err := utils.DB.Query(`
		SELECT id, name, rating, release_date
		FROM movies
		ORDER BY release_date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moviesWithActors []models.MovieWithActors

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Rating, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}

		actors, err := GetActorByMovieID(movie.ID)
		if err != nil {
			return nil, err
		}

		movieWithActors := models.MovieWithActors{
			Movie: models.Movie{
				ID:          movie.ID,
				Name:        movie.Name,
				Rating:      movie.Rating,
				ReleaseDate: movie.ReleaseDate,
			},
			Actors: *actors,
		}

		moviesWithActors = append(moviesWithActors, movieWithActors)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &moviesWithActors, nil
}

func FindMoviesByMovieName(nameFragment string) (*[]models.MovieWithActors, error) {
	query := `
		SELECT m.id, m.name, m.rating, m.release_date
		FROM movies m
		WHERE LOWER(m.name) LIKE '%' || LOWER($1) || '%'
		ORDER BY m.rating DESC
	`

	rows, err := utils.DB.Query(query, nameFragment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moviesWithActors []models.MovieWithActors

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Rating, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}

		actors, err := GetActorByMovieID(movie.ID)
		if err != nil {
			return nil, err
		}

		movieWithActors := models.MovieWithActors{
			Movie: models.Movie{
				ID:          movie.ID,
				Name:        movie.Name,
				Rating:      movie.Rating,
				ReleaseDate: movie.ReleaseDate,
			},
			Actors: *actors,
		}

		moviesWithActors = append(moviesWithActors, movieWithActors)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &moviesWithActors, nil
}

func FindMoviesByActorName(nameFragment string) (*[]models.MovieWithActors, error) {
	query := `
		SELECT m.id, m.name, m.rating, m.release_date
		FROM movies m
		INNER JOIN actors_movies am ON m.id = am.movie_id
		INNER JOIN actors a ON am.actor_id = a.id
		WHERE LOWER(CONCAT(a.first_name, ' ', a.last_name)) LIKE '%' || LOWER($1) || '%'
		ORDER BY m.rating DESC
	`

	rows, err := utils.DB.Query(query, nameFragment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moviesWithActors []models.MovieWithActors

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Rating, &movie.ReleaseDate)
		if err != nil {
			return nil, err
		}

		actors, err := GetActorByMovieID(movie.ID)
		if err != nil {
			return nil, err
		}

		movieWithActors := models.MovieWithActors{
			Movie: models.Movie{
				ID:          movie.ID,
				Name:        movie.Name,
				Rating:      movie.Rating,
				ReleaseDate: movie.ReleaseDate,
			},
			Actors: *actors,
		}

		moviesWithActors = append(moviesWithActors, movieWithActors)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &moviesWithActors, nil
}

func UpdateMovie(id int, newMovie models.MoviesCreate) error {
	tx, err := utils.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Обновляем информацию о фильме в таблице movies
	_, err = tx.Exec(`
		UPDATE movies
		SET name = $1, rating = $2, release_date = $3
		WHERE id = $4
	`, newMovie.Name, newMovie.Rating, newMovie.ReleaseDate, id)
	if err != nil {
		return err
	}

	// Удаляем все зависимости в таблице actors_movies для данного фильма
	_, err = tx.Exec(`
		DELETE FROM actors_movies WHERE movie_id = $1
	`, id)
	if err != nil {
		return err
	}

	// Заново связываем актеров и фильмы в таблице actors_movies
	for _, actorID := range newMovie.ActorsID {
		_, err := tx.Exec(`
			INSERT INTO actors_movies (actor_id, movie_id)
			VALUES ($1, $2)
		`, actorID, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteMovieByID(id int) error {
	tx, err := utils.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Удаляем запись из таблицы movies по ID
	_, err = tx.Exec(`
		DELETE FROM movies
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	// Удаляем все записи с movie_id равным данному ID из таблицы actors_movies
	_, err = tx.Exec(`
		DELETE FROM actors_movies
		WHERE movie_id = $1
	`, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAllMovies() error {
	_, err := utils.DB.Exec(`
		DELETE FROM movies
	`)
	if err != nil {
		return errors.New("failed to delete actors: " + err.Error())
	}
	DeleteAllActorsMovies()
	return nil
}

func DeleteAllActorsMovies() error {
	_, err := utils.DB.Exec(`
		DELETE FROM actors_movies
	`)
	if err != nil {
		return errors.New("failed to delete actors: " + err.Error())
	}
	return nil
}
