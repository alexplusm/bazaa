ENTITIES

- game
// How we will create new game?
// - proto: on the admin web-page must be form "create new game"
// - deep prod: automatically after load "job scheduler" from SOAP service. (???)
// How we will choose "game" while uploading zip with screenshots? 
// - In <input type="select"> on the admin web-page (where show game_id/game_start_timestamp)
attribs:
	* game_id // uuid
	* start_date // TIMESTAMP NOT NULL
	* answer_type (answer_type_enum)
	* O2M task (one game, many task)
    * question (+ options)
	

- external_system
// need form on the admin web-page for adding new external system?
attribs:
	* external_system_id
	* description
	* M2M user

- user
// if we receive request with new user and exsisting external_system -> create new user and join them
attribs:
	* user_id
	* M2M external_system

- task
attribs:
	* task_id
	* image_url
	* O2M game (one game, many task)


// TODO: Если существует запись ответа (answer) и предпринимается попытка создать новую запись, то:
-- Удалять старую и фиксировать новый ответ (update существующей запись с новым ответом)
-- Игнорировать новый ответ


- answer
attribs:
// должны уметь доставать все ответы на одно и тоже задание
	* O2O user
	* O2O task
	* answer (? for various games - various asnwer type !)
	* 

// NOTE
* how work with TIMESTAMP in postgresql
- new Date().toISOString() from JavasSript give correct date ISO string for postgresql!
CREATE TABLE test_ts (
	ts TIMESTAMP
);

INSERT INTO test_ts (ts)
VALUES ('2020-10-25T16:22:15.251Z')

SELECT * from test_ts;


