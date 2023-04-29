package handler

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"post.loc/app/pkg/model"
)

func Home(w http.ResponseWriter, r *http.Request) {

	posts := model.GetPosts()

	files := getTemplateString("home")

	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, posts)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	model.Delete(id)
	http.Redirect(w, r, "/list", 301)
}

func Post(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	post := model.GetPost(id)
	files := getTemplateString("post")
	ts, err := template.ParseFiles(files...)
	errorHandler(err, w)
	err = ts.Execute(w, post)
	errorHandler(err, w)
}

func List(w http.ResponseWriter, r *http.Request) {

	posts := model.GetPosts()
	files := getTemplateString("list")

	tmpl, _ := template.ParseFiles(files...)
	tmpl.Execute(w, posts)

}

func Create(w http.ResponseWriter, r *http.Request) {

	files := getTemplateString("create")

	ts, err := template.ParseFiles(files...)
	errorHandler(err, w)

	err = ts.Execute(w, nil)
	errorHandler(err, w)
}

func Redact(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	post := model.GetPost(id)

	files := getTemplateString("redact")
	ts, err := template.ParseFiles(files...)
	errorHandler(err, w)

	err = ts.Execute(w, post)
	errorHandler(err, w)
}

func Update(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	text := r.FormValue("text")
	id := r.FormValue("id")

	model.Update(id, title, text)

	http.Redirect(w, r, "/list", 301)
}

func Store(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}

	filename := fileSave(w, r)
	title := r.FormValue("title")
	text := r.FormValue("text")

	model.Store(title, text, filename)

	http.Redirect(w, r, "/list", 301)
}

func fileSave(w http.ResponseWriter, r *http.Request) string {

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")

	errorHandler(err, w)

	defer file.Close()

	f, err := os.OpenFile("../../ui/image/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	errorHandler(err, w)

	filename := handler.Filename

	defer f.Close()

	io.Copy(f, file)

	return filename
}

func getTemplateString(page string) []string {

	return []string{
		"../../ui/html/pages/" + page + ".page.tmpl",
		"../../ui/html/layout/base.layout.tmpl",
		"../../ui/html/partials/footer.partial.tmpl",
	}
}

func errorHandler(err error, w http.ResponseWriter) {

	if err != nil {
		fmt.Println(err.Error())
		//http.Error(w, "Internal Server Error", 500)
	}
}
