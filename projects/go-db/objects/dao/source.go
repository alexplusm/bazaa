package dao

const (
	ArchiveSourceType = iota
	ScheduleSourceType
)

//TODO: schedules

type SourceDAO struct {
	Type      int
	CreatedAt int64
	GameID    string
}

type Source2DAO struct {
	SourceID string
	Type     int
}
