package params

type CreateUserParams struct {
	Username string `json:"username" validate:"required,min=2"`
}
