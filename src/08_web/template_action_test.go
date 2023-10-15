package web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Pages struct {
	Title string
	Name  string
}

type Template struct {
	Title string
	Name  string
}

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Pages{
		Title: "Template Action If",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	CheckError(err)

	fmt.Println(string(body))
}

/*
Operator Perbandingan
eq artinya arg1 == arg2
ne artinya arg1 != arg2
lt artinya arg1 < arg2
le artinya arg1 <= arg2
gt artinya arg1 > arg2
ge artinya arg1 >= arg2
*/

func TemplateActionIfElse(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 50,
	})
}
func TestIfElse(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIfElse(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	CheckError(err)

	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title":   "Template Action Range",
		"Hobbies": []string{"Game", "Read", "Code"},
	})
}
func TestRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	CheckError(err)

	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(w, "address.gohtml", map[string]interface{}{
		"Name":    "Template Action Range",
		"Address": map[string]interface{}{"Street": "Jagakarsa", "Country": "Indonesia"}})
}
func TestActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	CheckError(err)

	fmt.Println(string(body))
}
