//
// Our sample application will respond with ‘Hello World’ to any GET request.
//
 
package main

import (
   "log"
   "net/http"
   "net/http/httptest"
   "testing"
   "strings"
)

type Server struct{}

// statusHandler is an http.Handler that writes an empty response using itself
// as the response status code.
type statusHandler int

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
   w.WriteHeader(http.StatusOK)
   w.Header().Set("Content-Type", "application/json")
   w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
   s := &Server{}
   http.Handle("/", s)
   log.Fatal(http.ListenAndServe(":8080", nil))
}

func TestServeHTTP(t *testing.T) {
        s := &Server{}
  	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if b := w.Body.String(); !strings.Contains(b, "hello world") {
		t.Fatalf("body = %s, want no", b)
	}
}
