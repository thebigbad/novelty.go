package hello

import (
  "http"
  "strings"
  "template"
)

var question = "Is Ryan at the Office?"
var answer = "yes"
var rootTemplate = template.Must(template.New("").ParseFile("index.template"))

type Context struct{
  Question string
  Answer string
}

type RequestHandler func(w http.ResponseWriter, r *http.Request)

var routes = make(map[string]map[string]RequestHandler)

func addRoute(path string, method string, handler RequestHandler) {
  if routes[path] == nil {
    routes[path] = make(map[string]RequestHandler)
  }
  routes[path][method] = handler
}

func router(w http.ResponseWriter, r *http.Request) {
  route := routes[r.URL.Path]
  if route == nil {
    http.Error(w, "Not Found", http.StatusNotFound)
    return
  }
  handler := route[r.Method]
  if handler == nil {
    //TODO: Is there a cleaner way to get a comma-seperated list of keys?
    methods := make([]string, len(route))
    i := 0
    for m, _ := range route {
      methods[i] = m
      i++
    }
    w.Header().Set("Allow", strings.Join(methods, ","))
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
  }
  handler(w, r)
}

func init() {
  addRoute("/", "GET", func(w http.ResponseWriter, r *http.Request) {
    // TODO: Is there a way to skip the type and use an inline struct?
    context := Context{ Question: question, Answer: answer }
    err := rootTemplate.Execute(w, context)
    if err != nil {
      http.Error(w, err.String(), http.StatusInternalServerError)
    }
  })

  addRoute("/yes", "GET", func(w http.ResponseWriter, r *http.Request) {
    answer = "yes"
    http.Redirect(w, r, "", http.StatusFound)
  })

  addRoute("/no", "GET", func(w http.ResponseWriter, r *http.Request) {
    answer = "no"
    http.Redirect(w, r, "", http.StatusFound)
  })

  http.HandleFunc("/", router)
}
