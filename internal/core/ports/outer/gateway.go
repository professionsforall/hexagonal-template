package outer

import (
	"context"

	"github.com/professionsforall/hexagonal-template/internal/core/models"
)

type TaskUseCase interface {
	AllTasks(ctx context.Context) ([]*models.TaskModel, error)
	SaveTask(ctx context.Context, task models.TaskModel) (*models.TaskModel, error)
	GetTask(ctx context.Context, id int) (*models.TaskModel, error)
	UpdateTask(ctx context.Context, id int, columns map[string]interface{}) (*models.TaskModel, error)
	DeleteTask(ctx context.Context, id int) (bool, error)
}
