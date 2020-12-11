package bo

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/timeutils"
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

// ---

type StatisticAnswersDateSlicedBO struct {
	Date       time.Time
	Statistics StatisticAnswersBO
}

type StatisticAnswersBO struct {
	TotalScreenshots int
	RightAnswers     int
	MatchWithExpert  int
	AverageAccuracy  float64
}

// todo
func (s *StatisticAnswersBO) Increase(answer, expertAnswer, usersAnswer string) {
	s.TotalScreenshots++

	if usersAnswer == answer {
		s.RightAnswers++
	}

	// TODO: write doc
	if expertAnswer != "" {
		if s.MatchWithExpert != -1 {
			if answer == expertAnswer {
				s.MatchWithExpert++
			}
		} else {
			if answer == expertAnswer {
				s.MatchWithExpert = 1
			}
		}
	}
}

func StatisticAnswersDateSlicedBOToDTOTotalOnly(stats []StatisticAnswersDateSlicedBO) dto.StatisticUserAnswersTotalDTO {
	total := dto.StatisticUsersInnerDTO{}
	total.AverageAccuracy = 0

	for _, s := range stats {
		total.TotalScreenshots += s.Statistics.TotalScreenshots
		total.RightAnswers += s.Statistics.RightAnswers
		total.MatchWithExpert += s.Statistics.MatchWithExpert
	}

	if total.TotalScreenshots != 0 {
		total.AverageAccuracy = float64(total.RightAnswers) / float64(total.TotalScreenshots)
	}

	return dto.StatisticUserAnswersTotalDTO{Total: total}
}

func StatisticAnswersDateSlicedBOToDTO(stats []StatisticAnswersDateSlicedBO) dto.StatisticUserAnswersDTO {
	total := dto.StatisticUsersInnerDTO{}
	history := make([]dto.StatisticsUserDTO, 0, len(stats))

	for _, s := range stats {
		d := dto.StatisticsUserDTO{}
		d.Date = timeutils.FromTimeToStrTimestamp(s.Date)
		d.Statistics.MatchWithExpert = s.Statistics.MatchWithExpert
		d.Statistics.RightAnswers = s.Statistics.RightAnswers
		d.Statistics.TotalScreenshots = s.Statistics.TotalScreenshots
		d.Statistics.AverageAccuracy = s.Statistics.AverageAccuracy

		// TODO: не место!!!
		if d.Statistics.TotalScreenshots != 0 {
			d.Statistics.AverageAccuracy = float64(d.Statistics.RightAnswers) / float64(d.Statistics.TotalScreenshots)
		} else {
			d.Statistics.AverageAccuracy = 0
		}

		history = append(history, d)

		total.TotalScreenshots += s.Statistics.TotalScreenshots
		total.RightAnswers += s.Statistics.RightAnswers
		if s.Statistics.MatchWithExpert != -1 {
			total.MatchWithExpert += s.Statistics.MatchWithExpert
		}
	}
	if total.TotalScreenshots != 0 {
		total.AverageAccuracy = float64(total.RightAnswers) / float64(total.TotalScreenshots)
	} else {
		total.AverageAccuracy = 0
	}

	val := dto.StatisticUserAnswersDTO{Total: total, History: history}

	return val
}
