package services

import (
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dto"
)

type ScreenshotUserAnswerService struct {
	AnswerRepo     interfaces.IAnswerRepo
	ScreenshotRepo interfaces.IScreenshotRepo
}

func (service *ScreenshotUserAnswerService) BuildUserAnswerResponse(
	userID string, answersBO []bo.UserAnswerCacheBO,
) dto.UserAnswerResponseData {
	// TODO: для разных типов игр по разному вычисляем "Result"
	// TODO: refactor!!!
	answers := make([]dto.UserAnswerDTO, 0, len(answersBO))
	finished := service.ScreenshotIsFinished(answersBO)

	if finished {
		rightAnswer, answerDefined := getRightAnswerCategoryType(answersBO)

		if !answerDefined {

			for _, answer := range answersBO {
				answerDTO := dto.UserAnswerDTO{
					UserID: answer.UserID, Answer: answer.Answer, Result: consts.UserResultUndefined,
				}
				answers = append(answers, answerDTO)
			}

			return dto.UserAnswerResponseData{
				Finished: finished, UserResult: consts.UserResultUndefined, Answers: answers,
			}

		} else {

			var currentUserResult string

			for _, answer := range answersBO {

				var userResult string
				if rightAnswer == answer.Answer {
					userResult = consts.UserResultRight
				} else {
					userResult = consts.UserResultWrong
				}

				if userID == answer.UserID {
					currentUserResult = userResult
				}

				answerDTO := dto.UserAnswerDTO{
					UserID: answer.UserID, Answer: answer.Answer, Result: userResult,
				}
				answers = append(answers, answerDTO)
			}

			return dto.UserAnswerResponseData{
				Finished: finished, UserResult: currentUserResult, Answers: answers,
			}
		}
	}

	for _, answer := range answersBO {
		answerDTO := dto.UserAnswerDTO{
			UserID: answer.UserID, Answer: answer.Answer, Result: consts.UserResultInProcess,
		}
		answers = append(answers, answerDTO)
	}

	return dto.UserAnswerResponseData{
		Finished: finished, UserResult: consts.UserResultInProcess, Answers: answers,
	}
}

func (service *ScreenshotUserAnswerService) ScreenshotIsFinished(
	answers []bo.UserAnswerCacheBO,
) bool {
	return len(answers) == consts.RequiredAnswerCountToFinishScreenshot
}

func (service *ScreenshotUserAnswerService) SaveUsersAnswers(
	answers []bo.UserAnswerCacheBO, gameID, screenshotID string,
) {
	answersCountMap := make(map[string]int)
	// TODO: config?! "-1": default usersAnswer
	// (когда пользователи не выбрали ответ достаточно однозначно)
	//  < 7 одинаковых ответов
	usersAnswer := "-1"
	answersDAO := make([]dao.AnswerInsertDAO, 0, len(answers))
	for _, answer := range answers {
		answerDAO := dao.AnswerInsertDAO{}
		answerDAO.FromCacheBO(answer, gameID, screenshotID)
		answersCountMap[answerDAO.Value]++
		answersDAO = append(answersDAO, answerDAO)
	}

	for key, value := range answersCountMap {
		if value >= consts.RightAnswerThreshold {
			usersAnswer = key
		}
	}
	err := service.ScreenshotRepo.UpdateUsersAnswer(screenshotID, usersAnswer)
	if err != nil {
		log.Error("save users answer: ", err)
	}
	service.AnswerRepo.InsertList(answersDAO)
}

func getRightAnswerCategoryType(answers []bo.UserAnswerCacheBO) (string, bool) {
	answerMap := make(map[string]int)

	for _, answer := range answers {
		answerMap[answer.Answer]++
	}
	for key, item := range answerMap {
		if item >= consts.RightAnswerThreshold {
			return key, true
		}
	}

	return "", false
}
