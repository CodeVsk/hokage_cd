package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrIDIsRequired = errors.New("id is required")
	ErrNameIsRequired = errors.New("name is required")
	ErrSlugIsRequired = errors.New("slug is required")
	ErrCreatedByIsRequired = errors.New("create by is required")
	
	ErrNameCharacterLimitExceeded = errors.New("character limit exceeded for name")
	ErrSlugCharacterLimitExceeded = errors.New("character limit exceeded for slug")
	ErrRepositoryUrlCharacterLimitExceeded = errors.New("character limit exceeded for repository url")
	ErrCreatedByCharacterLimitExceeded = errors.New("character limit exceeded for created by")
	ErrUpdatedByCharacterLimitExceeded = errors.New("character limit exceeded for updated by")
)

type Application struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	RepositoryUrl string    `json:"repository_url"`
	IsEnabled     bool      `json:"is_enabled"`
	AutoRollback  bool      `json:"auto_rollback"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedBy     string    `json:"updated_by"`
}

func NewApplication(name string, slug string, repository_url string, created_by string) (*Application, error) {
	application := &Application{
		ID: uuid.New(),
		Name: name,
		Slug: slug,
		RepositoryUrl: repository_url,
		CreatedAt: time.Now(),
		CreatedBy: created_by,
		IsEnabled: true,
		AutoRollback: false,
	}

	err := application.Validate()
	if err != nil {
		return nil, err;
	}

	return application, nil
}

func (a *Application) Validate() error {
	if(a.ID.String() == "") {
		return ErrIDIsRequired
	}

	if(a.Name == "") {
		return ErrNameIsRequired
	}

	if(a.Slug == "") {
		return ErrSlugIsRequired
	}

	if(a.CreatedBy == "") {
		return ErrCreatedByIsRequired
	}

	if(len(a.Name) > 40){
		return ErrNameCharacterLimitExceeded
	}

	if(len(a.Slug) > 40){
		return ErrSlugCharacterLimitExceeded
	}

	if(len(a.RepositoryUrl) > 255){
		return ErrRepositoryUrlCharacterLimitExceeded
	}

	if(len(a.CreatedBy) > 30){
		return ErrCreatedByCharacterLimitExceeded
	}

	if(len(a.UpdatedBy) > 30){
		return ErrUpdatedByCharacterLimitExceeded
	}

	return nil
}