package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is empty")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}

}

func ResponseCode2(w http.ResponseWriter, r *http.Request) {
	jurusan := r.URL.Query().Get("jurusan")

	if jurusan == "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Jurusan is not found")
	} else {
		fmt.Fprintf(w, "Jurusan kamu : %s", jurusan)
	}
}

func TestResponseCodeError(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Fajar", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseJurusan(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseCode2(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
