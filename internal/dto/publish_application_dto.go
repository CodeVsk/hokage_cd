package dto

type PublishApplicationInputDTO struct {
	ImageName       string `json:"image_name" validate:"required"`
	ImageTag        string `json:"image_tag" validate:"required"`
	ApplicationSlug string `json:"application_slug" validate:"required"`
}