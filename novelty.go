package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"html/template"
	"net/http"
	"os"
)

type Answer struct {
	Value string
}

type Password struct {
	Value string
}

func main() {
	http.HandleFunc("/", getAnswer)
	http.HandleFunc("/yes", setAnswer("yes"))
	http.HandleFunc("/no", setAnswer("no"))
	appengine.Main()
}

func getAnswer(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	k := datastore.NewKey(c, "Answer", "answer", 0, nil)
	a := new(Answer)
	if err := datastore.Get(c, k, a); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := template.Must(template.ParseFiles("index.template"))
	if err := t.Execute(w, a); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setAnswer(answer string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, p, ok := r.BasicAuth(); !ok || p != os.Getenv("PASSWORD") {
			w.Header().Set("WWW-Authenticate", "Basic")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		c := appengine.NewContext(r)
		k := datastore.NewKey(c, "Answer", "answer", 0, nil)
		a := Answer{
			Value: answer,
		}
		if _, err := datastore.Put(c, k, &a); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
