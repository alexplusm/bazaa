package bo

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

// TODO: rename to AnswerBO
type UserAnswerBO struct {
	ExtSystemID string `validate:"required"`
	UserID      string `validate:"required"`
	Answer      string `validate:"required"`
}

// TODO: rename to AnswerBO
type UserAnswerCacheBO struct {
	UserID string
	Answer string
}

func (userAnswer *UserAnswerBO) Create(src dto.UserAnswerRequestBody, validate *validator.Validate) error {
	userAnswer.ExtSystemID = src.ExtSystemID
	userAnswer.UserID = src.UserID
	userAnswer.Answer = src.Answer

	if err := validate.Struct(userAnswer); err != nil {
		return fmt.Errorf("user answer: %v", err)
	}
	return nil
}
