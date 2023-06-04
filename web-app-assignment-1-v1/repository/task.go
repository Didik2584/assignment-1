package repository

import (
	"a21hc3NpZ25tZW50/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(categoryID int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	return err
}

func (t *taskRepository) Update(task *model.Task) error {
	err := t.db.Save(task).Error
	return err
}

func (t *taskRepository) Delete(id int) error {
	err := t.db.Delete(&model.Task{}, id).Error
	return err
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	var tasks []model.Task
	err := t.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskRepository) GetTaskCategory(categoryID int) ([]model.TaskCategory, error) {
	var taskCategories []model.TaskCategory
	err := t.db.Table("tasks").
		Select("tasks.id, tasks.title, categories.name as category").
		Joins("JOIN categories ON tasks.category_id = categories.id").
		Where("categories.id = ?", categoryID).
		Find(&taskCategories).Error
	if err != nil {
		return nil, err
	}
	return taskCategories, nil
}
