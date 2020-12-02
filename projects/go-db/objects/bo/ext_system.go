package bo

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type ExtSystemBO struct {
	ID             string
	Description    string `validate:"required"`
	PostResultsURL string `validate:"required"`
}

func (extSystem *ExtSystemBO) FromDTO(src dto.CreateExtSystemRequestBody, validate *validator.Validate) error {
	extSystem.ID = src.ID
	extSystem.Description = src.Description
	extSystem.PostResultsURL = src.PostResultsURL

	if err := validate.Struct(extSystem); err != nil {
		return fmt.Errorf("ExtSystem from DTO validation: %v", err)
	}

	return nil
}

func (extSystem *ExtSystemBO) HasID() bool {
	return extSystem.ID != ""
}

func (extSystem *ExtSystemBO) ToDTO() dto.ExtSystemListItem {
	return dto.ExtSystemListItem{
		ID:             extSystem.ID,
		Description:    extSystem.Description,
		PostResultsURL: extSystem.PostResultsURL,
	}
}
