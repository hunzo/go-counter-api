package validate

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	DefaultApiKeyHeaderName = "x-api-key"
	DefaultApiKeyQueryName  = "key"
)

func HeaderValidator(c *fiber.Ctx) error {

	apikey := os.Getenv("API_KEY")
	// fmt.Println(apikey)
	// export API_KEY=xxx
	// fmt.Printf(c.Request().Header.String())

	header := c.Get(DefaultApiKeyHeaderName)
	query := c.Query(DefaultApiKeyQueryName)

	if header == "" && query == "" {
		return fiber.ErrUnauthorized
	}

	if header != "" && header == apikey {
		return c.Next()
	}

	if query != "" && query == apikey {
		return c.Next()
	}

	return fiber.ErrUnauthorized
}
