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
