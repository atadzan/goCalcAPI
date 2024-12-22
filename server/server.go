package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atadzan/goCalcAPI/pkg/handlers"
)

func Run(handlerInstance *handlers.Handler, port string) (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handlerInstance.Calculate)

	handler := validateHTTPMethodMiddleware(mux)

	log.Printf("GoCalcAPI is listening on port %s...", port)
	if err = http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("can't run server. Err: %v", err)
	}
	return
}

func validateHTTPMethodMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err := fmt.Fprintln(w, "Method Not Allowed")
			if err != nil {
				log.Println(err)
				return
			}
		}
	})
}
