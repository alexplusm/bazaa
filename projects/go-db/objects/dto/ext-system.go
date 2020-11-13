package dto

type CreateExtSystemRequestBody struct {
	ID             string `json:"ext_system_id"`
	Description    string `json:"description"`
	PostResultsURL string `json:"post_results_url"`
}
