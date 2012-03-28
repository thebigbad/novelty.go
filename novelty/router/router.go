package router

import (
  "http"
  "strings"
)

func allowedMethods(r map[string]requestHandler) string {
  //TODO: Is there a cleaner way to get a comma-seperated list of keys?
  ms := make([]string, len(r))
  i := 0
  for m, _ := range r {
    ms[i] = m
    i++
  }
  return strings.Join(ms, ",")
}

type requestHandler func(w http.ResponseWriter, r *http.Request)

type Router struct {
  Routes map[string]map[string]requestHandler
}

func (r *Router) addRoute(p string, m string, h requestHandler) {
  if r.Routes == nil {
    r.Routes = make(map[string]map[string]requestHandler)
  }
  if r.Routes[p] == nil {
    r.Routes[p] = make(map[string]requestHandler)
  }
  r.Routes[p][m] = h
}

func (r *Router) Get(p string, h requestHandler) {
  r.addRoute(p, "GET", h)
}

func (router *Router) HandleRequest(w http.ResponseWriter, r *http.Request) {
  route := router.Routes[r.URL.Path]
  if route == nil {
    http.Error(w, "Not Found", http.StatusNotFound)
    return
  }
  h := route[r.Method]
  if h == nil {
    w.Header().Set("Allow", allowedMethods(route))
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
  }
  h(w, r)
}
