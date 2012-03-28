package novelty

import (
  "http"
  "novelty/router"
  "template"
)

var answer = "yes"
var rootTemplate = template.Must(template.New("").ParseFile("index.template"))

type Context struct{
  Answer string
}

func init() {
  router := router.Router{}
  router.Get("/", func(w http.ResponseWriter, r *http.Request) {
    // TODO: Is there a way to skip the type and use an inline struct?
    context := Context{ Answer: answer }
    err := rootTemplate.Execute(w, context)
    if err != nil {
      http.Error(w, err.String(), http.StatusInternalServerError)
    }
  })

  router.Get("/yes", func(w http.ResponseWriter, r *http.Request) {
    answer = "yes"
    http.Redirect(w, r, "/", http.StatusFound)
  })

  router.Get("/no", func(w http.ResponseWriter, r *http.Request) {
    answer = "no"
    http.Redirect(w, r, "/", http.StatusFound)
  })

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    router.HandleRequest(w, r)
  })
}
