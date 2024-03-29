package web

import (
	"fmt"
	"net/http"
	"testing"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	CheckError(err)
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ini Homepage")
	})

	mux.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ini Students Page")
	})

	mux.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ini Teacher Page")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	CheckError(err)
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	CheckError(err)
}
