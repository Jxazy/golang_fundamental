package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"title": "Template Data Map",
		"name":  "Fajar",
		"Address": map[string]interface{}{
			"street": "jalan belum ada map"},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	CheckError(err)
	fmt.Println(string(body))
}

type Address struct {
	street string
}

type Page struct {
	title   string
	name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		title: "Template Data Struct",
		name:  "Eko",
		Address: Address{
			street: "Jalan Belum Ada",
		},
	})
}

func TestDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	CheckError(err)
	fmt.Println(string(body))

}
