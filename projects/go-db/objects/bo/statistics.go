package bo

import (
	"strconv"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type StatisticsUserBO struct {
	Date       time.Time
	Statistics StatisticsUsersInner
}

type StatisticsUsersInner struct {
	TotalScreenshots int
	RightAnswers     int
	MatchWithExpert  int
	AverageAccuracy  float64
}

func StatsToDTO(stats []StatisticsUserBO) dto.StatsUserDTO {
	total := dto.StatisticsUsersInnerDTO{}
	history := make([]dto.StatisticsUserDTO, 0, len(stats))

	for _, s := range stats {
		d := dto.StatisticsUserDTO{}
		d.Date = strconv.FormatInt(s.Date.Unix(), 10)
		d.Statistics.MatchWithExpert = s.Statistics.MatchWithExpert
		d.Statistics.RightAnswers = s.Statistics.RightAnswers
		d.Statistics.TotalScreenshots = s.Statistics.TotalScreenshots
		d.Statistics.AverageAccuracy = s.Statistics.AverageAccuracy

		// TODO: не место!!!
		d.Statistics.AverageAccuracy = float64(d.Statistics.RightAnswers) / float64(d.Statistics.TotalScreenshots)

		history = append(history, d)

		total.TotalScreenshots += s.Statistics.TotalScreenshots
		total.RightAnswers += s.Statistics.RightAnswers
		total.MatchWithExpert += s.Statistics.MatchWithExpert
	}
	total.AverageAccuracy = float64(total.RightAnswers) / float64(total.TotalScreenshots)

	return dto.StatsUserDTO{Total: total, History: history}
}

func StatsToTotalOnlyDTO(stats []StatisticsUserBO) dto.StatsUserTotalOnlyDTO {
	total := dto.StatisticsUsersInnerDTO{}

	for _, s := range stats {
		total.TotalScreenshots += s.Statistics.TotalScreenshots
		total.RightAnswers += s.Statistics.RightAnswers
		total.MatchWithExpert += s.Statistics.MatchWithExpert
	}
	total.AverageAccuracy = float64(total.RightAnswers) / float64(total.TotalScreenshots)

	return dto.StatsUserTotalOnlyDTO{Total: total}
}
