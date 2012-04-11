package novelty

import (
	"appengine"
	"appengine/datastore"
	"encoding/base64"
	"html/template"
	"net/http"
	"strings"
)

type Answer struct {
	Value string
}

type Password struct {
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

func authorized(r *http.Request) bool {
	h := r.Header.Get("Authorization")
	if !strings.HasPrefix(h, "Basic ") {
		return false
	}
	a, _ := base64.StdEncoding.DecodeString(strings.TrimLeft(h, "Basic "))
	fs := strings.Split(string(a), ":")
	if len(fs) != 2 {
		return false
	}
	c := appengine.NewContext(r)
	k := datastore.NewKey(c, "Password", "password", 0, nil)
	p := new(Password)
	if err := datastore.Get(c, k, p); err != nil {
		// If password is not set, seed with whatever password was passed in.
		// See: http://golang.org/misc/dashboard/app/build/key.go
		dp := Password{
			Value: fs[1],
		}
		if _, err := datastore.Put(c, k, &dp); err != nil {
			return false
		}
		return true
	}
	return p.Value == fs[1]
}

func setAnswer(answer string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !authorized(r) {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"novelty.go\"")
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
