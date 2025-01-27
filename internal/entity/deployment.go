package entity

import (
	"time"

	"github.com/google/uuid"
)

type Deployment struct {
	ID            uuid.UUID `json:"id"`
	ApplicationId uuid.UUID `json:"application_id"`
	ImageName     string    `json:"image_name"`
	ImageTag      string    `json:"image_tag"`
	IsFailed      bool      `json:"is_enabled"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     string    `json:"created_by"`
}

func NewDeployment(applicationId uuid.UUID, imageName string, imageTag string, isFailed bool, createdBy string) *Deployment {
	return &Deployment{
		ID: uuid.New(),
		ApplicationId: applicationId,
		ImageName: imageName,
		ImageTag: imageTag,
		IsFailed: isFailed,
		CreatedAt: time.Now(),
		CreatedBy: createdBy,
	}
}