package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type AnswerService struct {
	AnswerRepo interfaces.IAnswerRepository
}

func (service *AnswerService) GetUserStatistics(
	userID, totalOnly, gameIDs, from, to string,
) error {
	//var toBO, fromBO time.Time
	//var result

	toBO := time.Now()
	fromBO := time.Now()

	if to != "" {
		i, _ := strconv.ParseInt(to, 10, 64)
		toBO = time.Unix(i, 0)
	} else {
		toBO = time.Now()
	}
	if from != "" {
		i, _ := strconv.ParseInt(from, 10, 64)
		fromBO = time.Unix(i, 0)
	} else {
		// TODO: достать первую игру пользователя
		fromBO = time.Date(2019, 1, 1, 1, 0, 0, 0, time.UTC)
	}

	fmt.Println(toBO, fromBO)

	gIds := strings.Split(gameIDs, ",")

	if len(gIds) == 0 {
		//service.AnswerRepo.
	}

	// gameIDs if == "" retrieve

	return nil
}
