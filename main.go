package main

import (
	artGenerator "ascii-art-web-stylize/generator"
	"ascii-art-web-stylize/handlers"
	"fmt"
	"log"
	"net/http"
)

var port = 3000

func main() {
	// This tells Go: "If a request starts with /static/, look inside the ./static folder"
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 1. Routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.FormHandler)

	welcomeText := "Ascii art web!"
	banner := "./banners/standard.txt"
	welcomeTextArt, err := artGenerator.GenerateArt(welcomeText, banner)

	if err != nil {
		// handle art generation error
		fmt.Printf("Error occured: %v\n", err)
		return
	}

	// welcomeTextArt to be displayed on server start
	fmt.Print(welcomeTextArt)

	fmt.Printf("port running on http://localhost:%v/\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}
}
