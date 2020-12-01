package dto

type StatisticsUsersInnerDTO struct {
	TotalScreenshots int     `json:"totalScreenshots"`
	RightAnswers     int     `json:"rightAnswers"`
	MatchWithExpert  int     `json:"matchWithExpert"`
	AverageAccuracy  float64 `json:"averageAccuracy"`
}

type StatisticsUserDTO struct {
	Date       string                  `json:"date"`
	Statistics StatisticsUsersInnerDTO `json:"statistics"`
}

type StatsUserDTO struct {
	Total   StatisticsUsersInnerDTO `json:"total"`
	History []StatisticsUserDTO     `json:"history"`
}

type StatsUserTotalOnlyDTO struct {
	Total StatisticsUsersInnerDTO `json:"total"`
}

type UserAnswerForScreenshotResultDTO struct {
	UserID string `json:"userId"`
	Answer string `json:"answer"`
	Result string `json:"result"`
}

type ScreenshotResultsDTO struct {
	Finished bool                               `json:"finished"`
	Answers  []UserAnswerForScreenshotResultDTO `json:"answers"`
}

//{
//"screenshotsResolved": 653,
//"screenshotsLeft": 10265,
//"usersUnique": 210,
//"usersActive": 154
//}
type GameStatsDTO struct {
	ScreenshotsResolved int `json:"screenshotsResolved"`
	ScreenshotsLeft     int `json:"screenshotsLeft"`
	UsersUnique         int `json:"usersUnique"`
	UsersActive         int `json:"usersActive"`
}
