package profile

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
	CreateProfile(profile model.Profile) error
	GetProfile(id string) (model.Profile, error)
	UpdateProfile(profile model.Profile, id string) error
	DeleteProfile(id string) error
	GetAllTask(id string) ([]model.Task, error)
}

// NewService will initialize service struct with Database pointer and returned as Service
func NewService(app *app.App) Service {
	return &service{
		DB: app.DB,
	}
}

// CreateProfile will create profile in database
func (s *service) CreateProfile(profile model.Profile) error {
	result := s.DB.Create(&profile)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// GetProfile will return profile
func (s *service) GetProfile(id string) (model.Profile, error) {
	var profile model.Profile

	result := s.DB.First(&profile, id)
	profileId, err := strconv.Atoi(id)
	if result.Error != nil || err != nil || profileId != int(profile.Id) {
		return model.Profile{}, errors.New("ERROR: Profile not found for given ID")
	}

	return profile, nil
}

// GetAllTask will return all ask attached with particular user
func (s *service) GetAllTask(id string) ([]model.Task, error) {
	var tasks []model.Task

	s.DB.Raw("SELECT * FROM tasks where profile_id = ?", id).Scan(&tasks)

	return tasks, nil
}

// UpdateProfile will update the profile in database
func (s *service) UpdateProfile(profile model.Profile, id string) error {
	dbprofile, err := s.GetProfile(id)
	if err != nil {
		return err
	}

	dbprofile.Name = profile.Name
	dbprofile.Email = profile.Email
	dbprofile.Password = profile.Password

	result := s.DB.Save(&dbprofile)
	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// DeleteProfile will delete the profile from database
func (s *service) DeleteProfile(id string) error {
	var profile model.Profile

	// Permanently delete the profile instead of soft delete
	result := s.DB.Unscoped().Delete(&profile, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
