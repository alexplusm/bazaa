package dto

type UserAnswerRequestBody struct {
	ExtSystemID string `json:"extSystemId"`
	UserID      string `json:"userId"`
	Answer      string `json:"answer"`
}
