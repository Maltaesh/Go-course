package main

import (
	"fmt"
	"net/http"
	"time"
)

func LogRequestInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("Date: %d/%d/%d, Time: %d:%d, ", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute())
		fmt.Println("URL:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
