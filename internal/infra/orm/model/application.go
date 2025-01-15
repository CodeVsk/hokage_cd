package model

import (
	"time"

	"github.com/codevsk/codevsk_golang_cd/internal/entity"
	"github.com/google/uuid"
)

type Application struct {
	ID            uuid.UUID `gorm:"type:UNIQUEIDENTIFIER;primary_key"`
	Name          string    `gorm:"type:varchar(40)"`
	Slug          string    `gorm:"type:varchar(40);size:40;unique"`
	RepositoryUrl string    `gorm:"type:varchar(255)"`
	IsEnabled     bool      `gorm:"type:bit"`
	AutoRollback  bool      `gorm:"type:bit"`
	CreatedAt     time.Time `gorm:"type:datetime"`
	CreatedBy     string    `gorm:"type:varchar(30)"`
	UpdatedAt     time.Time `gorm:"type:datetime"`
	UpdatedBy     string    `gorm:"type:varchar(30)"`
}

func (Application) TableName() string {
	return "application"
}

func (a *Application) ToModel(application entity.Application) Application {
	return Application {
		ID: application.ID,
		Name: application.Name,
		Slug: application.Slug,
		RepositoryUrl: application.RepositoryUrl,
		IsEnabled: application.IsEnabled,
		AutoRollback: application.AutoRollback,
		CreatedAt: application.CreatedAt,
		CreatedBy: application.CreatedBy,
		UpdatedAt: application.UpdatedAt,
		UpdatedBy: application.UpdatedBy,
	}
}

func (a *Application) ToEntity(application Application) entity.Application {
	return entity.Application {
		ID: application.ID,
		Name: application.Name,
		Slug: application.Slug,
		RepositoryUrl: application.RepositoryUrl,
		IsEnabled: application.IsEnabled,
		AutoRollback: application.AutoRollback,
		CreatedAt: application.CreatedAt,
		CreatedBy: application.CreatedBy,
		UpdatedAt: application.UpdatedAt,
		UpdatedBy: application.UpdatedBy,
	}
}