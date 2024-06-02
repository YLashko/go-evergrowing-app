package main_page

import (
	"html/template"
	"net/http"
)

type MainPageData struct {
	Welcoming_message string
	Paragraphs        []MainPageParagraph
}

type MainPageParagraph struct {
	Title    string
	Contents string
}

func Template() *template.Template {
	return template.Must(template.ParseFiles("views/main.html"))
}

func Main_page(templ *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := MainPageData{
			Welcoming_message: "Hello! This is the evergrowing web app made using golang. It grows as I learn new golang features.",
			Paragraphs: []MainPageParagraph{
				{
					Title:    "This is a sample paragraph",
					Contents: "And this is its contents",
				},
				{
					Title:    "This is a sample 2 paragraph",
					Contents: "And this is its contents",
				},
			},
		}
		templ.Execute(w, data)
	}
}
