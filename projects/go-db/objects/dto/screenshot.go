package dto

type UserAnswerForScreenshotResultDTO struct {
	UserID string `json:"userId"`
	Answer string `json:"answer"`
	Result string `json:"result"`
}

type ScreenshotResultsDTO struct {
	Finished    bool                               `json:"finished"`
	UsersAnswer string                             `json:"usersAnswer"`
	Answers     []UserAnswerForScreenshotResultDTO `json:"answers"`
}
