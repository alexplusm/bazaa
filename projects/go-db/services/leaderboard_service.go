package services

import (
	"fmt"
	"sort"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type LeaderboardService struct {
	AnswerService interfaces.IAnswerService
}

func (service *LeaderboardService) GetLeaderboard(gameIDs []string, from, to time.Time, limit int) dto.LeadersResponseDTO {
	// TODO: refactor
	listDAO := make([]dao.AnswerStatLeadDAO, 0, 1024)

	for _, gameID := range gameIDs {
		res, err := service.AnswerService.ABC(gameID, from, to)
		if err != nil {
			// TODO: LOG
			continue
		}
		fmt.Println("Res: ", res)
		listDAO = append(listDAO, res...)
	}

	// MAP

	userAnswerMap := make(map[string]bo.StatisticsUsersInner)

	for _, listItem := range listDAO {
		_, ok := userAnswerMap[listItem.UserID]
		if !ok {
			userAnswerMap[listItem.UserID] = bo.StatisticsUsersInner{}
		}

		val := userAnswerMap[listItem.UserID]

		if listItem.Value == listItem.UsersAnswer {
			val.RightAnswers++
			//userAnswerMap[listItem.UserID].RightAnswers++
		}
		if listItem.ExpertAnswer == listItem.Value {
			val.MatchWithExpert++
			//userAnswerMap[listItem.UserID].MatchWithExpert++
		}
		val.TotalScreenshots++
		//userAnswerMap[listItem.UserID].TotalScreenshots++

		userAnswerMap[listItem.UserID] = val
	}

	leaders := make([]dto.LeadersDTO, 0, 1024)

	for key, item := range userAnswerMap {
		sts := dto.StatisticUsersInnerDTO{
			TotalScreenshots: item.TotalScreenshots,
			RightAnswers:     item.RightAnswers,
			MatchWithExpert:  item.MatchWithExpert,
			AverageAccuracy:  item.AverageAccuracy,
		}
		dtooo := dto.LeadersDTO{UserID: key, Statistics: sts}
		leaders = append(leaders, dtooo)
	}

	sort.Slice(leaders, func(i, j int) bool {
		return leaders[i].Statistics.RightAnswers > leaders[j].Statistics.RightAnswers
	})

	leaders = leaders[:limit]

	resp := dto.LeadersResponseDTO{Leaders: leaders}

	return resp
}
