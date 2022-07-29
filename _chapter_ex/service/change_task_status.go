package service

import (
	"context"
	"fmt"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
)

type ChangeTaskStatus struct {
	DB interface {
		store.Execer
		store.Queryer
	}
	Repo TaskChangeStatuser
}

func (s ChangeTaskStatus) ChangeTaskStatus(ctx context.Context, id entity.TaskID, status entity.TaskStatus) (*entity.Task, error) {
	t, err := s.Repo.ChangeTaskStatus(ctx, s.DB, id, status)
	if err != nil {
		return nil, fmt.Errorf("failed to change status: %w", err)
	}
	return t, nil
}
