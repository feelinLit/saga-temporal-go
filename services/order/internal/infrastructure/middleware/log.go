package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type LogMiddleware struct {
	next http.Handler
}

func (mw *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] received request %s %s\n",
		time.Now().Format("02-01-2006 15:04:05"),
		r.Method,
		r.URL)

	mw.next.ServeHTTP(w, r)
}

func WithLog(handler http.Handler) *LogMiddleware {
	return &LogMiddleware{next: handler}
}
