package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(string(body))

}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("x-powered-by", "Fajar Octhaviano")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)

	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-powered-by"))
}
