package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "nama-lengkap"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	_, err := fmt.Fprint(w, "Create Cookie Successful")
	CheckError(err)
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("nama-lengkap")
	if err != nil {
		_, err := fmt.Fprint(w, "No Cookies")
		CheckError(err)

	} else {
		name := cookie.Value
		_, err := fmt.Fprintf(w, "Hello %s", name)
		CheckError(err)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	CheckError(err)

}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Fajar", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookies %s: %s \n", cookie.Name, cookie.Value)

	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "nama-lengkap"
	cookie.Value = "Fajar"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	CheckError(err)
	fmt.Println(string(body))
}
