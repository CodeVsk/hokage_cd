package model

import (
	"time"

	"github.com/codevsk/codevsk_golang_cd/internal/entity"
	"github.com/google/uuid"
)

type Deployment struct {
	ID						string    `gorm:"unique;type:varchar(36);primary_key"`
	ApplicationId	string    `gorm:"unique;type:varchar(36);"`
	ImageName			string    `gorm:"type:varchar(255)"`
	ImageTag			string    `gorm:"type:varchar(255);"`
	IsFailed			bool    	`gorm:"type:bit"`
	CreatedAt     time.Time `gorm:"type:datetime"`
	CreatedBy     string    `gorm:"type:varchar(30)"`
}

func (Deployment) TableName() string {
	return "deployment"
}


func ToDeploymentModel(deployment *entity.Deployment) *Deployment {
	return &Deployment {
		ID: deployment.ID.String(),
		ApplicationId: deployment.ApplicationId.String(),
		ImageName: deployment.ImageName,
		ImageTag: deployment.ImageTag,
		IsFailed: deployment.IsFailed,
		CreatedAt: deployment.CreatedAt,
		CreatedBy: deployment.CreatedBy,
	}
}

func ToDeploymentEntity(deployment *Deployment) *entity.Deployment {
	return &entity.Deployment {
		ID: uuid.Must(uuid.Parse(deployment.ID)),
		ApplicationId: uuid.Must(uuid.Parse(deployment.ApplicationId)),
		ImageName: deployment.ImageName,
		ImageTag: deployment.ImageTag,
		IsFailed: deployment.IsFailed,
		CreatedAt: deployment.CreatedAt,
		CreatedBy: deployment.CreatedBy,
	}
}