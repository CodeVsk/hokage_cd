package result

type ResultType int

const (
	Success    ResultType = 200
	Created    ResultType = 201
	NoContent  ResultType = 204
	BadRequest ResultType = 400
	Conflict   ResultType = 409
	Internal   ResultType = 500
)

type Metadata struct {
	TotalPages int `json:"totalPages,omitempty"`
}

type Body struct {
	Message  string    `json:"message,omitempty"`
	Value    any       `json:"value,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
}

type Result struct {
	StatusCode ResultType
	Body       Body
	Errors     error
}

func NewSuccessResult(value any, totalPages int) Result {
	return Result{
		StatusCode: Success,
		Body: Body{
			Value: value,
			Metadata: &Metadata{
				TotalPages: totalPages,
			},
		},
	}
}

func NewCreatedResult() Result {
	return Result{
		StatusCode: Created,
		Body:       Body{},
	}
}

func NewValidationResult(message string) Result {
	return Result{
		StatusCode: BadRequest,
		Body: Body{
			Message: message,
		},
	}
}

func NewInternalResult(err error) Result {
	return Result{
		StatusCode: Internal,
		Body: Body{
			Message: "Ops! Ocorreu um erro inesperado.",
		},
		Errors: err,
	}
}