package main

import (
	"api/validate"
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	REDIS_SERVER := os.Getenv("REDIS_SERVER")
	app := fiber.New()

	app.Use(cors.New())

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     REDIS_SERVER,
		Password: "",
		DB:       0,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"info":          "counter api",
			"reset_counter": "GET /api/set",
			"get_counter":   "GET /api/get",
			"incr_counter":  "GET /api/count",
		})
	})

	route := app.Group("/api", validate.HeaderValidator)

	route.Get("/set", func(c *fiber.Ctx) error {
		err := rdb.Set(ctx, "counter", 0, 0).Err()
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"info":  "reset_counter",
			"start": 0,
		})
	})

	route.Get("/get", func(c *fiber.Ctx) error {
		val, err := rdb.Get(ctx, "counter").Result()
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"value": val,
		})
	})

	route.Get("/count", func(c *fiber.Ctx) error {
		result, err := rdb.Incr(ctx, "counter").Result()
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"success": true,
			"result":  result,
		})
	})

	app.Listen(":8080")
}
