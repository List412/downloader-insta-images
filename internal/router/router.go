package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"instagram-downloader/internal/instagram"
	"log"
	"net/http"
)

func NewRouter() {
	router := mux.NewRouter()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router.Use(MiddlewareJson)
	router.Use(PanicHandler)
	router.Use(LogRequest)

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", health).Methods(http.MethodGet)

	api.HandleFunc("/upload", upload)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func health(w http.ResponseWriter, r *http.Request) {
	// dont check anything now but should be
	err := json.NewEncoder(w).Encode(map[string]string{
		"im": "ok",
	})
	if err != nil {
		panic(err)
	}
}

type uploadJson struct {
	Url  string
	Path string
}

func upload(w http.ResponseWriter, r *http.Request) {
	data := uploadJson{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}

	go instagram.TakeImages(data.Url)
}
