package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/video/add/{filename}", handleAddVideo).Methods("GET")
	router.HandleFunc("/video/{id}", handleGetVideo).Methods("GET")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func handleAddVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	video := uploadVideo(filename)

	json.NewEncoder(w).Encode(map[string]string{"message": "Video uploaded with id: " + video.Id})

	//v := videoInfo("P2f_PyPrxgY")
	//json.NewEncoder(w).Encode(map[string]string{"message": strconv.Itoa(int((v.FileDetails.DurationMs / 1000)))})
}

func handleGetVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	v := videoInfo(id)
	json.NewEncoder(w).Encode(map[string]string{
		"message": strconv.Itoa(int((v.FileDetails.DurationMs / 1000))),
		"a":       v.ContentDetails.Duration,
		"b":       v.ContentDetails.Dimension,
		"c":       v.Player.EmbedHtml,
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}
