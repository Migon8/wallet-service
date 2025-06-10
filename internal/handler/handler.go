package handler

import (
	"fmt"
	"net/http"
)

func InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Wallet Service is running!")
	})

	return mux
}
