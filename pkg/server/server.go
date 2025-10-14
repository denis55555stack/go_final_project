package server

import (
	"fmt"
	"net/http"
)

func Run() error {
	port := 7540
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
