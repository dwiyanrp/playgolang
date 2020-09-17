package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber"
)

var count = 0

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) {
	  count++
	  log.Println(count)

	  time.Sleep(time.Second)

    c.Send("Hello, World 👋!")
  })

  app.Listen(3000)
}