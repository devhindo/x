package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, x server 👋!")
    })

    http.ListenAndServe(":3000", nil)
}

/*
package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func main() {
	server := fiber.New()
	server.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello, xCLI 👋!")
        return c.SendString("Hello, World 👋!")
    })

    server.Listen(":3000")
}
*/