package dao

const (
	ArchiveSourceType = iota
	ScheduleSourceType
)

type SourceDAO struct {
	Type      int
	CreatedAt int64
	GameID    string
}
