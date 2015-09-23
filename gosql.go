package main

import (
	"./DBReader/VideoReader"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/videos", GetAllVideos)
	http.ListenAndServe(":8000", r)
}

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videoObj := Videoreader.GetAllVideos()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(videoObj)
}
