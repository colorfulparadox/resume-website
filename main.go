package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"resume/main/components"
	"resume/main/models"

	"github.com/a-h/templ"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var markdown goldmark.Markdown = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
	goldmark.WithExtensions(
		highlighting.NewHighlighting(
			highlighting.WithStyle("monokai"),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
			),
		),
	),
)

func unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func loadMarkDownFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read MD file")
		fmt.Println(err)
		return "<h1>404</h1>"
	}

	var buf bytes.Buffer
	if err = markdown.Convert([]byte(file), &buf); err != nil {
		fmt.Println("Error Turning MD to html")
		fmt.Println(err)
		return "<h1>404</h1>"
	}

	return buf.String()
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	var pages map[string]templ.Component = make(map[string]templ.Component)

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

	dir, _ := os.ReadDir("./project_pages")
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		var name string = file.Name()
		projectPage := components.ProjectPage(unsafe(loadMarkDownFile("./project_pages/" + name)))
		pages[name[:len(name)-3]] = components.MainHTML("Preston Resume", aboutMe, projectPage)
	}

	pages["main"] = components.MainHTML("Preston Resume", aboutMe, components.MainContent(mainContent))

	pages["404"] = components.MainHTML("Preston Resume", aboutMe, components.Page404())

	http.HandleFunc("/project/", func(w http.ResponseWriter, r *http.Request) {
		var path string = r.URL.Path[9:]
		page, exists := pages[path]
		if exists {
			page.Render(r.Context(), w)
		} else {
			pages["404"].Render(r.Context(), w)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			pages["main"].Render(r.Context(), w)
		} else {
			pages["404"].Render(r.Context(), w)
		}

	})

	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting server!")
		fmt.Println(err)
	}

}
