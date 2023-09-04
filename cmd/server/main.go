package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

var (
	domainOriginServer = "origin.cors.com"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	routerOriginServer := mux.NewRouter()
	routerOriginServer.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		setCookies(writer)
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

	routerPermittedCORSServer := mux.NewRouter()
	routerPermittedCORSServer.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		setCookies(writer)
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		_, err := writer.Write([]byte(`{"message": "Hello, World!"}`))
		if err != nil {
			panic(err)
		}
	}).Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodHead)

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://" + domainOriginServer, "https://" + domainOriginServer}),
		handlers.AllowedHeaders([]string{"Content-Type", "Custom-Allowed-Header"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}),
	)(routerPermittedCORSServer)

	fmt.Println("Listening permitted CORS api server on http://127.0.0.1:8081...")
	go func() {
		defer wg.Done()

		err := http.ListenAndServe(":8081", handler)
		if err != nil {
			panic(err)
		}
	}()

	routerUnPermittedCORSServer := mux.NewRouter()
	routerUnPermittedCORSServer.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		setCookies(writer)
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		_, err := writer.Write([]byte(`{"message": "Hello, World!"}`))
		if err != nil {
			panic(err)
		}
	}).Methods(http.MethodPost, http.MethodPut, http.MethodGet, http.MethodHead)

	fmt.Println("Listening unpermitted CORS api server on http://127.0.0.1:8082...")
	go func() {
		defer wg.Done()

		err := http.ListenAndServe(":8082", routerUnPermittedCORSServer)
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()
}

func setCookies(writer http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "HttpOnly",
		Value:    "v",
		HttpOnly: true,
	}
	http.SetCookie(writer, cookie)

	cookie = &http.Cookie{
		Name:   "Secure", // requires https
		Value:  "v",
		Secure: true,
	}
	http.SetCookie(writer, cookie)

	cookie = &http.Cookie{
		Name:     "SSNone",
		Value:    "v",
		SameSite: http.SameSiteNoneMode, // requires Secure
		Secure:   true,                  // requires https
	}
	http.SetCookie(writer, cookie)

	cookie = &http.Cookie{
		Name:     "SSDefault",
		Value:    "v",
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(writer, cookie)

	cookie = &http.Cookie{
		Name:     "SSLax",
		Value:    "v",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(writer, cookie)

	cookie = &http.Cookie{
		Name:     "SSStrict",
		Value:    "v",
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(writer, cookie)

}
