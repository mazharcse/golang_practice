package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		start := time.Now()
		next.ServeHTTP(w, r)

		fmt.Println("Cloud Drive", start.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, time.Since(start))
	})
}