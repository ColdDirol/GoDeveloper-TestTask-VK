package repository

import (
	"GoDeveloperVK-testTask/models"
	"GoDeveloperVK-testTask/server/repository"
	"GoDeveloperVK-testTask/tests"
	"GoDeveloperVK-testTask/utils"
	"database/sql"
	"errors"
	"fmt"
	"testing"
)

func TestAddActor(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllActors()

	actor := models.Actor{
		FirstName:   "John",
		LastName:    "Doe",
		Sex:         "male",
		DateOfBirth: "2022-12-02T00:00:00Z",
	}

	err := repository.AddActor(actor)
	if err != nil {
		t.Errorf("Error adding actor: %v", err)
	}

	actors, err := repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}

	fmt.Println()

	if len(*actors) == 0 {
		t.Error("Expected at least one actor, got none")
	} else if !models.ActorsEquals((*actors)[0], actor) {
		fmt.Println((*actors)[0], actor)
		t.Error("First actor does not match the expected actor")
	}

	repository.DeleteAllActors()
}

func TestGetAllActors(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllActors()
	actors, err := repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}
	if len(*actors) != 0 {
		t.Errorf("Expected empty list of actors, got %d", len(*actors))
	}
	actor1 := models.Actor{
		FirstName:   "John",
		LastName:    "Doe",
		Sex:         "male",
		DateOfBirth: "1990-01-01T00:00:00Z",
	}
	actor2 := models.Actor{
		FirstName:   "Jane",
		LastName:    "Smith",
		Sex:         "female",
		DateOfBirth: "1985-05-05T00:00:00Z",
	}
	repository.AddActor(actor1)
	repository.AddActor(actor2)

	actors, err = repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}
	if len(*actors) != 2 {
		t.Errorf("Expected 2 actors, got %d", len(*actors))
	}
	if !models.ActorsEquals((*actors)[0], actor1) || !models.ActorsEquals((*actors)[1], actor2) {
		t.Error("Actors slice not match the actors1 and actors2")
	}

	repository.DeleteAllActors()
}

func TestGetActorByMovieID(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllActors()
	actors, err := repository.GetActorByMovieID(1)
	if err != nil {
		t.Errorf("Error getting actors by movie ID: %v", err)
	}
	if len(*actors) != 0 {
		t.Errorf("Expected empty list of actors, got %d", len(*actors))
	}
	actor1 := models.Actor{
		FirstName:   "John",
		LastName:    "Doe",
		Sex:         "male",
		DateOfBirth: "1990-01-01T00:00:00Z",
	}
	actor2 := models.Actor{
		FirstName:   "Jane",
		LastName:    "Smith",
		Sex:         "female",
		DateOfBirth: "1985-05-05T00:00:00Z",
	}
	repository.AddActor(actor1)
	repository.AddActor(actor2)

	actors, err = repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}
	if len(*actors) != 2 {
		t.Errorf("Expected 2 actors, got %d", len(*actors))
	}
	if !models.ActorsEquals((*actors)[0], actor1) || !models.ActorsEquals((*actors)[1], actor2) {
		t.Error("Actors slice not match the actors1 and actors2")
	}

	_, err = utils.DB.Exec("INSERT INTO actors_movies (actor_id, movie_id) VALUES ($1, $2)", (*actors)[0].ID, 1)
	if err != nil {
		t.Errorf("Error assigning actors to movie: %v", err)
	}
	_, err = utils.DB.Exec("INSERT INTO actors_movies (actor_id, movie_id) VALUES ($1, $2)", (*actors)[1].ID, 1)
	if err != nil {
		t.Errorf("Error assigning actors to movie: %v", err)
	}

	actors, err = repository.GetActorByMovieID(1)
	if err != nil {
		t.Errorf("Error getting actors by movie ID: %v", err)
	}
	if len(*actors) != 2 {
		t.Errorf("Expected 2 actors, got %d", len(*actors))
	}

	repository.DeleteAllActors()
}

