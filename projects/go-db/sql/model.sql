DROP TYPE IF EXISTS ANSWER_TYPE; -- 'IF NOT EXISTS' don't work with 'CREATE TYPE'
CREATE TYPE ANSWER_TYPE AS ENUM (
	'1', 
	'2',
	'3',
	'4'
);
-- TODO: add descriptions for each type

-- delete
-- options_list
CREATE TABLE IF NOT EXISTS options_list (
	"options_id" SERIAL,
	"options_csv" TEXT NOT NULL,
	
	PRIMARY KEY ("options_id")
);

INSERT INTO options_list ("options_csv")
	VALUES ('yes, no');

INSERT INTO options_list ("options_csv")
	VALUES ('yes, no, nope, yep') RETURNING "options_id"; -- return "options_id" after insert
-- end options_list

-- games
CREATE TABLE IF NOT EXISTS games (
	"game_id" uuid DEFAULT uuid_generate_v4(),
	"start_date" TIMESTAMP NOT NULL,
	"end_date" TIMESTAMP NOT NULL,
	"answer_type" ANSWER_TYPE NOT NULL,
	"question" TEXT NOT NULL,
	"options_csv" VARCHAR, -- не обязательный

	PRIMARY KEY ("game_id"),
);

INSERT INTO games ("start_date", "answer_type", "question", "options_id")
	VALUES
	('2020-10-25T17:30:39.417Z', '1', 'some question', 1);

-- получить список options связанных со своими game
SELECT * FROM options_list
	INNER JOIN games ON options_list.options_id = games.options_id;


-- создать опции и игру вместе
WITH insert_result AS (
	INSERT INTO options_list ("options_csv")
		VALUES ('yes, no') RETURNING "options_id"
)
INSERT INTO games ("start_date", "answer_type", "question", "options_id")
	VALUES
	(
		'2020-10-25T17:30:39.417Z', 
		'1', 
		'some question',
		(SELECT options_id FROM insert_result)
	)
	RETURNING "game_id";


-- end games

-- external_systems
CREATE TABLE IF NOT EXISTS external_systems (
	-- "external_system_id" VARCHAR NOT NULL,
	"external_system_id" uuid DEFAULT uuid_generate_v4(),
	"post_task_results_url" VARCHAR,
	"description" TEXT, -- TODO: тужно ли нам это поле?

	PRIMARY KEY ("external_system_id")
);

INSERT INTO external_systems (external_system_id, "description")
	VALUES ('ex-1', 'desc1');
-- external_systems end

-- users
CREATE TABLE IF NOT EXISTS users (
	"inner_user_id" uuid DEFAULT uuid_generate_v4(),
	"user_id" VARCHAR NOT NULL,

	PRIMARY KEY ("user_id")
);

INSERT INTO users (user_id)
	VALUES ('vasya_pup123');
-- users end

-- users__external_systems
CREATE TABLE IF NOT EXISTS users__external_systems (
	"external_system_id" VARCHAR NOT NULL,
	"user_id" VARCHAR NOT NULL,

	FOREIGN KEY ("external_system_id") REFERENCES external_systems("external_system_id"),
	FOREIGN KEY ("user_id") REFERENCES users("user_id")
);
-- users__external_systems end

-- tasks
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
-- tasks end

-- answers
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
	FOREIGN KEY ("task_id") REFERENCES tasks("task_id"),
	FOREIGN KEY ("game_id") REFERENCES games("game_id"), -- TODO: above
	FOREIGN KEY ("user_id") REFERENCES users("user_id"),
	FOREIGN KEY ("external_system_id") REFERENCES external_systems("external_system_id")
);
-- answers end


ExternalSystemId 
Date1
Date2

У кого больше правильный ответов - тот и лидер

Select game_id FROM games
WHERE games.start_date >= Date1 OR games.start_date <= Date2
INNER JOIN answers
ON answers.game_id = game_id
WHERE answers.external_system_id = ExternalSystemId
