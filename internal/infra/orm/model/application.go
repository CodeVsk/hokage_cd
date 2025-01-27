package model

import (
	"time"

	"github.com/codevsk/codevsk_golang_cd/internal/entity"
	"github.com/google/uuid"
)

type Application struct {
	ID            	string 		`gorm:"unique;type:varchar(36);primary_key"`
	Name          	string    `gorm:"type:varchar(40)"`
	Slug          	string    `gorm:"type:varchar(40);size:40;unique"`
	RepositoryUrl 	string    `gorm:"type:varchar(255)"`
	RepositoryType 	string 		`gorm:"type:varchar(15)"`
	RepositoryName 	string 		`gorm:"type:varchar(30)"`
	RepositoryOwner string 		`gorm:"type:varchar(30)"`
	RepositoryToken string 		`gorm:"type:varchar(40)"`
	FilePath 				string 		`gorm:"type:varchar(60)"`
	FileName 				string 		`gorm:"type:varchar(25)"`
	IsEnabled     	bool      `gorm:"type:bit"`
	AutoRollback  	bool      `gorm:"type:bit"`
	CreatedAt     	time.Time `gorm:"type:datetime"`
	CreatedBy     	string    `gorm:"type:varchar(30)"`
	UpdatedAt     	time.Time `gorm:"type:datetime"`
	UpdatedBy     	string    `gorm:"type:varchar(30)"`
}

func (Application) TableName() string {
	return "application"
}

func ToApplicationModel(application *entity.Application) *Application {
	return &Application {
		ID: application.ID.String(),
		Name: application.Name,
		Slug: application.Slug,
		RepositoryUrl: application.RepositoryUrl,
		RepositoryType: application.RepositoryType,
		RepositoryName: application.RepositoryName,
		RepositoryOwner: application.RepositoryOwner,
		RepositoryToken: application.RepositoryToken,
		FilePath: application.FilePath,
		FileName: application.FileName,
		IsEnabled: application.IsEnabled,
		AutoRollback: application.AutoRollback,
		CreatedAt: application.CreatedAt,
		CreatedBy: application.CreatedBy,
		UpdatedAt: application.UpdatedAt,
		UpdatedBy: application.UpdatedBy,
	}
}

func ToApplicationEntity(application *Application) *entity.Application {
	return &entity.Application {
		ID: uuid.Must(uuid.Parse(application.ID)),
		Name: application.Name,
		Slug: application.Slug,
		RepositoryUrl: application.RepositoryUrl,
		RepositoryType: application.RepositoryType,
		RepositoryName: application.RepositoryName,
		RepositoryOwner: application.RepositoryOwner,
		RepositoryToken: application.RepositoryToken,
		FilePath: application.FilePath,
		FileName: application.FileName,
		IsEnabled: application.IsEnabled,
		AutoRollback: application.AutoRollback,
		CreatedAt: application.CreatedAt,
		CreatedBy: application.CreatedBy,
		UpdatedAt: application.UpdatedAt,
		UpdatedBy: application.UpdatedBy,
	}
}