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
func ListenAndServeWithLogging(addr string, handler http.Handler) error {
	if handler == nil {
		handler = http.DefaultServeMux
	}
	return http.ListenAndServe(addr, withLogging(handler))
}

// ListenAndServeTLSWithLogging: Like ListenAndServeWithLogging, but with
// encryption.
func ListenAndServeTLSWithLogging(
	addr string, handler http.Handler, crtPath, keyPath string) error {

	if handler == nil {
		handler = http.DefaultServeMux
	}

	var err error

	if crtPath, err = ExpandPath(crtPath); err != nil {
		panic(err)
	}

	if keyPath, err = ExpandPath(keyPath); err != nil {
		panic(err)
	}

	return http.ListenAndServeTLS(addr, crtPath, keyPath, withLogging(handler))
}

// HttpToHttps: Accept connections on port 80 and forward to the same host
// using https. This will run in an infinite loop, or panic if it can't listen.
func HttpToHttps() {
	err := http.ListenAndServe(":80", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + r.Host + r.URL.RequestURI()
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		}))
	if err != nil {
		panic(err)
	}
}
