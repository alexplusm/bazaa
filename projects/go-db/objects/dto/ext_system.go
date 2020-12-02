package dto

type CreateExtSystemRequestBody struct {
	ID             string `json:"extSystemId"`
	Description    string `json:"description"`
	PostResultsURL string `json:"postResultsUrl"`
}

type CreateExtSystemResponseBody struct {
	ID string `json:"extSystemId"`
}

type ExtSystemListItem struct {
	ID             string `json:"extSystemId"`
	Description    string `json:"description"`
	PostResultsURL string `json:"postResultsUrl"`
}

type ExtSystemListResponseBody struct {
	ExtSystems []ExtSystemListItem `json:"extSystems"`
}
