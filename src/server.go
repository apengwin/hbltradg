package hello

import (
    "net/http"
    "html/template"
    "io/ioutil"
    "fmt"
    "regexp"
)

type Image struct {
    Name string
    Path string
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func loadTemplates() map[string]*template.Template{
    pages := ["index", "contact", "locks"]
    templates := make(map[string]*template.Template)
    for _, page := range (pages) {
        templatePath := fmt.Sprintf("./pages/%s.tmpl", page)
        templates[page] = template.Must(template.ParseFiles("pages/base.tmpl", templatePath))
    }
    return templates
}

var templates = loadTemplates()

func renderTemplate(w http.ResponseWriter, name string, args []Image) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    err := templates[name].ExecuteTemplate(w, "base", args)
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
    http.HandleFunc("/", handler)
    http.HandleFunc("/contact", contactHandler)
    http.HandleFunc("/fill", fillHandler)
    http.HandleFunc("/order", orderHandler)
}

