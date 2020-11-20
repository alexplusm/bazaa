package services

const (
	extSystemIDKey   = "extSystemID"
	screenshotsKey   = "--screenshots"
	gameKey          = "--game"
	screenshotURLKey = "url"

	// INFO: количество служебных полей: "url"
	nonAnswerFieldsCount = 1

	// INFO: дефолтное значение ответа
	initAnswerValue = "null"
)

var serviceKeyMap map[string]bool

func init() {
	serviceKeyMap = make(map[string]bool)

	serviceKeyMap[screenshotURLKey] = true
}
