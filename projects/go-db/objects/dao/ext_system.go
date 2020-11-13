package dao

import "github.com/Alexplusm/bazaa/projects/go-db/objects/bo"

type ExtSystemDAO struct {
	ID             string
	Description    string
	PostResultsURL string
}

func (extSystem *ExtSystemDAO) FromBO(bo bo.ExtSystemBO) {
	extSystem.ID = bo.ID
	extSystem.Description = bo.Description
	extSystem.PostResultsURL = bo.PostResultsURL
}

func (extSystem *ExtSystemDAO) HasID() bool {
	return extSystem.ID != ""
}
