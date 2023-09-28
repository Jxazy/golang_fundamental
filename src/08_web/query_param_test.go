package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Fajar", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(string(body))
}

func FullName(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	if firstName == "" && lastName == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Welcome, %s %s", firstName, lastName)

	}
}

func TestMultipleQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Fajar&last_name=Octhaviano", nil)
	recorder := httptest.NewRecorder()

	FullName(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(string(body))
}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, "Selamat datang, ", strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://loaclhost:8080/hello?name=Fajar&name=Ahmad&name=Octhaviano", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(string(body))
}
