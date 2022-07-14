package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	app := fiber.New()

	// | METHOD | REQUEST                               | RESPONSE                      |
	// |--------|---------------------------------------|-------------------------------|
	// | GET    | `/name/{PARAM}`                       | body: `Hello, PARAM!`         |
	// | GET    | `/bad`                                | Status: `500`                 |
	// | POST   | `/data` + Body `PARAM`                | body: `I got message:\nPARAM` |
	// | POST   | `/headers`+ Headers{"a":"2", "b":"3"} | Header `"a+b": "5"`           |
	app.Get("/name/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString(fmt.Sprintf("Hello, %s!", name))
	})
	app.Get("/bad", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusInternalServerError)
	})
	app.Post("/data", func(c *fiber.Ctx) error {
		body := c.Body()
		return c.SendString(fmt.Sprintf("I got message:\n%s", body))
	})
	app.Post("/headers", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		a, foundA := headers["A"]
		b, foundB := headers["B"]
		if !foundA || !foundB {
			return c.SendStatus(http.StatusInternalServerError)
		}
		aInt, convErrA := strconv.Atoi(a)
		bInt, convErrB := strconv.Atoi(b)
		if convErrA != nil || convErrB != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		c.Set("a+b", strconv.Itoa(aInt+bInt))
		for key, value := range headers {
			c.Set(key, value)
		}
		return c.SendStatus(http.StatusOK)
	})

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))

	if err := app.Listen(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
