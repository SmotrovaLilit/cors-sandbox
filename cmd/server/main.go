package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	routerOriginServer := mux.NewRouter()
	routerOriginServer.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte(`{"message": "Hello, World!"}`))
		if err != nil {
			panic(err)
		}
	}).Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodHead)
	routerOriginServer.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	fmt.Println("Listening origin server with static scripts on http://127.0.0.1:3000...")
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(":3000", routerOriginServer)
		if err != nil {
			panic(err)
		}
	}()

	routerAnotherServer := mux.NewRouter()
	routerAnotherServer.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte(`{"message": "Hello, World!"}`))
		if err != nil {
			panic(err)
		}
	}).Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodHead)

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://127.0.0.1:3000"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}),
	)(routerAnotherServer)

	fmt.Println("Listening another api server on http://127.0.0.1:8081...")
	go func() {
		defer wg.Done()

		err := http.ListenAndServe(":8081", handler)
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}
