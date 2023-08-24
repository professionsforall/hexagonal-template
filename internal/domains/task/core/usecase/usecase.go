package usecase

import (
	"context"

	"github.com/professionsforall/hexagonal-template/internal/domains/task/core/models"
	"github.com/professionsforall/hexagonal-template/internal/domains/task/core/ports/dependency"
	"github.com/professionsforall/hexagonal-template/internal/domains/task/core/ports/outer"
)

type TaskHandler struct {
	TaskRepositoryDependency dependency.TaskRepositoryDependency
}

func NewTaskHandler(taskRepositoryDepenndency dependency.TaskRepositoryDependency) outer.TaskUseCase {
	return &TaskHandler{TaskRepositoryDependency: taskRepositoryDepenndency}
}

func (r *TaskHandler) SaveTask(ctx context.Context, task models.TaskModel) error {
	return r.TaskRepositoryDependency.Save(ctx, task)
}

func (r *TaskHandler) GetTask(ctx context.Context, id int) (*models.TaskModel, error) {
	return r.TaskRepositoryDependency.Get(ctx, id)
}

func (r *TaskHandler) AllTasks(ctx context.Context) ([]*models.TaskModel, error) {
	return r.TaskRepositoryDependency.All(ctx)
}

func (r *TaskHandler) UpdateTask(ctx context.Context, id int, columns map[string]interface{}) error {
	return r.TaskRepositoryDependency.Update(ctx, id, columns)
}

func (r *TaskHandler) DeleteTask(ctx context.Context, id int) error {
	return r.TaskRepositoryDependency.Delete(ctx, id)
}
