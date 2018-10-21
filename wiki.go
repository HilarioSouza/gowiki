package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

<<<<<<< HEAD
//Page estrutura da página da wiki
=======
//Page : Estrutura da página
>>>>>>> 1b30a8bc6b9bf2b6904579d917c6ea4cbbd956cf
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<h1>%s</h'><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name =\"body\">%s</textarea><br>"+
		"<input type=\"sumit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}
