package util

import (
	"fmt"
	"net/http"
)

//Middleware for format response json data
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Println(r.RequestURI)
	})
}
