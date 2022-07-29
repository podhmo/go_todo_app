package store

import (
	"context"

	"github.com/budougumi0617/go_todo_app/entity"
)

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	sql := `INSERT INTO task
			(title, status, created, modified)
	VALUES (?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, t.Title, t.Status,
		r.Clocker.Now(), r.Clocker.Now(),
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT 
				id, title,
				status, created, modified 
			FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) ChangeTaskStatus(
	ctx context.Context, db interface {
		Execer
		Queryer
	}, id entity.TaskID, status entity.TaskStatus,
) (*entity.Task, error) {
	sql := `Update task
				SET status=?,
				WHERE id=?
			;`

	_, err := db.ExecContext(ctx, sql, id, status)
	if err != nil {
		return nil, err
	}

	return r.GetTaskByID(ctx, db, id)
}

func (r *Repository) GetTaskByID(
	ctx context.Context, db Queryer, id entity.TaskID,
) (*entity.Task, error) {
	task := entity.Task{}
	sql := `SELECT 
				id, title,
				status, created, modified 
			FROM task
			WHERE
				id=?
			;`

	if err := db.GetContext(ctx, &task, sql, id); err != nil {
		return nil, err
	}
	return &task, nil
}
