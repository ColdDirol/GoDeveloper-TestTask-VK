package repository

import (
	"GoDeveloperVK-testTask/models"
	"GoDeveloperVK-testTask/server/repository"
	"GoDeveloperVK-testTask/tests"
	"database/sql"
	"errors"
	"testing"
)

func TestAddMovie(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()

	movie := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}

	err := repository.AddMovie(movie)
	if err != nil {
		t.Errorf("Error adding movie: %v", err)
	}

	movies, err := repository.GetAllMovies()
	if err != nil {
		t.Errorf("Error getting movies slice: %v", err)
	}

	if len(*movies) == 0 {
		t.Error("Expected at least one movie, got none")
	} else if (*movies)[0].Movie.Name != movie.Name || (*movies)[0].Movie.Rating != movie.Rating || (*movies)[0].Movie.ReleaseDate != movie.ReleaseDate {
		t.Error("First movie does not match the expected movie")
	}

	repository.DeleteAllMovies()
}

func TestGetAllMovies(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()
	movies, err := repository.GetAllMovies()
	if err != nil {
		t.Errorf("Error getting movies slice: %v", err)
	}
	if len(*movies) != 0 {
		t.Errorf("Expected empty list of movies, got %d", len(*movies))
	}
	movie1 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}
	movie2 := models.MoviesCreate{
		Name:        "Inception",
		Rating:      8,
		ReleaseDate: "2010-07-16T00:00:00Z",
		ActorsID:    []int{3, 4},
	}
	repository.AddMovie(movie1)
	repository.AddMovie(movie2)

	movies, err = repository.GetAllMovies()
	if err != nil {
		t.Errorf("Error getting movies slice: %v", err)
	}
	if len(*movies) != 2 {
		t.Errorf("Expected 2 movies, got %d", len(*movies))
	}
	if (*movies)[0].Movie.Name != movie1.Name || (*movies)[1].Movie.Name != movie2.Name {
		t.Error("Movies slice does not match the expected movies")
	}

	repository.DeleteAllMovies()
}

func TestGetMovieByID(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()
	movie, err := repository.GetMovieByID(10000)
	if err == nil || !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("Expected movie not found error, got: %v", err)
	}
	if movie != nil {
		t.Errorf("Expected nil movie, got %+v", movie)
	}
	movie1 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}
	repository.AddMovie(movie1)

	movies, err := repository.GetAllMovies()

	movie, err = repository.GetMovieByID((*movies)[0].Movie.ID)
	if err != nil {
		t.Errorf("Error getting movie by ID: %v", err)
	}
	if movie.Movie.ID != (*movies)[0].Movie.ID {
		t.Errorf("Expected movie ID 1, got %d", movie.Movie.ID)
	}
	if movie.Movie.Name != movie1.Name || movie.Movie.Rating != movie1.Rating || movie.Movie.ReleaseDate != movie1.ReleaseDate {
		t.Error("Movie does not match the expected movie")
	}

	repository.DeleteAllMovies()
}

func TestUpdateMovie(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()
	movie1 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}
	repository.AddMovie(movie1)
	movies, err := repository.GetAllMovies()
	if err != nil {
		t.Errorf("Error getting movies slice: %v", err)
	}
	newMovie := models.MoviesCreate{
		Name:        "Inception",
		Rating:      8,
		ReleaseDate: "2010-07-16T00:00:00Z",
		ActorsID:    []int{3, 4},
	}
	err = repository.UpdateMovie((*movies)[0].Movie.ID, newMovie)
	if err != nil {
		t.Errorf("Error updating movie: %v", err)
	}
	updatedMovie, err := repository.GetMovieByID((*movies)[0].Movie.ID)
	if err != nil {
		t.Errorf("Error getting updated movie: %v", err)
	}
	if updatedMovie.Movie.Name != newMovie.Name || updatedMovie.Movie.Rating != newMovie.Rating || updatedMovie.Movie.ReleaseDate != newMovie.ReleaseDate {
		t.Error("Movie does not equals with the expected movie")
	}

	repository.DeleteAllMovies()
}

