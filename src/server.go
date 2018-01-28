package hello

import (
    "net/http"
    "html/template"
    "io/ioutil"
    "fmt"
    "regexp"
)

type Page struct {
    Title string
    Body  []byte
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func loadPage()

var templates = template.Must(template.ParseFiles("web/pages/edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, name string, p *Page) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    err := templates.ExecuteTemplate(w, name+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "contact", nil)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "order", nil)
}

func fillHandler(w http.ResponseWriter, r *http.Request) {
   return
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          m := validPath.FindStringSubmatch(r.URL.Path)
          if m == nil {
              http.NotFound(w, r)
              return
          }
          fn(w, r, m[2])
      }
  }


func init() {
    /* Load templates */
    http.HandleFunc("/", handler)
    http.HandleFunc("/contact", contactHandler)
    http.HandleFunc("/fill", fillHandler)
    http.HandleFunc("/order", orderHandler)
    loadTemplates()
}

