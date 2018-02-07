package hello

import (
    "net/http"
    "html/template"
    _"io/ioutil"
    "fmt"
    "regexp"
    "encoding/csv"
    "os"
    "log"
)

type Image struct {
    Item string
    Name string
    Price float32
    Case string
    Pack string
    Path string
}

var dataFiles =
    [...]string {
        "mousetrap",
    }


var templates = loadTemplates()
var ImageData = loadData()

func loadTemplates() map[string]*template.Template {
    pages := [...]string{"index", "contact", "catalog"}
    templates := make(map[string]*template.Template)
    for _, page := range (pages) {
        templatePath := fmt.Sprintf("./pages/%s.tmpl", page)
        templates[page] = template.Must(template.ParseFiles("pages/base.tmpl", templatePath))
    }
    return templates
}

func loadData() map[string][]*Image {

    data := make(map[string][]*Image)

    for _, fileName := range (dataFiles) {
         dataFile := fmt.Sprintf("./data/%s.csv", fileName)
         var page [30]*Image
         f, err := os.Open(dataFile)
         if err != nil {
             panic(err)
         }
         lines, err := csv.NewReader(f).ReadAll()
         if err != nil {
             panic(err)
         }
         for i, line := range lines {
             data := Image {
                 Item : line[0],
                 Name : line[1],
                 Price : 0,
                 Case : line[2],
                 Pack : line[3],
                 Path : fmt.Sprintf("./images/%s.JPG", line[0]),
             }
             page[i] = &data
         }

    }
    return data
}

func renderTemplate(w http.ResponseWriter, name string, args []*Image) {
    log.Println(name)
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

func catalogHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "catalog", ImageData["mousetrap"])
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
    http.HandleFunc("/order", orderHandler)
    http.HandleFunc("/mousetrap", catalogHandler)
}