func TestDeleteMovieByID(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()
	movie1 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}
	repository.AddMovie(movie1)
	movies, err := repository.GetAllMovies()
	if err != nil {
		t.Errorf("Error getting movies slice: %v", err)
	}
	err = repository.DeleteMovieByID((*movies)[0].Movie.ID)
	if err != nil {
		t.Errorf("Error deleting movie by ID: %v", err)
	}
	movie, err := repository.GetMovieByID((*movies)[0].Movie.ID)
	if err == nil || !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("Expected movie not found error, got: %v", err)
	}
	if movie != nil {
		t.Errorf("Expected nil movie, got %+v", movie)
	}

	repository.DeleteAllMovies()
}

func TestDeleteAllMovies(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()
	movie1 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}
	movie2 := models.MoviesCreate{
		Name:        "Inception",
		Rating:      8,
		ReleaseDate: "2010-07-16T00:00:00Z",
		ActorsID:    []int{3, 4},
	}
	repository.AddMovie(movie1)
	repository.AddMovie(movie2)

	movies, err := repository.GetAllMovies()
	if err != nil {
		t.Errorf("Error getting movies slice: %v", err)
	}
	if len(*movies) != 2 {
		t.Errorf("Expected 2 movies, got %d", len(*movies))
	}
	err = repository.DeleteAllMovies()
	if err != nil {
		t.Errorf("Error deleting all movies: %v", err)
	}
	movies, err = repository.GetAllMovies()
	if err != nil {
		t.Errorf("Error getting movies slice: %v", err)
	}
	if len(*movies) != 0 {
		t.Errorf("Expected 0 movies, got %d", len(*movies))
	}
}

func TestFindMoviesByMovieName(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()
	movie1 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}
	movie2 := models.MoviesCreate{
		Name:        "Inception",
		Rating:      8,
		ReleaseDate: "2010-07-16T00:00:00Z",
		ActorsID:    []int{3, 4},
	}
	repository.AddMovie(movie1)
	repository.AddMovie(movie2)

	foundMovies, err := repository.FindMoviesByMovieName("Matrix")
	if err != nil {
		t.Errorf("Error finding movies by name: %v", err)
	}
	if len(*foundMovies) != 1 {
		t.Errorf("Expected 1 movie to be found, found %d", len(*foundMovies))
	}
	if (*foundMovies)[0].Movie.Name != "The Matrix" {
		t.Errorf("Expected movie name to be 'The Matrix', got %s", (*foundMovies)[0].Movie.Name)
	}

	foundMovies, err = repository.FindMoviesByMovieName("Unknown")
	if err != nil {
		t.Errorf("Error finding movies by name: %v", err)
	}
	if len(*foundMovies) != 0 {
		t.Errorf("Expected no movies to be found, found %d", len(*foundMovies))
	}

	repository.DeleteAllMovies()
}

func TestFindMoviesByActorName(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()
	repository.DeleteAllActors()

	actor1 := models.Actor{
		FirstName:   "DiCaprio",
		LastName:    "Doe",
		Sex:         "male",
		DateOfBirth: "1990-01-01T00:00:00Z",
	}
	repository.AddActor(actor1)

	actors, err := repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}

	movie1 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{(*actors)[0].ID},
	}
	movie2 := models.MoviesCreate{
		Name:        "Inception",
		Rating:      8,
		ReleaseDate: "2010-07-16T00:00:00Z",
		ActorsID:    []int{(*actors)[0].ID},
	}
	repository.AddMovie(movie1)
	repository.AddMovie(movie2)

	foundMovies, err := repository.FindMoviesByActorName("DiCaprio")
	if err != nil {
		t.Errorf("Error finding movies by actor name: %v", err)
	}
	if len(*foundMovies) != 2 {
		t.Errorf("Expected 1 movie to be found, found %d", len(*foundMovies))
	}
	if (*foundMovies)[1].Movie.Name != "Inception" {
		t.Errorf("Expected movie name to be 'Inception', got %s", (*foundMovies)[1].Movie.Name)
	}

	foundMovies, err = repository.FindMoviesByActorName("Unknown")
	if err != nil {
		t.Errorf("Error finding movies by actor name: %v", err)
	}
	if len(*foundMovies) != 0 {
		t.Errorf("Expected no movies to be found, found %d", len(*foundMovies))
	}

	repository.DeleteAllMovies()
	repository.DeleteAllActors()
}

