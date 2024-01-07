package auth

import (
	"github.com/gofiber/fiber"
	auth "github.com/roronoazor/InventoryBackendQL/auth/models"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	user := new(auth.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

}
