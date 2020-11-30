package services

import (
	"fmt"
	"sort"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type AnswerService struct {
	AnswerRepo interfaces.IAnswerRepository
}

// TODO: total only !!!
func (service *AnswerService) GetUserStatistics(
	userID string, totalOnly bool, games []bo.GameBO, from, to time.Time,
) ([]bo.StatisticsUserBO, error) {
	gameIds := make([]string, 0, len(games))

	for _, game := range games {
		gameIds = append(gameIds, game.GameID)
	}

	res, err := service.AnswerRepo.SelectAnswersByUser(userID, gameIds, from, to)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, fmt.Errorf("get user statistics: %v", err)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i].AnswerDate < res[j].AnswerDate
	})

	//fmt.Printf("Res: %+v | %v\n", res, len(res))

	year, month, day := time.Unix(res[0].AnswerDate, 0).Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	end := time.Unix(res[len(res)-1].AnswerDate, 0)

	resuults := make([]bo.StatisticsUserBO, 0, len(res))

	currentDay := start
	for currentDay.Before(end) {
		for _, r := range res {

			date := time.Unix(r.AnswerDate, 0)
			next := currentDay.AddDate(0, 0, 1)
			//fmt.Println("Date: ", date.UTC())
			if currentDay.Before(date) && date.Before(next) {
				i := sort.Search(len(resuults), func(i int) bool {
					return resuults[i].Date.Equal(currentDay)
				})

				if i < len(resuults) {

					//fmt.Println("total screenshots: ", resuults[i].Statistics.TotalScreenshots)

					resuults[i].Statistics.TotalScreenshots++
					if r.Value == r.ExpertAnswer {
						if resuults[i].Statistics.MatchWithExpert == -1 {
							resuults[i].Statistics.MatchWithExpert = 1
						} else {
							resuults[i].Statistics.MatchWithExpert++
						}
					}
					if r.Value == r.UsersAnswer {
						resuults[i].Statistics.RightAnswers++
					}
				} else {
					s := bo.StatisticsUsersInner{MatchWithExpert: -1}
					newR := bo.StatisticsUserBO{
						Date:       currentDay,
						Statistics: s,
					}

					resuults = append(resuults, newR)
				}
			}
		}
		currentDay = currentDay.AddDate(0, 0, 1)
	}

	//fmt.Printf("WATF: %+v\n", resuults)

	return resuults, nil
}
