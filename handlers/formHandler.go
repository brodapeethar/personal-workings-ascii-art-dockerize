package handlers

import (
	artGenerator "ascii-art-web-stylize/generator"
	"fmt"
	"net/http"
	"os"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get data from the form
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Validate input (400 Bad Request)
	if text == "" || banner == "" {
		http.Error(w, "400 Bad Request: Missing text or banner selection", http.StatusBadRequest)
		return
	}

	// Check if the banner file exists
	bannerPath := fmt.Sprintf("./banners/%v.txt", banner)
	if _, err := os.Stat(bannerPath); os.IsNotExist(err) {
		http.Error(w, "404 Banner Not Found", http.StatusNotFound)
		return
	}

	// generating the ASCII art.
	asciiArtResult, err := artGenerator.GenerateArt(text, bannerPath)

	if err != nil {
		// handle art generation error
		fmt.Printf("Error occured: %v\n", err)
		ee := fmt.Sprintf("%v", err)
		http.Error(w, ee, http.StatusInternalServerError)
		return
	}

	// Sending result back to home page
	renderTemplate(w, "./templates/index.html", &PageData{Result: asciiArtResult, InputText: text})
}
