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

type StatisticUserAnswersDTO struct {
	Total   StatisticsUsersInnerDTO `json:"total"`
	History []StatisticsUserDTO     `json:"history"`
}

type StatisticUserAnswersTotalDTO struct {
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

type GameStatsDTO struct {
	ScreenshotsResolved int `json:"screenshotsResolved"`
	ScreenshotsLeft     int `json:"screenshotsLeft"`
	UsersUnique         int `json:"usersUnique"`
	UsersActive         int `json:"usersActive"`
}

type LeadersDTO struct {
	UserID     string                  `json:"userId"`
	Statistics StatisticsUsersInnerDTO `json:"statistics"`
}

type LeadersResponseDTO struct {
	Leaders []LeadersDTO `json:"leaders"`
}
