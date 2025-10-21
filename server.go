package main

import (
    "github.com/gofiber/fiber/v2"
    "time"
)

// This struct will be used to create our JSON object
type jsonTest struct {
    Message string `json:"message"`
    Timestamp int64 `json:"timestamp"`
}

// Like a constructor, accepts messsage as arg and generates timestamp on creation
func newJsonTest(message string) *jsonTest {
    j := jsonTest{Message: message}
    j.Timestamp = time.Now().UnixMilli()
    return &j
}

func main() {
    app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) error {
        // Convert the struct to a json object
        return c.JSON(newJsonTest("My name is Stephen"))
    })

    app.Listen(":80")
}

