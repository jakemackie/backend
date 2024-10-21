package controllers

import (
	"encoding/json"
	"net/http"

	"server/models"
)

var Videos models.Videos

func PostVideo(w http.ResponseWriter, r *http.Request) {
    var video models.Video
    err := json.NewDecoder(r.Body).Decode(&video)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    Videos = append(Videos, video)

    json.NewEncoder(w).Encode(video)
}

func AllVideos(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Videos)
}