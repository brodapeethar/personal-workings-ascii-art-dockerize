package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// PageData holds information we want to send to the HTML template
type PageData struct {
	Result    string
	InputText string
}

// Helper function to render template
func renderTemplate(w http.ResponseWriter, tempPath string, data *PageData) {
	tmpl, err := template.ParseFiles(tempPath)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Check template execution error here!
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("error executing template: ", err)
		fmt.Printf("Execution error: %v\n", err)
	}
}
