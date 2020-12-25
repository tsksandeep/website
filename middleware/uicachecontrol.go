package middleware

import (
	"fmt"
	"net/http"
	"regexp"
)

const (
	cacheControl        = "Cache-Control"
	cacheControlNoCache = "no-cache"
	cacheControlNoStore = "no-store"
)

var (
	noCachePaths = []string{"/index.html"}
	cachePaths   = regexp.MustCompile(`^/.*\.(css|js|json|png|jpg|jpeg|ico|svg|ttf|woff|woff2)$`)
)

func UICacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cacheControlValue := fmt.Sprintf("%s, %s", cacheControlNoCache, cacheControlNoStore)
		w.Header().Set(cacheControl, cacheControlValue)
		next.ServeHTTP(w, r)
	})
}

func shouldCache(path string) bool {
	for _, f := range noCachePaths {
		if path == f {
			return false
		}
	}
	return cachePaths.MatchString(path)
}