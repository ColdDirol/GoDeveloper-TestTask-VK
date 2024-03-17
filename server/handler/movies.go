package handler

import (
	"GoDeveloperVK-testTask/auth/middleware"
	"GoDeveloperVK-testTask/server/service"
	"GoDeveloperVK-testTask/utils"
	"GoDeveloperVK-testTask/utils/logger"
	"net/http"
)

const (
	rating      = "rating"
	name        = "name"
	releaseDate = "releaseDate"
	movie       = "movie"
	actor       = "actor"
)

func InitMoviesHandlers() {
	http.HandleFunc("/movies", middleware.Middleware(moviesHandler))
	http.HandleFunc("/movies/sort/", middleware.Middleware(sortedMoviesHandler))
	http.HandleFunc("/movies/find/", middleware.Middleware(findMoviesHandler))
	http.HandleFunc("/movies/", middleware.Middleware(movieByIDHandler))
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		service.GetMovies(w)
	case http.MethodPost:
		service.PostMovies(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func movieByIDHandler(w http.ResponseWriter, r *http.Request) {
	movieID, err := utils.ExtractIDFromURL(r.URL.Path)
	if err != nil {
		utils.LOG.Error("invalid id", logger.Err(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		service.UpdateMovies(w, r, movieID)
	case http.MethodDelete:
		service.DeleteMovies(w, movieID)
	default:
		utils.LOG.Error("invalid http method", http.StatusMethodNotAllowed)
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func sortedMoviesHandler(w http.ResponseWriter, r *http.Request) {
	switch utils.ExtractLastStringParameterFromURL(r.URL.Path) {
	case rating:
		service.GetSortedMoviesByRating(w)
	case name:
		service.GetSortedMoviesByName(w)
	case releaseDate:
		service.GetSortedMoviesByReleaseDate(w)
	default:
		utils.LOG.Error("invalid sort type", http.StatusBadRequest)
		http.Error(w, "invalid sort type", http.StatusBadRequest)
	}
}

func findMoviesHandler(w http.ResponseWriter, r *http.Request) {
	switch utils.ExtractLastStringParameterFromURL(r.URL.Path) {
	case movie:
		service.FindMoviesByMovieName(w, r)
	case actor:
		service.FindMoviesByActorName(w, r)
	default:
		utils.LOG.Error("invalid name for find", http.StatusBadRequest)
		http.Error(w, "invalid name for find", http.StatusBadRequest)
	}
}
