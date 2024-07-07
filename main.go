package main

import (
    "net/http"
    "github.com/cerdas-buatan/be/module"
)

func main() {
    http.HandleFunc("/chat", handler.ChatHandler)
    http.ListenAndServe(":8080", nil)
}
