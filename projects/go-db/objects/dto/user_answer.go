package dto

type UserAnswerRequestBody struct {
	ExtSystemID string `json:"extSystemId"`
	UserID      string `json:"userId"`
	Answer      string `json:"answer"`
}

type UserAnswerResponseData struct {
	Finished   bool            `json:"finished"`
	UserResult string          `json:"userResult"`
	Answers    []UserAnswerDTO `json:"answers"`
}

type UserAnswerDTO struct {
	UserID string `json:"userId"`
	Answer string `json:"answer"`
	Result string `json:"result"`
}
