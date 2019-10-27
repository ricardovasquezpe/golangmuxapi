package middle

import (
	"log"
	"net/http"
)

func MiddlewareOne(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			log.Println("middleware one")
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}
