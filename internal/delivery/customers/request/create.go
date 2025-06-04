package request

type CreateCustomerRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type CustomerQueryParams struct {
	Q     string `url:"q"`
	Page  int    `url:"page"`
	Limit int    `url:"limit"`
}
