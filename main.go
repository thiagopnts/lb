package main

import (
	"lb"
	"net/http"
)

func main() {
  s, _ := lb.NewLoadBalancer("http://74.125.234.232", "http://74.125.234.233", "http://74.125.234.230")

  http.Handle("/", s)
  http.ListenAndServe(":8000", nil)
}