func TestGetActorByID(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllActors()
	repository.DeleteAllMovies()
	actor, err := repository.GetActorByID(1)
	if err == nil || !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("Expected actor not found error, got: %v", err)
	}
	if actor != nil {
		t.Errorf("Expected nil actor, got %+v", actor)
	}
	actor1 := models.Actor{
		FirstName:   "John",
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
		Name:        "awesomeMovie",
		Rating:      10,
		ReleaseDate: "2024-01-01T00:00:00Z",
		ActorsID:    []int{(*actors)[0].ID},
	}
	repository.AddMovie(movie1)

	movies, err := repository.GetAllMovies()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}
	actor, err = repository.GetActorByID((*actors)[0].ID)
	if err != nil {
		t.Errorf("Error getting actor by ID: %v", err)
	}
	if actor.Actor.ID != (*actors)[0].ID {
		t.Errorf("Expected actor ID 1, got %d", actor.Actor.ID)
	}
	if !models.ActorsEquals(actor1, actor.Actor) {
		fmt.Println((*actors)[0], actor)
		t.Error("First actor does not match the expected actor")
	}
	if !models.MoviesEquals((*movies)[0].Movie, actor.Movies[0]) {

	}

	repository.DeleteAllActors()
	repository.DeleteAllMovies()
}

func TestUpdateActor(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllActors()
	actor1 := models.Actor{
		FirstName:   "John",
		LastName:    "Doe",
		Sex:         "male",
		DateOfBirth: "1990-01-01T00:00:00Z",
	}
	repository.AddActor(actor1)
	actors, err := repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}
	newActor := models.Actor{
		FirstName:   "Jane",
		LastName:    "Smith",
		Sex:         "female",
		DateOfBirth: "1985-05-05T00:00:00Z",
	}
	err = repository.UpdateActor((*actors)[0].ID, newActor)
	if err != nil {
		t.Errorf("Error updating actor: %v", err)
	}
	updatedActor, err := repository.GetActorByID((*actors)[0].ID)
	if err != nil {
		t.Errorf("Error getting updated actor: %v", err)
	}
	if !models.ActorsEquals(updatedActor.Actor, newActor) {
		t.Error("Actor does not equals with the expected actor")
	}

	repository.DeleteAllActors()
}

func TestDeleteActorByID(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllActors()
	actor1 := models.Actor{
		FirstName:   "John",
		LastName:    "Doe",
		Sex:         "male",
		DateOfBirth: "1990-01-01T00:00:00Z",
	}
	repository.AddActor(actor1)
	actors, err := repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}
	err = repository.DeleteActorByID((*actors)[0].ID)
	if err != nil {
		t.Errorf("Error deleting actor by ID: %v", err)
	}
	actor, err := repository.GetActorByID((*actors)[0].ID)
	if err == nil || !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("Expected actor not found error, got: %v", err)
	}
	if actor != nil {
		t.Errorf("Expected nil actor, got %+v", actor)
	}

	repository.DeleteAllActors()
}

func TestDeleteAllActors(t *testing.T) {
	tests.StartDB()
	repository.DeleteAllActors()
	actor1 := models.Actor{
		FirstName:   "John",
		LastName:    "Doe",
		Sex:         "male",
		DateOfBirth: "1990-01-01T00:00:00Z",
	}
	actor2 := models.Actor{
		FirstName:   "Jane",
		LastName:    "Smith",
		Sex:         "female",
		DateOfBirth: "1985-05-05T00:00:00Z",
	}
	repository.AddActor(actor1)
	repository.AddActor(actor2)

	actors, err := repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}
	if len(*actors) != 2 {
		t.Errorf("Expected 2 actors, got %d", len(*actors))
	}
	err = repository.DeleteAllActors()
	if err != nil {
		t.Errorf("Error deleting all actors: %v", err)
	}
	actors, err = repository.GetAllActors()
	if err != nil {
		t.Errorf("Error getting actors slice: %v", err)
	}
	if len(*actors) != 0 {
		t.Errorf("Expected 0 actors, got %d", len(*actors))
	}
}
