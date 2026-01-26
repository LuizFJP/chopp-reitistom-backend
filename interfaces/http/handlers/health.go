package handlers

import (
	"fmt"
	"net/http"
)

func BuildLivenessHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}

	return http.HandlerFunc(fn)
}
