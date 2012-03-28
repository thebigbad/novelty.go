package router

import (
  "bytes"
  "http"
  "http/httptest"
  "testing"
)

func Expect(t *testing.T, e, a interface{}) {
  if a != e {
    t.Errorf("expected %s got %s", e, a)
  }
}

func TestAllowedMethods(t *testing.T) {
  handler := func(w http.ResponseWriter, r *http.Request) {}
  route := map[string]requestHandler{}

  route["bees"] = handler
  Expect(t, "bees", allowedMethods(route))
  route["birds"] = handler
  Expect(t, "birds,bees", allowedMethods(route))
}

func TestNotFound(t *testing.T) {
  req, _ := http.NewRequest("DELETE", "/bees", nil)
  res := httptest.NewRecorder()
  router := Router{}
  router.HandleRequest(res, req)

  Expect(t, 404, res.Code)
  Expect(t, "Not Found\n", res.Body.String())
}

func TestMethodNotAllowed(t *testing.T) {
  req, _ := http.NewRequest("DELETE", "/bees", nil)
  res := httptest.NewRecorder()
  handler := func(w http.ResponseWriter, r *http.Request) {}
  router := Router{}
  router.Get("/bees", handler)
  router.HandleRequest(res, req)

  Expect(t, 405, res.Code)
  Expect(t, "GET", res.Header().Get("Allow"))
  Expect(t, "Method Not Allowed\n", res.Body.String())
}

func TestGet(t *testing.T) {
  req, _ := http.NewRequest("GET", "/bees", nil)
  res := httptest.NewRecorder()
  handler := func(w http.ResponseWriter, r *http.Request) {
    bytes.NewBufferString("bees!").WriteTo(w)
  }
  router := Router{}
  router.Get("/bees", handler)
  router.HandleRequest(res, req)

  Expect(t, 200, res.Code)
  Expect(t, "bees!", res.Body.String())
}
