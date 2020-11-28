package dto

type CreateGameRequestBody struct {
	ExtSystemID string `json:"extSystemId"`
	Name        string `json:"name"`
	AnswerType  int    `json:"answerType"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Question    string `json:"question"`
	Options     string `json:"options"`
}

type CreateGameResponseBody struct {
	GameID string `json:"gameId"`
}

type PrepareGameResponseBody struct {
	GameID string `json:"gameId"`
}

type GameListResponseBody struct {
	Games []GameItemResponseBody `json:"games"`
}

type GameItemResponseBody struct {
	GameID string `json:"gameId"`
	Name   string `json:"name"`
	Status string `json:"status"`
	From   string `json:"from"`
	To     string `json:"to"`
}
