package main

import (
	"net/http"

	"github.com/Shohehe/graphql-sample/gateway/handler"
)

func main() {
	http.HandleFunc("/graphql", handler.Handler)
	http.ListenAndServe(":8000", nil)
}
