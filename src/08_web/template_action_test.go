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
