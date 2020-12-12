package services

import (
	"sort"
	"time"

	log "github.com/sirupsen/logrus"

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
	listDAO := make([]dao.AnswerRetrieve2DAO, 0, 1024)

	for _, gameID := range gameIDs {
		res, err := service.AnswerService.ABC(gameID, from, to)
		if err != nil {
			log.Error("get leaderboard statistic: ", err)
			continue
		}
		listDAO = append(listDAO, res...)
	}

	// MAP

	userStatisticMap := make(map[string]bo.StatisticsUsersInner)

	for _, listItem := range listDAO {
		_, exist := userStatisticMap[listItem.UserID]
		if !exist {
			userStatisticMap[listItem.UserID] = bo.StatisticsUsersInner{}
		}

		statistic := userStatisticMap[listItem.UserID]

		if listItem.Value == string(listItem.UsersAnswer) { // TODO: refactor
			statistic.RightAnswers++
		}
		if listItem.Value == string(listItem.ExpertAnswer) { // TODO: refactor
			statistic.MatchWithExpert++
		}
		statistic.TotalScreenshots++

		userStatisticMap[listItem.UserID] = statistic
	}

	leaders := make([]dto.LeadersDTO, 0, 1024)

	for key, item := range userStatisticMap {
		stats := dto.StatisticUsersInnerDTO{
			TotalScreenshots: item.TotalScreenshots,
			RightAnswers:     item.RightAnswers,
			MatchWithExpert:  item.MatchWithExpert,
			AverageAccuracy:  item.AverageAccuracy,
		}
		leaders = append(leaders, dto.LeadersDTO{UserID: key, Statistics: stats})
	}

	sort.Slice(leaders, func(i, j int) bool {
		return leaders[i].Statistics.RightAnswers > leaders[j].Statistics.RightAnswers
	})

	leaders = leaders[:limit]

	return dto.LeadersResponseDTO{Leaders: leaders}
}
