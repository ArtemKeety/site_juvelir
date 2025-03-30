package paigesrander

import (
	"fmt"
	"html/template"
	"net/http"
)

//func Render(w http.ResponseWriter, temp string, data interface{}) {
//	tmp, _ := template.ParseFiles(temp)
//	tmp.Execute(w, data)
//}

func Render(w http.ResponseWriter, temp string, data interface{}) {
	tmp, err := template.ParseFiles(temp)
	if err != nil {
		// Log the error and return an internal server error
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, data)
	if err != nil {
		// Log the error and return an internal server error
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}
