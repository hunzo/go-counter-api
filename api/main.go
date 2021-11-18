package main

import (
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
			"info":                 "counter api",
			"set_your_counter_key": "GET /api/set?key=yourKey",
			"get_counter":          "GET /api/get?key=yourKey",
			"incr_counter":         "GET /api/inc?key=yourKey",
			"decr_counter":         "GET /api/dec?key=yourKey",
		})
	})

	route := app.Group("/api")

	route.Get("/set", func(c *fiber.Ctx) error {
		key := c.Query("key")

		if key == "" {
			return c.JSON(fiber.ErrBadRequest)
		}

		err := rdb.Set(ctx, key, 0, 0).Err()
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"info":  "call with ?key=yourkey",
			"start": 0,
		})
	})

	route.Get("/get", func(c *fiber.Ctx) error {
		key := c.Query("key")
		if key == "" {
			return c.JSON(fiber.ErrBadRequest)
		}
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"value": val,
		})
	})

	route.Get("/inc", func(c *fiber.Ctx) error {
		key := c.Query("key")
		if key == "" {
			return c.JSON(fiber.ErrBadRequest)
		}

		result, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"success": true,
			"result":  result,
		})
	})

	route.Get("/dec", func(c *fiber.Ctx) error {
		key := c.Query("key")
		if key == "" {
			return c.JSON(fiber.ErrBadRequest)
		}

		result, err := rdb.Decr(ctx, key).Result()
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
