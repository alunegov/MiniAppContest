package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Setup new HTTP server mux to handle different paths
	mux := httprouter.New()

	mux.HandleOPTIONS = true
	mux.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r)

		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Headers", "Content-Type, Ngrok-Skip-Browser-Warning")
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	// This serves
	mux.HandlerFunc(http.MethodGet, "/goods", goods())
	// This serves
	mux.HandlerFunc(http.MethodPost, "/order", order())

	server := http.Server{
		Addr:    ":4001",
		Handler: mux,
	}

	// Start the webserver
	if err := server.ListenAndServe(); err != nil {
		panic("failed to listen and serve: " + err.Error())
	}
}

func goods() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)

		w.Header().Set("Access-Control-Allow-Origin", "*")

		g := []struct {
			Id     int    `json:"id"`
			Name   string `json:"name"`
			Price  int    `json:"price"`
			Pic    string `json:"pic"`
			PicAlt string `json:"picAlt"`
		}{
			{1, "Pen", 10, "https://miniappcontest.work.gd/images/pen.svg", "Pen"},
			{2, "Pineapple", 20, "https://miniappcontest.work.gd/images/pineapple.svg", "Pineapple"},
			{3, "Apple", 30, "https://miniappcontest.work.gd/images/apple.svg", "Apple"},
		}

		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func order() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)

		w.Header().Set("Access-Control-Allow-Origin", "*")

		var o []struct {
			Id  int `json:"id"`
			Qty int `json:"qty"`
		}

		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(o)

		// TODO: persist order

		w.WriteHeader(http.StatusNoContent)
	}
}
