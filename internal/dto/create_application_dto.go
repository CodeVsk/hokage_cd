package dto

type CreateApplicationInputDTO struct {
	Name            string `json:"name,omitempty"`
	Slug            string `json:"slug,omitempty"`
	Repository_url  string `json:"repository_url,omitempty"`
	RepositoryType  string `json:"repository_type,omitempty"`
	RepositoryName  string `json:"repository_name,omitempty"`
	RepositoryOwner string `json:"repository_owner,omitempty"`
	RepositoryToken string `json:"repository_token,omitempty"`
	FilePath        string `json:"file_path,omitempty"`
	FileName        string `json:"file_name,omitempty"`
}