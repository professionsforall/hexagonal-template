package repository

import (
	"context"
	"database/sql"
	"fmt"

	boilerModel "github.com/professionsforall/hexagonal-template/internal/adapters/models/sqlboiler/mysql"
	"github.com/professionsforall/hexagonal-template/internal/core/models"
	"github.com/professionsforall/hexagonal-template/internal/core/ports/dependency"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) dependency.TaskRepositoryDependency {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Save(ctx context.Context, task models.TaskModel) error {
	data := boilerModel.Task{
		Title:       null.NewString(task.Title, true),
		Description: null.NewString(task.Description, true),
		Color:       null.NewString(task.Color, true),
		StartsAt:    null.NewTime(task.StartsAt, true),
		DoneAt:      null.NewTime(task.DoneAt, true),
	}
	err := data.Insert(ctx, t.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
func (t *TaskRepository) Get(ctx context.Context, id int) (*models.TaskModel, error) {
	task, err := boilerModel.Tasks(boilerModel.TaskWhere.ID.EQ(uint64(id))).One(ctx, t.db)
	if err != nil {
		fmt.Println(err)
	}
	if task == nil {
		return nil, err
	}
	return &models.TaskModel{
		ID:          int(task.ID),
		Title:       task.Title.String,
		Description: task.Description.String,
		Color:       task.Color.String,
		StartsAt:    task.StartsAt.Time,
		DoneAt:      task.DoneAt.Time,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}, err
}

func (t *TaskRepository) All(ctx context.Context) (tasks []*models.TaskModel, err error) {
	allTasks, err := boilerModel.Tasks().All(ctx, t.db)
	for _, task := range allTasks {
		taskModel := new(models.TaskModel)
		taskModel.ID = int(task.ID)
		taskModel.Title = task.Title.String
		taskModel.Description = task.Description.String
		taskModel.Color = task.Color.String
		taskModel.StartsAt = task.StartsAt.Time
		taskModel.DoneAt = task.DoneAt.Time
		taskModel.CreatedAt = task.CreatedAt.Time
		taskModel.UpdatedAt = task.UpdatedAt.Time
		taskModel.DeletedAt = task.DeletedAt.Time
		tasks = append(tasks, taskModel)
	}
	return tasks, err
}

func (t *TaskRepository) Update(ctx context.Context, id int, columns map[string]interface{}) error {
	_, err := boilerModel.Tasks(boilerModel.TaskWhere.ID.EQ(uint64(id))).UpdateAll(ctx, t.db, columns)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) Delete(ctx context.Context, id int) error {
	rows, err := boilerModel.Tasks(boilerModel.TaskWhere.ID.EQ(uint64(id))).DeleteAll(ctx, t.db)

	if rows == 0 {
		return err
	}
	return nil
}
