package services

import (
	"context"
	"fmt"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"strings"
)

// TODO: move to consts
const (
	extSystemIDKey = "extSystemID"
	screenshotsKey = "screenshots"
	answerCount    = 5 // TODO: rename + move to config
	// INFO: "url" ...
	serviceFields = 1
)

type RedisService struct {
	RedisClient interfaces.IRedisHandler
}

func (r *RedisService) Method() {
	gameID := "bd255325-e7d1-44bd-8f76-1ff4796e71a2"

	//ctx := context.Background()
	//conn := r.RedisClient.GetConn()

	// CHECK GAME <-> EXTERNAL_SYSTEM_ID
	//sysID := "ExtSys-123-id"
	//e := r.checkGame(gameID, sysID)
	//fmt.Println("1) Has GAME!", e)
	//r.insertGame(gameID, sysID)
	//e = r.checkGame(gameID, sysID)
	//fmt.Println("2) Has GAME!", e)
	// CHECK GAME <-> EXTERNAL_SYSTEM_ID END

	// --------------------------------------
	//r.insertScreenshots(gameID, screenshots)
	//keys, err := conn.Keys(ctx, "*").Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("KEYS: ", keys)
	//fmt.Println("LEN KEYS:", len(keys))
	//fmt.Println("-----")
	//for _, key := range keys {
	//	v, err := conn.HGet(ctx, key, "url").Result()
	//	fmt.Println("###", key, "|", v, err)
	//}
	//fmt.Println(len(screenshots))
	// --------------------------------------

	hasID, id := r.getScreenshotID(gameID)
	fmt.Println("===", hasID, id)
}

// TODO: remove
//func (r *RedisService) prepareCache(gameID string, screenshots []dao.ScreenshotDAOFull) {
//	// достаем последнюю игру (за один час до начала игры) | todo: а если несколько игр?!
//	// достаем все скрины из этой игры | todo: замешивание из прошлой игры с таким же типом
//	// закидываем в редис имена скриншотов (в массив?!)
//
//	y := make([]interface{}, len(screenshots))
//	for i := range screenshots {
//		y[i] = screenshots[i].ScreenshotID
//	}
//
//	c := r.RedisClient.GetConn()
//	ctx := context.Background()
//
//	val, err := c.RPush(ctx, screenshotsKey, y...).Result()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println("Value: ", val)
//
//	len, err := c.LLen(ctx, screenshotsKey).Result()
//
//	fmt.Println("Len: ", len)
//	values, err := c.LRange(ctx, screenshotsKey, 0, len).Result()
//
//	fmt.Println("VALUES", values)
//}

func (r *RedisService) insertGame(gameID, externalSystemID string) {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

	conn.HSet(ctx, gameID, extSystemIDKey, externalSystemID)
}

func (r *RedisService) insertScreenshots(gameID string, screenshots []dao.ScreenshotDAOFull) {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

	screenshotIDs := make([]interface{}, len(screenshots))
	for i := range screenshots {
		id := screenshots[i].ScreenshotID
		conn.HSet(ctx, id, "url", screenshots[i].Filename) // TODO: build URL
		screenshotIDs[i] = id
	}

	conn.LPush(ctx, getScreenshotsListKey(gameID), screenshotIDs...)
}

// скриншот, который можно отдать пользователю
func (r *RedisService) getScreenshotID(gameID string) (bool, string) {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

	listLength, err := conn.LLen(ctx, getScreenshotsListKey(gameID)).Result()
	var index int64
	var id string
	var hasID = true
	flag := true

	for flag {
		// todo: incr index to find not full answered screenshot
		if index+1 == listLength {
			hasID = false
			id = ""
			break
		}
		id, err = conn.LIndex(ctx, getScreenshotsListKey(gameID), index).Result()

		if err != nil {
			fmt.Println(err)
		}
		flag = r.checkAnsweredScreenshot(id)
		index++
	}

	fmt.Println(err) // TODO: process error

	return hasID, id
}

// TODO: check out of range

//fmt.Println("IDs: ", id)

//conn.HSet(ctx, id, "user1", "kek")
//conn.HSet(ctx, id, "user2", "lol")
//conn.HSet(ctx, id, "user3", "kek")
//conn.HSet(ctx, id, "user4", "kek")

//fmt.Println("BOOL:", r.checkAnsweredScreenshot(id))

//url := conn.HGet(ctx, id, "url").Val()
//fmt.Println("Url:", url)

func (r *RedisService) checkAnsweredScreenshot(screenshotID string) bool {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

	keys, err := conn.HKeys(ctx, screenshotID).Result()
	fmt.Println("Keys:", keys)

	fmt.Println(err) // TODO: process error
	// INFO: проверка длинны массива ключей ответов с учетом наличия поля служебных полей
	return len(keys) < answerCount+serviceFields // TODO: concurrency
}

func getScreenshotsListKey(gameID string) string {
	return strings.Join([]string{screenshotsKey, gameID}, ":")
}

/*
$gameID : $externalSystemID
"screenshot:$gameID" :list: [scrID1, scrID2, scrID3, ... scrID_n]

// Когда отдаем задание пользователю с $userID
// 		создаем поле в структуре $scrID1->$userID со значением "null"
// Когда получаем ответ от пользователя с $userID
//		обновляем поле в структуре $scrID1->$userID со значением ответа
// Если скриншот был выдан уже 10-ти пользователям: то мы выдаем следующий скриншот
// 		проверяем кол-во ключей в хэше (исключая поле "url") // из константы
// Когда все ответы по скриншоту получены
//		* проверка в момент получения ответа от пользователя
// 		1) проверяем кол-во ключей в хэше (исключая поле "url") // из константы
//		2) у всех полей значение не "null"
// 		3) записываем в базу все ответы всех пользователей

$scrID_1 :hash:
	url: /url/abc_1.jpg
	$userID1: answer1 | null
	$userID2: answer2 | null
	...
	$userID10: answer10 | null
...
$scrID_n :hash:
	url: /url/abc_n.jpg
	$userID1: answer1 | null
	$userID2: answer2 | null
	...
	$userID10: answer10 | null
*/
