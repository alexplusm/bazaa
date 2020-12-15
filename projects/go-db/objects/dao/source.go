package dao

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type SourceBaseDAO struct {
	GameID    string
	Type      int
	CreatedAt int64
	Value     string
}

type SourceInsertDAO struct {
	SourceBaseDAO
}

type SourceRetrieveDAO struct {
	SourceID string
	SourceBaseDAO
}

func (s *SourceRetrieveDAO) ToBO() bo.SourceBO {
	// Value
	return bo.SourceBO{
		SourceID:  s.SourceID,
		CreatedAt: time.Unix(s.CreatedAt, 0),
		Type:      s.Type,
	}
}
