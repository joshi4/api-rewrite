package main

import (
	"net/http"
	"os"
)

func main() {
	branch := os.Getenv("RENDER_GIT_BRANCH")
	http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		if branch == "" || branch == "master" || branch == "main" {
			w.Write([]byte("main branch"))
			return
		}
		w.Write([]byte("feature branch"))
	}))
}
