package service

import (
	"GoDeveloperVK-testTask/models"
	"GoDeveloperVK-testTask/server/repository"
	"GoDeveloperVK-testTask/utils"
	"GoDeveloperVK-testTask/utils/logger"
	"encoding/json"
	"net/http"
)

func GetMovies(w http.ResponseWriter) {
	moviesWithActors, err := repository.GetAllMovies()
	if err != nil {
		utils.LOG.Error("failed to get movies", logger.Err(err))
		http.Error(w, "failed to get movie", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(moviesWithActors)
	if err != nil {
		utils.LOG.Error("failed to marshal movies data", logger.Err(err))
		http.Error(w, "failed to marshal movies data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetSortedMoviesByRating(w http.ResponseWriter) {
	moviesWithActors, err := repository.GetMoviesSortedByRating()
	if err != nil {
		utils.LOG.Error("failed to get movies sorted by rating", logger.Err(err))
		http.Error(w, "failed to get movies sorted by rating", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(moviesWithActors)
	if err != nil {
		utils.LOG.Error("failed to marshal movies data", logger.Err(err))
		http.Error(w, "failed to marshal movies data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetSortedMoviesByName(w http.ResponseWriter) {
	moviesWithActors, err := repository.GetMoviesSortedByName()
	if err != nil {
		utils.LOG.Error("failed to get movies sorted by name", logger.Err(err))
		http.Error(w, "failed to get movies sorted by name", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(moviesWithActors)
	if err != nil {
		utils.LOG.Error("failed to marshal movies data", logger.Err(err))
		http.Error(w, "failed to marshal movies data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetSortedMoviesByReleaseDate(w http.ResponseWriter) {
	moviesWithActors, err := repository.GetMoviesSortedByReleaseDate()
	if err != nil {
		utils.LOG.Error("failed to get movies sorted by release date", logger.Err(err))
		http.Error(w, "failed to get movies sorted by release date", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(moviesWithActors)
	if err != nil {
		utils.LOG.Error("failed to marshal movies data", logger.Err(err))
		http.Error(w, "failed to marshal movies data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func FindMoviesByMovieName(w http.ResponseWriter, r *http.Request) {
	var movieByName models.MoviesByName
	err := json.NewDecoder(r.Body).Decode(&movieByName)
	if err != nil {
		utils.LOG.Error("failed to decode movies by movie name data", logger.Err(err))
		http.Error(w, "failed to decode movies by movie name data", http.StatusBadRequest)
		return
	}

	movies, err := repository.FindMoviesByMovieName(movieByName.Name)
	if err != nil {
		utils.LOG.Error("failed to find movies by movie name", logger.Err(err))
		http.Error(w, "failed to find movies by movie name", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(movies)
	if err != nil {
		utils.LOG.Error("failed to marshal movies data", logger.Err(err))
		http.Error(w, "failed to marshal movies data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func FindMoviesByActorName(w http.ResponseWriter, r *http.Request) {
	var movieByName models.MoviesByName
	err := json.NewDecoder(r.Body).Decode(&movieByName)
	if err != nil {
		utils.LOG.Error("failed to decode movies by actor name data", logger.Err(err))
		http.Error(w, "failed to decode movies by actor name data", http.StatusBadRequest)
		return
	}

	movies, err := repository.FindMoviesByActorName(movieByName.Name)
	if err != nil {
		utils.LOG.Error("failed to find movies by actor name", logger.Err(err))
		http.Error(w, "failed to find movies by actor name", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(movies)
	if err != nil {
		utils.LOG.Error("failed to marshal movies data", logger.Err(err))
		http.Error(w, "failed to marshal movies data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func PostMovies(w http.ResponseWriter, r *http.Request) {
	var movie models.MoviesCreate
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		utils.LOG.Error("failed to decode movie data", logger.Err(err))
		http.Error(w, "failed to decode movie data", http.StatusBadRequest)
		return
	}

	err = repository.AddMovie(movie)
	if err != nil {
		utils.LOG.Error("failed to insert movie", logger.Err(err))
		http.Error(w, "failed to insert movie", http.StatusBadRequest)
		return
	}
}

func UpdateMovies(w http.ResponseWriter, r *http.Request, movieID int) {
	var movie models.MoviesCreate
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		utils.LOG.Error("failed to decode movie data", logger.Err(err))
		http.Error(w, "failed to decode movie data", http.StatusInternalServerError)
		return
	}

	err = repository.UpdateMovie(movieID, movie)
	if err != nil {
		utils.LOG.Error("failed to update movie", logger.Err(err))
		http.Error(w, "failed to update movie", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteMovies(w http.ResponseWriter, movieID int) {
	err := repository.DeleteMovieByID(movieID)
	if err != nil {
		utils.LOG.Error("failed to delete movie", logger.Err(err))
		http.Error(w, "failed to delete movie", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
