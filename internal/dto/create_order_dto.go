package dto

type CreateApplicationInputDTO struct {
	Name           string
	Slug           string
	Repository_url string
}

type CreateApplicationOutputDTO struct{}