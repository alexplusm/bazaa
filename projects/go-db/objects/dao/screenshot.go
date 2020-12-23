package dao

type ScreenshotCreateDAO struct {
	Filename string
	GameID   string
	SourceID string
	// INFO: if ExpertAnswer == "" -> without expert answer
	ExpertAnswer string
}

type ScreenshotRetrieveDAO struct {
	ScreenshotID string
	GameID       string
	SourceID     string
	Filename     string
	ExpertAnswer string
	UsersAnswer  string
}

// TODO: refactor !!!

type ScreenshotURLDAO struct {
	ScreenshotID string
	ImageURL     string
}

type AnsweredScreenshotsDAO struct {
	UserID []string
	Count  int
}
