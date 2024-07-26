package dto

type FindByEmailDTO struct {
	Email string `json:"email" validate:"required,email"`
}

type FindByIDDTO struct {
	ID string `json:"id" validate:"required"`
}

type CreateEmailDTO struct {
	Email  string `json:"email" validate:"required,email"`
	Status string `json:"status" validate:"required,oneof=PENDING SENDED"`
}

type DeleteEmailDTO struct {
	ID string `json:"id" validate:"required"`
}

type UpdateEmailDTO struct {
	ID     string `json:"id" validate:"required"`
	Email  string `json:"email" validate:"omitempty,email"`
	Status string `json:"status" validate:"omitempty,oneof=PENDING SENDED"`
}
