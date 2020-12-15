package bo

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type AttachGameParams struct {
	SourceGameID string
	Answer       string
}

func (p *AttachGameParams) FromDTO(value dto.AttachGameResultsRequestBody) {
	p.SourceGameID = value.SourceGameID
	p.Answer = value.Answer
}
