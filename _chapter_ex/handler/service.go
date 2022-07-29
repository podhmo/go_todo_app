package handler

import (
	"context"

	"github.com/budougumi0617/go_todo_app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService ChangeTaskStatusService RegisterUserService
type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type ChangeTaskStatusService interface {
	ChangeTaskStatus(ctx context.Context, id entity.TaskID, status entity.TaskStatus) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}
