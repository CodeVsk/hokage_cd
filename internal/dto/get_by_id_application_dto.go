package dto

import "github.com/google/uuid"

type GetByIdApplicationOutputDTO struct {
	Id             uuid.UUID `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Repository_url string `json:"repository_url"`
}