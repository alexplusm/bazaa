package dao

// TODO: refactor !!!

// TODO: rename CREATE
type ScreenshotDAO struct {
	Filename string
	GameID   string
	SourceID string
	// INFO: if ExpertAnswer == "" -> without expert answer
	ExpertAnswer string
}

// RETRIEVE
type ScreenshotDAOFull struct {
	ScreenshotID string
	GameID       string
	SourceID     string
	Filename     string
	ExpertAnswer string
	UsersAnswer  string
}

type ScreenshotURLDAO struct {
	ScreenshotID string
	ImageURL     string
}

type AnsweredScreenshotsDAO struct {
	UserID []string
	Count  int
}
