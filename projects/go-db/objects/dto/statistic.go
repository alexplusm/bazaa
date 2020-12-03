package dto

type StatisticUserAnswersDTO struct {
	Total   StatisticUsersInnerDTO `json:"total"`
	History []StatisticsUserDTO    `json:"history"`
}

type StatisticUserAnswersTotalDTO struct {
	Total StatisticUsersInnerDTO `json:"total"`
}

type StatisticUsersInnerDTO struct {
	TotalScreenshots int     `json:"totalScreenshots"`
	RightAnswers     int     `json:"rightAnswers"`
	MatchWithExpert  int     `json:"matchWithExpert"`
	AverageAccuracy  float64 `json:"averageAccuracy"`
}

type StatisticsUserDTO struct {
	Date       string                 `json:"date"`
	Statistics StatisticUsersInnerDTO `json:"statistics"`
}

type StatisticGameDTO struct {
	ScreenshotsResolved int `json:"screenshotsResolved"`
	ScreenshotsLeft     int `json:"screenshotsLeft"`
	UsersUnique         int `json:"usersUnique"`
	UsersActive         int `json:"usersActive"`
}

type LeadersDTO struct {
	UserID     string                 `json:"userId"`
	Statistics StatisticUsersInnerDTO `json:"statistics"`
}

type LeadersResponseDTO struct {
	Leaders []LeadersDTO `json:"leaders"`
}
