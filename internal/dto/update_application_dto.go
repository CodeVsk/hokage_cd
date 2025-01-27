package dto

type UpdateApplicationInputDTO struct {
	Name           string `json:"name,omitempty"`
	Slug           string `json:"slug,omitempty"`
	Repository_url string `json:"repository_url,omitempty"`
}