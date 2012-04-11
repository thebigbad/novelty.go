package novelty

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"net/http"
)

type Answer struct {
	Value string
}

func init() {
	http.HandleFunc("/", getAnswer)
	http.HandleFunc("/yes", setAnswer("yes"))
	http.HandleFunc("/no", setAnswer("no"))
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
