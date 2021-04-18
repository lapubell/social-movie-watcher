package main

import "net/http"

func noCacheStaticAssets(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-store, max-age=0")
		h.ServeHTTP(w, r)
	}
}
