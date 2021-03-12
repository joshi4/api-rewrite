package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	branch := os.Getenv("RENDER_GIT_BRANCH")
	ex := os.Getenv("RENDER_EXTERNAL_URL")
	http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("external_url is: %s", ex)))
		if branch == "" || branch == "master" || branch == "main" {
			w.Write([]byte("main branch\n"))
			return
		}
		w.Write([]byte("feature branch"))
	}))
}
