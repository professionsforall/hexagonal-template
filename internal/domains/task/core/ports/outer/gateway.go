package outer

import (
	"context"

	"github.com/professionsforall/hexagonal-template/internal/domains/task/core/models"
)

type TaskUseCase interface {
	AllTasks(ctx context.Context) ([]*models.TaskModel, error)
	SaveTask(ctx context.Context, task models.TaskModel) error
	GetTask(ctx context.Context, id int) (*models.TaskModel, error)
	UpdateTask(ctx context.Context, id int, columns map[string]interface{}) error
	DeleteTask(ctx context.Context, id int) error
}
