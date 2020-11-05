-- INFO: sync schema with "docs/model_description_for_customer.md"
-- INFO: to generate UML diagrams: https://app.sqldbm.com/

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TYPE IF EXISTS ANSWER_TYPE; -- INFO: 'IF NOT EXISTS' don't work with 'CREATE TYPE'
CREATE TYPE ANSWER_TYPE AS ENUM (
	'1', -- Текст
	'2', -- Категориальный
	'3', -- Координаты прямоугольника
	'4'  -- Полигональный
);

CREATE TABLE IF NOT EXISTS games (
	"game_id"       uuid            DEFAULT uuid_generate_v4(),
	"name"          VARCHAR         NOT NULL,
	"start_date"    BIGINT          NOT NULL,
	"end_date"      BIGINT          NOT NULL,
	"answer_type"   ANSWER_TYPE     NOT NULL,
	"question"      TEXT            NOT NULL,
	"options_csv"   VARCHAR         DEFAULT NULL, -- INFO: обязательное поле только для answer_type == 2

	PRIMARY KEY ("game_id")
);

-- TODO: schedules (see docs)

CREATE TABLE IF NOT EXISTS external_systems (
	"external_system_id"    VARCHAR DEFAULT uuid_generate_v4(),
	"post_task_results_url" VARCHAR DEFAULT NULL,

	PRIMARY KEY ("external_system_id")
);

CREATE TABLE IF NOT EXISTS users (
	"inner_user_id" uuid    DEFAULT uuid_generate_v4(),
	"user_id"       VARCHAR NOT NULL,

	PRIMARY KEY ("user_id")
);

CREATE TABLE IF NOT EXISTS users__external_systems (
	"external_system_id"    VARCHAR NOT NULL,
	"user_id"               VARCHAR NOT NULL,

	FOREIGN KEY ("external_system_id") REFERENCES external_systems("external_system_id"),
	FOREIGN KEY ("user_id") REFERENCES users("user_id")
);

-- TODO: tasks -> screenshots
CREATE TABLE IF NOT EXISTS tasks (
	"task_id" uuid DEFAULT uuid_generate_v4(),
	"game_id" uuid NOT NULL,
	"image_url" VARCHAR NOT NULL,
	-- INFO: в этом поле будет лежать ответ эксперта, который мы получаем:
		-- либо при загрузке zip архива, из файловой структуры
		-- либо в ответе внешней системы, которая проверила "task" и дала экспертный ответ
	-- INFO: как работать с ответами понимаем из "game.answer_type"

	-- INFO: *_answer
		-- if game.answerType == 1 (Текст) -> store 'user free text'
		-- if game.answerType == 2 (Категориальный) -> store option index  '0' | '1' | ...
		-- if game.answerType == 3 (Координаты прямоугольника) -> store JSON string
		-- if game.answerType == 4 (Полигональный) -> store JSON string
	"expert_answer" VARCHAR DEFAULT 'undefined',
	"users_answer" VARCHAR DEFAULT 'undefined', 
	-- INFO: как только 10 пользователей даст ответ, переводим в TRUE
	"complete_by_users" BOOLEAN DEFAULT FALSE,

	PRIMARY KEY ("task_id"),
	FOREIGN KEY ("game_id") REFERENCES games("game_id")
);

CREATE TABLE IF NOT EXISTS answers (
	"answer_id" BIGSERIAL,
	
	-- INFO: value
		-- if game.answerType == 1 (Текст) -> store 'user free text'
		-- if game.answerType == 2 (Категориальный) -> store option index  '0' | '1' | ...
		-- if game.answerType == 3 (Координаты прямоугольника) -> store JSON string
		-- if game.answerType == 4 (Полигональный) -> store JSON string
	"value" VARCHAR NOT NULL,
	"task_id" uuid NOT NULL,
	"game_id" uuid NOT NULL,
	"user_id" VARCHAR NOT NULL,
	"external_system_id" VARCHAR NOT NULL,

	PRIMARY KEY ("answer_id"),

	-- INFO: большое количество constraints в виде FK могут навредить insert'ам
		-- в случае просадки скорости insert'ов рассмотреть вариант убрать часть FK
	FOREIGN KEY ("task_id") REFERENCES tasks("task_id"), -- TODO: task -> screenshot
	FOREIGN KEY ("game_id") REFERENCES games("game_id"), -- TODO: above
	FOREIGN KEY ("user_id") REFERENCES users("user_id"),
	FOREIGN KEY ("external_system_id") REFERENCES external_systems("external_system_id")
);
