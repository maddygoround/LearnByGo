package main

import (
	"./DBConnection"
	"./DBReader/VideoReader"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	SQLCtx.InitDB()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/videos", GetAllVideos)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./WebApp/")))
	http.Handle("/", r)
	//fmt.Println("server started on : 8080", http.ListenAndServe(":8000", nil))
}

func GetAllVideos(w http.ResponseWriter, r *http.Request) {
	videoObj := Videoreader.GetAllVideos()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	json.NewEncoder(gz).Encode(videoObj)
}
