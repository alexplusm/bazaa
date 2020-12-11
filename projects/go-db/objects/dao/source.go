package dao

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type SourceInsertDAO struct {
	Type      int
	CreatedAt int64
	GameID    string
}

type SourceRetrieveDAO struct {
	SourceID  string
	CreatedAt int64
	Type      int
}

func (s *SourceRetrieveDAO) ToBO() bo.SourceBO {
	return bo.SourceBO{
		SourceID:  s.SourceID,
		CreatedAt: time.Unix(s.CreatedAt, 0),
		Type:      s.Type,
	}
}
