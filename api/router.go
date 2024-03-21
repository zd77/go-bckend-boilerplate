package api

import (
	"fmt"
	"go-boilerplate/api/middleware"
	"go-boilerplate/stores"

	"github.com/gofiber/fiber/v2"
)

type RouterDependencies struct {
	TaskStore stores.TaskStore
	UserStore stores.UserStore
}

func RouterApi(app *fiber.App, dependencies RouterDependencies) {
	fmt.Println("Configurando rutas principales")
	apiGroup := app.Group("/api", middleware.AuthMiddleware)
	configureTaskRoutes(apiGroup, dependencies.TaskStore)
	configureUserRoutes(apiGroup, dependencies.UserStore)
}

func configureTaskRoutes(apiGroup fiber.Router, taskStore stores.TaskStore) {
	fmt.Println("Configurando rutas de tareas")
	taskHandler := NewTaskHandler(taskStore)
	apiGroup.Get("/tasks", taskHandler.HandleGetTasks)
}

func configureUserRoutes(apiGroup fiber.Router, userStore stores.UserStore) {
	fmt.Println("Configurando rutas de usuarios")
	taskHandler := NewUserHandler(userStore)
	apiGroup.Get("/users", taskHandler.HandlerGetUsers)
}
