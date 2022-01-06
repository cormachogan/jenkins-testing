//
// Since we are building a CD pipeline, we should have some tests in place
// Our code is so simple that it only needs one test case; 
// ensuring that we receive the correct string when we hit the root URL
//

package main

import (
   "log"
   "net/http"
)

type Server struct{}

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

