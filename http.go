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

// ListenAndServeWithLoggingTLS: Like ListenAndServeWithLogging, but with
// encryption.
func ListenAndServeWithLoggingTLS(addr, crtPath, keyPath string) {
	var err error

	if crtPath, err = ExpandPath(crtPath); err != nil {
		panic(err)
	}

	if keyPath, err = ExpandPath(keyPath); err != nil {
		panic(err)
	}

	http.ListenAndServeTLS(
		addr, crtPath, keyPath, withLogging(http.DefaultServeMux))
}

// HttpToHttps: Accept connections on port 80 and forward to the same address
// using https. This will run in an infinite loop, or panic if it can't listen.
func HttpToHttps() {
	err := http.ListenAndServe(":80", http.HandlerFunc(httpToHttpsHandler))
	if err != nil {
		panic(err)
	}
}

func httpToHttpsHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	url.Scheme = "https"
	http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
}
