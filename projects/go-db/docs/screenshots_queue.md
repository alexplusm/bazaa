### tech description

---
$gameID : $externalSystemID
"screenshot:$gameID" :list: [scrID1, scrID2, scrID3, ... scrID_n]
---
Когда отдаем задание пользователю с $userID
		создаем поле в структуре $scrID1->$userID со значением "null"
Когда получаем ответ от пользователя с $userID
		обновляем поле в структуре $scrID1->$userID со значением ответа
Если скриншот был выдан уже 10-ти пользователям: то мы выдаем следующий скриншот
		проверяем кол-во ключей в хэше (исключая поле "url") // из константы
Когда все ответы по скриншоту получены
		* проверка в момент получения ответа от пользователя
		1) проверяем кол-во ключей в хэше (исключая поле "url") // из константы
		2) у всех полей значение не "null"
		3) записываем в базу все ответы всех пользователей
---
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
