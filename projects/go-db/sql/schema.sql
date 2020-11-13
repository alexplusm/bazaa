-- INFO: sync schema with "docs/model_description_for_customer.md"
-- ####################################
-- INFO: to generate UML diagrams: https://app.sqldbm.com/
-- ####################################
-- INFO: большое количество constraints в виде FK могут навредить insert'ам
    -- в случае просадки скорости insert'ов рассмотреть вариант убрать часть FK
    -- "answers" table
-- ####################################
-- INFO: как работать с ответами понимаем из "game.answer_type"
-- answers.value | screenshots.expert_answer | screenshots.users_answer
-- *_answer
    -- if game.answerType == 1 (Текст) -> store 'user free text'
    -- if game.answerType == 2 (Категориальный) -> store option index  '0' | '1' | ...
    -- if game.answerType == 3 (Координаты прямоугольника) -> store JSON string
    -- if game.answerType == 4 (Полигональный) -> store JSON string

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS games (
	"game_id"           uuid        DEFAULT uuid_generate_v4(),
	"ext_system_id"     VARCHAR     NOT NULL,
	"name"              VARCHAR     NOT NULL,
	"start_date"        BIGINT      NOT NULL,
	"end_date"          BIGINT      NOT NULL,
	"answer_type"       INTEGER     NOT NULL,
	"question"          TEXT        NOT NULL,
	"options_csv"       VARCHAR     DEFAULT NULL, -- INFO: обязательное поле только для answer_type == 2

	PRIMARY KEY ("game_id")
);

CREATE TABLE IF NOT EXISTS ext_systems (
	"ext_system_id"         VARCHAR     DEFAULT uuid_generate_v4(),
	"description"           VARCHAR     NOT NULL,
	"post_results_url"      VARCHAR     NOT NULL,

	PRIMARY KEY ("ext_system_id")
);

CREATE TABLE IF NOT EXISTS sources (
    "source_id"     VARCHAR     DEFAULT uuid_generate_v4(),
    "game_id"       uuid        NOT NULL,
    "source_type"   INTEGER     NOT NULL,
    "created_at"    BIGINT      NOT NULL,

    PRIMARY KEY ("source_id"),

    FOREIGN KEY ("game_id") REFERENCES games("game_id")
);

CREATE TABLE IF NOT EXISTS screenshots (
	"screenshot_id"         uuid        DEFAULT uuid_generate_v4(),
	"game_id"               uuid        NOT NULL,
	"source_id"             VARCHAR     NOT NULL,
	"filename"              VARCHAR     NOT NULL,

	-- INFO: в этом поле будет лежать ответ эксперта, который мы получаем:
        -- либо при загрузке zip архива, из файловой структуры
        -- либо в ответе внешней системы, которая проверила "screenshot" и дала экспертный ответ
	"expert_answer"         VARCHAR     DEFAULT NULL,
	"users_answer"          VARCHAR     DEFAULT NULL,

	PRIMARY KEY ("screenshot_id"),

	FOREIGN KEY ("game_id")     REFERENCES games("game_id"),
	FOREIGN KEY ("source_id")   REFERENCES sources("source_id")
);

CREATE TABLE IF NOT EXISTS users (
	"inner_user_id"     uuid        DEFAULT uuid_generate_v4(),
	"user_id"           VARCHAR     NOT NULL,

	PRIMARY KEY ("user_id")
);

CREATE TABLE IF NOT EXISTS users__ext_systems (
	"ext_system_id"     VARCHAR     NOT NULL,
	"user_id"           VARCHAR     NOT NULL,

	FOREIGN KEY ("ext_system_id")   REFERENCES ext_systems("ext_system_id"),
	FOREIGN KEY ("user_id")         REFERENCES users("user_id")
);

CREATE TABLE IF NOT EXISTS answers (
	"answer_id"         BIGSERIAL,
	"screenshot_id"     uuid        NOT NULL,
	"game_id"           uuid        NOT NULL,
	"user_id"           VARCHAR     NOT NULL,
	"value"             VARCHAR     NOT NULL,

	PRIMARY KEY ("answer_id"),

	FOREIGN KEY ("screenshot_id")       REFERENCES screenshots("screenshot_id"),
	FOREIGN KEY ("game_id")             REFERENCES games("game_id"),
	FOREIGN KEY ("user_id")             REFERENCES users("user_id")
);

-- TODO: schedules (see docs)

-- INFO: drop tables
DROP TABLE games;
DROP TABLE ext_systems;
DROP TABLE sources;
DROP TABLE screenshots;
DROP TABLE users;
DROP TABLE users__ext_systems;
DROP TABLE answers;
