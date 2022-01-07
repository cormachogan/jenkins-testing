//
// Our sample application will respond with ‘Hello World’ to any GET request.
//
 
package main

import (
   "net/http"
   "net/http/httptest"
   "testing"
   "strings"
)


func TestServeHTTP(t *testing.T) {
        s := &Server{}
  	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if b := w.Body.String(); !strings.Contains(b, "hello world") {
		t.Fatalf("body = %s, want hello world", b)
	}
}
