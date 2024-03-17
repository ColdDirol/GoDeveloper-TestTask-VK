package service

import (
	"GoDeveloperVK-testTask/models"
	"GoDeveloperVK-testTask/server/repository"
	"GoDeveloperVK-testTask/utils"
	"GoDeveloperVK-testTask/utils/logger"
	"encoding/json"
	"net/http"
)

func GetActors(w http.ResponseWriter) {
	actors, err := repository.GetAllActors()
	if err != nil {
		utils.LOG.Error("failed to get actors", logger.Err(err))
		http.Error(w, "failed to get actors", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(actors)
	if err != nil {
		utils.LOG.Error("failed to marshal actors data", logger.Err(err))
		http.Error(w, "failed to marshal actors data", http.StatusInternalServerError)
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

func GetActorByID(w http.ResponseWriter, actorID int) {
	actorWithMovies, err := repository.GetActorByID(actorID)
	if err != nil {
		utils.LOG.Error("failed to get actor", logger.Err(err))
		http.Error(w, "failed to get actor", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(actorWithMovies)
	if err != nil {
		utils.LOG.Error("failed to marshal actors data", logger.Err(err))
		http.Error(w, "failed to marshal actors data", http.StatusInternalServerError)
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

func PostActors(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		utils.LOG.Error("failed to decode actor data", logger.Err(err))
		http.Error(w, "failed to decode actor data", http.StatusInternalServerError)
		return
	}

	err = repository.AddActor(actor)
	if err != nil {
		utils.LOG.Error("failed to insert actor", logger.Err(err))
		http.Error(w, "failed to insert actor", http.StatusBadRequest)
		return
	}
}

func UpdateActor(w http.ResponseWriter, r *http.Request, actorID int) {
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		utils.LOG.Error("failed to decode actor data", logger.Err(err))
		http.Error(w, "failed to decode actor data", http.StatusInternalServerError)
		return
	}

	err = repository.UpdateActor(actorID, actor)
	if err != nil {
		utils.LOG.Error("failed to update actor", logger.Err(err))
		http.Error(w, "failed to update actor", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteActor(w http.ResponseWriter, actorID int) {
	err := repository.DeleteActorByID(actorID)
	if err != nil {
		utils.LOG.Error("failed to delete actor", logger.Err(err))
		http.Error(w, "failed to delete actor", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
