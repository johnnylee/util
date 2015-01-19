package util

import (
	"net/http"
)

var httpLogger = NewPrefixLogger("http")

func withLogging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpLogger.Msg("%s %s %s %s",
			r.RemoteAddr, r.Method, r.URL, r.Referer())
		handler.ServeHTTP(w, r)
	})
}

// ListenAndServeWithLogging: A wrapper for the standard library's
// http.ListenAndServe function, logging requests using the standard logger.
// Uses http.DefaultServeMux.
func ListenAndServeWithLogging(addr string) {
	http.ListenAndServe(addr, withLogging(http.DefaultServeMux))
}
