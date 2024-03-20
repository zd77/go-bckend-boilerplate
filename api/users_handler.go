package api

import (
	"go-boilerplate/stores"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore stores.UserStore
}

func NewUserHandler(userStore stores.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (uh *UserHandler) HandlerGetUsers(c *fiber.Ctx) error {
	list, err := uh.userStore.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(list)
}
