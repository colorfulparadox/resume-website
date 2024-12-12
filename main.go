package main

import (
	"fmt"
	"net/http"
	"resume/main/components"
	"resume/main/models"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/project/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		http.ServeFile(w, r, "project.html")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		aboutMe, err := models.LoadJSON[models.AboutMe]("data/aboutme.json")
		if err != nil {
			fmt.Println("Error loading json file!")
			fmt.Println(err)
		}

		mainContent, err := models.LoadJSON[models.MainContent]("data/maincontent.json")
		if err != nil {
			fmt.Println("Error loading json file!")
			fmt.Println(err)
		}

		webpage := components.MainHTML("Preston Resume", aboutMe, mainContent)
		webpage.Render(r.Context(), w)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server!")
		fmt.Println(err)
	}

}
