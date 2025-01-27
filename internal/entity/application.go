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
	ErrCreatedByIsRequired = errors.New("created by is required")
	ErrRepositoryTypeIsRequired = errors.New("repository type is required")
	ErrRepositoryNameIsRequired = errors.New("repository name is required")
	ErrRepositoryOwnerIsRequired = errors.New("repository owner is required")
	ErrRepositoryTokenIsRequired = errors.New("repository token is required")
	ErrFilePathIsRequired = errors.New("file path is required")
	ErrFileNameIsRequired = errors.New("file name is required")

	ErrNameCharacterLimitExceeded = errors.New("character limit exceeded for name")
	ErrSlugCharacterLimitExceeded = errors.New("character limit exceeded for slug")
	ErrRepositoryUrlCharacterLimitExceeded = errors.New("character limit exceeded for repository url")
	ErrCreatedByCharacterLimitExceeded = errors.New("character limit exceeded for created by")
	ErrUpdatedByCharacterLimitExceeded = errors.New("character limit exceeded for updated by")
	ErrRepositoryTypeCharacterLimitExceeded = errors.New("character limit exceeded for repository type")
	ErrRepositoryNameCharacterLimitExceeded = errors.New("character limit exceeded for repository name")
	ErrRepositoryOwnerCharacterLimitExceeded = errors.New("character limit exceeded for repository owner")
	ErrRepositoryTokenCharacterLimitExceeded = errors.New("character limit exceeded for repository token")
	ErrFilePathCharacterLimitExceeded = errors.New("character limit exceeded for file path")
	ErrFileNameCharacterLimitExceeded = errors.New("character limit exceeded for file name")
)

type Application struct {
	ID            	uuid.UUID `json:"id"`
	Name          	string    `json:"name"`
	Slug          	string    `json:"slug"`
	RepositoryUrl 	string    `json:"repository_url"`
	RepositoryType 	string 		`json:"repository_type"`
	RepositoryName 	string 		`json:"repository_name"`
	RepositoryOwner string 		`json:"repository_owner"`
	RepositoryToken string 		`json:"repository_token"`
	FilePath 				string 		`json:"file_path"`
	FileName 				string 		`json:"file_name"`
	IsEnabled     	bool      `json:"is_enabled"`
	AutoRollback  	bool      `json:"auto_rollback"`
	CreatedAt     	time.Time `json:"created_at"`
	CreatedBy     	string    `json:"created_by"`
	UpdatedAt     	time.Time `json:"updated_at"`
	UpdatedBy     	string    `json:"updated_by"`
}

func NewApplication(name string, slug string, repository_url string, repository_type string, repository_name string, repository_owner string, repository_token string, file_path string, file_name string, created_by string) (*Application, error) {
	application := &Application{
		ID: uuid.New(),
		Name: name,
		Slug: slug,
		RepositoryUrl: repository_url,
		RepositoryType: repository_type,
		RepositoryName: repository_name,
		RepositoryOwner: repository_owner,
		RepositoryToken: repository_token,
		FilePath: file_path,
		FileName: file_name,
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

func (a *Application) SetName(name string) {
	a.Name = name
}

func (a *Application) SetSlug(slug string) {
	a.Slug = slug
}

func (a *Application) SetRepositoryUrl(repository_url string) {
	a.RepositoryUrl = repository_url
}

func (a *Application) SetUpdated(updatedBy string) {
	a.UpdatedBy = updatedBy
	a.UpdatedAt = time.Now()
}

func (a *Application) SetRepositoryType(repository_type string) {
	a.RepositoryType = repository_type
}

func (a *Application) SetRepositoryName(repository_name string) {
	a.RepositoryName = repository_name
}

func (a *Application) SetRepositoryOwner(repository_owner string) {
	a.RepositoryOwner = repository_owner
}

func (a *Application) SetRepositoryToken(repository_token string) {
	a.RepositoryToken = repository_token
}

func (a *Application) SetFilePath(file_path string) {
	a.FilePath = file_path
}

func (a *Application) SetFileName(file_name string) {
	a.FileName = file_name
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

	if(a.RepositoryType == "") {
		return ErrRepositoryTypeIsRequired
	}

	if(a.RepositoryName == "") {
		return ErrRepositoryNameIsRequired
	}

	if(a.RepositoryOwner == "") {
		return ErrRepositoryOwnerIsRequired
	}

	if(a.RepositoryToken == "") {
		return ErrRepositoryTokenIsRequired
	}

	if(a.FilePath == "") {
		return ErrFilePathIsRequired
	}

	if(a.FileName == "") {
		return ErrFileNameIsRequired
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

	if(len(a.RepositoryName) > 30){
		return ErrRepositoryNameCharacterLimitExceeded
	}

	if(len(a.RepositoryOwner) > 30){
		return ErrRepositoryOwnerCharacterLimitExceeded
	}

	if(len(a.RepositoryToken) > 40){
		return ErrRepositoryTokenCharacterLimitExceeded
	}

	if(len(a.RepositoryType) > 15){
		return ErrRepositoryTypeCharacterLimitExceeded
	}

	if(len(a.FilePath) > 60){
		return ErrFilePathCharacterLimitExceeded
	}

	if(len(a.FileName) > 25){
		return ErrFileNameCharacterLimitExceeded
	}

	if(len(a.CreatedBy) > 30){
		return ErrCreatedByCharacterLimitExceeded
	}

	if(len(a.UpdatedBy) > 30){
		return ErrUpdatedByCharacterLimitExceeded
	}

	return nil
}