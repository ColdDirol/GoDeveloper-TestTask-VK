package handler

import (
	"GoDeveloperVK-testTask/auth/middleware"
	"GoDeveloperVK-testTask/server/service"
	"GoDeveloperVK-testTask/utils"
	"GoDeveloperVK-testTask/utils/logger"
	"net/http"
)

func InitActorsHandlers() {
	http.HandleFunc("/actors", middleware.Middleware(actorsHandler))
	http.HandleFunc("/actors/", middleware.Middleware(actorByIDHandler))
}

func actorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		service.GetActors(w)
	case http.MethodPost:
		service.PostActors(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func actorByIDHandler(w http.ResponseWriter, r *http.Request) {
	actorID, err := utils.ExtractIDFromURL(r.URL.Path)
	if err != nil {
		utils.LOG.Error("invalid id", logger.Err(err))
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		service.GetActorByID(w, actorID)
	case http.MethodPut:
		service.UpdateActor(w, r, actorID)
	case http.MethodDelete:
		service.DeleteActor(w, actorID)
	default:
		utils.LOG.Error("invalid http method", http.StatusMethodNotAllowed)
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}
