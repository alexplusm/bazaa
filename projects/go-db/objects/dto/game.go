package dto

// TODO: camelCase !!! and refactor test script

type CreateGameRequestBody struct {
	ExtSystemID string `json:"ext_system_id"`
	Name        string `json:"name"`
	AnswerType  int    `json:"answer_type"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Question    string `json:"question"`
	Options     string `json:"options"`
}
