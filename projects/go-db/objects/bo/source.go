package bo

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/timeutils"
)

type SourceBO struct {
	SourceID  string
	Type      int
	CreatedAt time.Time
}

func (s *SourceBO) ToDTO() dto.SourceDTO {
	return dto.SourceDTO{
		SourceID:  s.SourceID,
		CreatedAt: timeutils.FromTimeToStrTimestamp(s.CreatedAt),
		Type:      s.Type,
	}
}
