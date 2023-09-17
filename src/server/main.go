package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // fmt.Fprintf(w, "Hello, x server 👋!")
        // attach 
    })

    // handle GET requests on /hello
    

    http.ListenAndServe(":8080", nil)
}


