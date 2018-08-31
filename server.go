package main

import(
	"html/template"
	"net/http"
	//Uncomment for Verbose logging
	//"fmt"
)

/*
TYPE : Page
struct for use with HTTP/TEMPLATE to display web pages.  Webpages internal data is stored here.
*/
type Page struct {
	Title string
	Body  string
}

func loadPage(title string) (*Page, error){
	return &Page{Title: title, Body: "blank..."}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage("Test")
  renderTemplate(w, "./html/index", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  t, _ := template.ParseFiles(tmpl + ".html")
  t.Execute(w, p)
}

func main() {
  http.HandleFunc("/", viewHandler)
	http.Handle("/css/",http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/fonts/",http.StripPrefix("/fonts/", http.FileServer(http.Dir("./fonts"))))
	http.Handle("/js/",http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.Handle("/vendor/",http.StripPrefix("/vendor/", http.FileServer(http.Dir("./vendor"))))
	http.Handle("/img/",http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	http.ListenAndServe(":8083", nil)
}