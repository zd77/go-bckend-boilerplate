package api

import (
	"go-boilerplate/stores"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskStore stores.TaskStore
}

func NewTaskHandler(taskStore stores.TaskStore) *TaskHandler {
	return &TaskHandler{
		taskStore: taskStore,
	}
}

func (th *TaskHandler) HandleGetTasks(c *fiber.Ctx) error {
	list, err := th.taskStore.GetTasks()
	if err != nil {
		return err
	}
	return c.JSON(list)
}
