package main

import (
	"net/http"

	"github.com/definev/gosocial/internal/handler"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/err", handler.TestErrHandler)
	m.HandleFunc("/succ", handler.TestHandler)

	http.ListenAndServe(":8080", m)
}
