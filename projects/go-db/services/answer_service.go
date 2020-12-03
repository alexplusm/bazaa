package services

import (
	"fmt"
	"sort"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils"
)

type AnswerService struct {
	AnswerRepo interfaces.IAnswerRepository
}

func (service *AnswerService) GetScreenshotResults(
	gameID, screenshotID string,
) ([]dto.UserAnswerForScreenshotResultDTO, error) {
	res, err := service.AnswerRepo.SelectScreenshotResult(gameID, screenshotID)
	if err != nil {
		return nil, fmt.Errorf("get screenshot results: %v", err)
	}

	res_len := len(res)
	list := make([]dto.UserAnswerForScreenshotResultDTO, 0, res_len)

	for _, r := range res {
		dtoo := dto.UserAnswerForScreenshotResultDTO{}
		dtoo.UserID = r.UserID
		dtoo.Answer = r.Value
		if res_len < consts.RequiredAnswerCountToFinishScreenshot {
			dtoo.Result = "inProcess"
		} else {
			if r.UsersAnswer == "-1" {
				dtoo.Result = "undefined"
			} else {
				if r.Value == r.UsersAnswer {
					dtoo.Result = "right"
				} else {
					dtoo.Result = "wrong"
				}
			}
		}
		list = append(list, dtoo)
	}

	fmt.Printf("res: %+v\n", res)

	return list, nil
}

// TODO: refactor | screenshotResult (добавить поле usersResult)
func (service *AnswerService) GetUserStatistics(
	userID string, gameIDs []string, from, to time.Time,
) ([]bo.StatisticsUserBO, error) {
	res, err := service.AnswerRepo.SelectAnswersByUser(userID, gameIDs, from, to)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, fmt.Errorf("get user statistics: %v", err)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i].AnswerDate < res[j].AnswerDate
	})

	start := utils.TrimTime(time.Unix(res[0].AnswerDate, 0))
	end := time.Unix(res[len(res)-1].AnswerDate, 0)

	results := make([]bo.StatisticsUserBO, 0, len(res))

	currentDay := start
	for currentDay.Before(end) {
		for _, r := range res {

			date := time.Unix(r.AnswerDate, 0)
			next := currentDay.AddDate(0, 0, 1)

			if currentDay.Before(date) && date.Before(next) {
				i := sort.Search(len(results), func(i int) bool {
					return results[i].Date.Equal(currentDay)
				})

				if i < len(results) {
					results[i].Statistics.TotalScreenshots++
					if r.Value == r.ExpertAnswer {
						if results[i].Statistics.MatchWithExpert == -1 {
							results[i].Statistics.MatchWithExpert = 1
						} else {
							results[i].Statistics.MatchWithExpert++
						}
					}
					if r.Value == r.UsersAnswer {
						results[i].Statistics.RightAnswers++
					}
				} else {
					s := bo.StatisticsUsersInner{MatchWithExpert: -1}
					newR := bo.StatisticsUserBO{
						Date:       currentDay,
						Statistics: s,
					}
					results = append(results, newR)
				}
			}
		}
		currentDay = currentDay.AddDate(0, 0, 1)
	}

	return results, nil
}

func (service *AnswerService) GetUsersAndScreenshotCountByGame(
	gameID string,
) (dao.AnsweredScreenshotsDAO, error) {
	return service.AnswerRepo.SelectAnsweredScreenshotsByGame(gameID)
}

func (service *AnswerService) ABC(gameID string, from, to time.Time) ([]dao.AnswerStatLeadDAO, error) {
	return service.AnswerRepo.SelectAnswersTODO(gameID, from, to)
}
