package task

import (
	"errors"
	"jwt-go/app"
	"jwt-go/model"
	"strconv"

	"gorm.io/gorm"
)

// service struct contains Database pointer
type service struct {
	DB *gorm.DB
}

// Service interface with all CRUD functions signature
type Service interface {
	CreateTask(task model.Task) error
	GetTask(id string) (model.Task, error)
	UpdateTask(task model.Task, id string) error
	DeleteTask(id string) error
}

// NewService will initialize service struct with Database pointer and returned as Service
func NewService(app *app.App) Service {
	return &service{
		DB: app.DB,
	}
}

// CreateTask wil create task in database
func (s *service) CreateTask(task model.Task) error {
	result := s.DB.Create(&task)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// GetTask will retrieve task from database matched with id
func (s *service) GetTask(id string) (model.Task, error) {
	var task model.Task

	result := s.DB.First(&task, id)
	TaskId, err := strconv.Atoi(id)
	if result.Error != nil || err != nil || TaskId != int(task.Id) {
		return model.Task{}, errors.New("ERROR: Book not found for given ID")
	}

	return task, nil
}

// UpdateTask will update task in database matched with id
func (s *service) UpdateTask(task model.Task, id string) error {
	dbTask, err := s.GetTask(id)
	if err != nil {
		return err
	}

	dbTask.Title = task.Title
	dbTask.Body = task.Body
	dbTask.ProfileId = task.ProfileId

	result := s.DB.Save(&dbTask)
	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// DeleteTask will delete the task from database matched with id
func (s *service) DeleteTask(id string) error {
	var task model.Task

	// Permanently delete the task instead of soft delete
	result := s.DB.Unscoped().Delete(&task, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