func TestGetMoviesSortedByRating(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()

	movie1 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}
	movie2 := models.MoviesCreate{
		Name:        "Inception",
		Rating:      8,
		ReleaseDate: "2010-07-16T00:00:00Z",
		ActorsID:    []int{3, 4},
	}
	movie3 := models.MoviesCreate{
		Name:        "Interstellar",
		Rating:      10,
		ReleaseDate: "2014-11-07T00:00:00Z",
		ActorsID:    []int{5, 6},
	}

	repository.AddMovie(movie1)
	repository.AddMovie(movie2)
	repository.AddMovie(movie3)

	expectedOrder := []string{"Interstellar", "The Matrix", "Inception"}

	movies, err := repository.GetMoviesSortedByRating()
	if err != nil {
		t.Errorf("Error getting sorted movies by rating: %v", err)
	}

	if len(*movies) != len(expectedOrder) {
		t.Errorf("Expected %d movies, got %d", len(expectedOrder), len(*movies))
	}

	for i, movie := range *movies {
		if movie.Movie.Name != expectedOrder[i] {
			t.Errorf("Expected movie at position %d to be '%s', got '%s'", i, expectedOrder[i], movie.Movie.Name)
		}
	}

	repository.DeleteAllMovies()
}

func TestGetMoviesSortedByName(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()

	movie1 := models.MoviesCreate{
		Name:        "Inception",
		Rating:      8,
		ReleaseDate: "2010-07-16T00:00:00Z",
		ActorsID:    []int{3, 4},
	}
	movie2 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}
	movie3 := models.MoviesCreate{
		Name:        "Interstellar",
		Rating:      10,
		ReleaseDate: "2014-11-07T00:00:00Z",
		ActorsID:    []int{5, 6},
	}

	repository.AddMovie(movie1)
	repository.AddMovie(movie2)
	repository.AddMovie(movie3)

	expectedOrder := []string{"Inception", "Interstellar", "The Matrix"}

	movies, err := repository.GetMoviesSortedByName()
	if err != nil {
		t.Errorf("Error getting sorted movies by name: %v", err)
	}

	if len(*movies) != len(expectedOrder) {
		t.Errorf("Expected %d movies, got %d", len(expectedOrder), len(*movies))
	}

	for i, movie := range *movies {
		if movie.Movie.Name != expectedOrder[i] {
			t.Errorf("Expected movie at position %d to be '%s', got '%s'", i, expectedOrder[i], movie.Movie.Name)
		}
	}

	repository.DeleteAllMovies()
}

func TestGetMoviesSortedByReleaseDate(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllMovies()

	movie1 := models.MoviesCreate{
		Name:        "Inception",
		Rating:      8,
		ReleaseDate: "2010-07-16T00:00:00Z",
		ActorsID:    []int{3, 4},
	}
	movie2 := models.MoviesCreate{
		Name:        "Interstellar",
		Rating:      10,
		ReleaseDate: "2014-11-07T00:00:00Z",
		ActorsID:    []int{5, 6},
	}
	movie3 := models.MoviesCreate{
		Name:        "The Matrix",
		Rating:      9,
		ReleaseDate: "1999-03-31T00:00:00Z",
		ActorsID:    []int{1, 2},
	}

	repository.AddMovie(movie1)
	repository.AddMovie(movie2)
	repository.AddMovie(movie3)

	expectedOrder := []string{"Interstellar", "Inception", "The Matrix"}

	movies, err := repository.GetMoviesSortedByReleaseDate()
	if err != nil {
		t.Errorf("Error getting sorted movies by release date: %v", err)
	}

	if len(*movies) != len(expectedOrder) {
		t.Errorf("Expected %d movies, got %d", len(expectedOrder), len(*movies))
	}

	for i, movie := range *movies {
		if movie.Movie.Name != expectedOrder[i] {
			t.Errorf("Expected movie at position %d to be '%s', got '%s'", i, expectedOrder[i], movie.Movie.Name)
		}
	}

	repository.DeleteAllMovies()
}
