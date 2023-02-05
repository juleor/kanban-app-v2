package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	result := []entity.Task{}
	resp := r.db.Raw("SELECT * FROM tasks WHERE user_id = ?", id).Scan(&result)
	if resp.Error != nil {
		return []entity.Task{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return []entity.Task{}, nil
	}
	return result, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	result := r.db.Create(task)
	if result.Error != nil {
		return 0, result.Error
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	result := entity.Task{}
	resp := r.db.Raw("SELECT * FROM tasks WHERE id = ?", id).Scan(&result)
	if resp.Error != nil {
		return entity.Task{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return entity.Task{}, nil
	}
	return result, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	result := []entity.Task{}
	resp := r.db.Raw("SELECT * FROM tasks WHERE category_id = ?", catId).Scan(&result)
	if resp.Error != nil {
		return []entity.Task{}, resp.Error
	}
	if resp.RowsAffected == 0 {
		return []entity.Task{}, nil
	}
	return result, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	result := r.db.Model(&entity.Task{}).Where("id = ?", task.ID).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	result := r.db.Delete(&entity.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
