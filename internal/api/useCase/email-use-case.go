package useCase

import (
	"FirstAPI/internal/api/customError"
	"FirstAPI/internal/api/model"
	"FirstAPI/internal/api/repository"
	"errors"
)

type EmailUseCase struct {
	repository *repository.EmailRepository
}

func NewEmailUseCase(repository *repository.EmailRepository) *EmailUseCase {
	return &EmailUseCase{
		repository: repository,
	}
}

func (e *EmailUseCase) GetEmails() ([]model.Email, error) {
	emails, err := e.repository.Find()
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (e *EmailUseCase) GetByID(id string) (model.Email, error) {
	email, err := e.repository.FindByID(id)
	if err != nil {
		return model.Email{}, err
	}

	return email, nil
}

func (e *EmailUseCase) Create(email string, status string) (model.Email, error) {
	emailExists, err := e.repository.FindByEmail(email)
	if err != nil {
		if !errors.Is(err, customError.ErrNotFound) {
			return model.Email{}, err
		}
	}

	if emailExists.ID != "" {
		return model.Email{}, errors.New("email already exists")
	}

	result, err := e.repository.Insert(email, status)
	if err != nil {
		return model.Email{}, err
	}

	return result, nil
}

func (e *EmailUseCase) Delete(id string) error {
	err := e.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmailUseCase) Update(id string, email string, status string) (model.Email, error) {
	emailExists, err := e.repository.FindByEmail(email)
	if err != nil {
		if !errors.Is(err, customError.ErrNotFound) {
			return model.Email{}, err
		}
	}

	if emailExists.ID != "" {
		return model.Email{}, errors.New("email already exists")
	}

	result, err := e.repository.Update(id, email, status)
	if err != nil {
		return model.Email{}, err
	}

	return result, nil
}
