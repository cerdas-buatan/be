package main

import (
    "net/http"
    "github.com/cerdas-buatan/be/module/handler"
)

func main() {
    http.HandleFunc("/chat", handler.ChatHandler)
    http.ListenAndServe(":8080", nil)
}
