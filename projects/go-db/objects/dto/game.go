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

// TODO!
//“sources”: [ { “type”: “schedule” | “file” | “url”, “sourceId”: “someid” } ]

type OptionDTO struct {
	Option int    `json:"option"`
	Text   string `json:"text"`
}

type SourceDTO struct {
	Type     string `json:"type"`
	SourceID string `json:"sourceId"`
}

type QuestionDTO struct {
	AnswerType int         `json:"answerType"`
	Text       string      `json:"text"`
	Options    []OptionDTO `json:"options"`
}

type GameDetailsResponseBody struct {
	StartDate  string      `json:"startDate"`
	FinishDate string      `json:"finishDate"`
	Question   QuestionDTO `json:"question"`
	Sources    []SourceDTO `json:"sources"`
}
