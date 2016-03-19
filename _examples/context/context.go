package main

import (
	"fmt"
	"gopkg.in/vinci-proxy/context.v0"
	"gopkg.in/vinci-proxy/vinci.v0"
	"net/http"
)

func main() {
	fmt.Printf("Server listening on port: %d\n", 3100)
	vs := vinci.NewServer(vinci.ServerOptions{Host: "localhost", Port: 3100})

	vs.Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
		context.Set(r, "foo", "bar")
		h.ServeHTTP(w, r)
	})

	vs.Use(func(w http.ResponseWriter, r *http.Request, h http.Handler) {
		w.Header().Set("foo", context.GetString(r, "foo"))
		h.ServeHTTP(w, r)
	})

	vs.Forward("http://httpbin.org")

	err := vs.Listen()
	if err != nil {
		fmt.Errorf("Error: %s\n", err)
	}
}
