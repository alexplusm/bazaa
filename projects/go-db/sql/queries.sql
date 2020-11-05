CREATE DATABASE godb0;
CREATE USER alex WITH ENCRYPTED PASSWORD 'admin';
GRANT ALL PRIVILEGES ON DATABASE godb0 TO alex;

-- postgres://{user}:{password}@{hostname}:{port}/{database-name}
-- postgres://postgres:postgres@localhost:5432/testik

-- to use uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- to test extension
SELECT uuid_generate_v1();

CREATE TYPE image_category AS ENUM (
	'no_violation',
	'violation',
	'undefined'
)

-- create table images
CREATE TABLE IF NOT EXISTS images (
	id uuid DEFAULT uuid_generate_v4 (),
	"url" VARCHAR NOT NULL,
	init_category image_category,
	users_answer_category image_category DEFAULT 'undefined',

	-- task_id (o-to-many)

	-- был ли дан ответ (+ поле самого ответа (мб использовать enum category))
	resolved BOOLEAN DEFAULT false, 
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tasks (
	-- task_id SERIAL,
	id uuid DEFAULT uuid_generate_v4 (),
	-- timestamp -- значение времени, когда был загружен архив на сервер
	PRIMARY KEY (id)
)

-- insert image
INSERT INTO images ("url", category)
VALUES ('/path/kek/123123123.png', 'no_violation');

-- TODO
CREATE TYPE image_category AS ENUM (
	'no_violation',
	'violation',
	'undefined'
)

CREATE TABLE IF NOT EXISTS jobs (
	id uuid DEFAULT uuid_generate_v4 (),
	load_timestamp, -- TODO: create_at (schedules)
	job_date, -- TODO: TIMESTAMP (schedules)

	PRIMARY KEY (id)
)

-- todo: rename jobs -> tasks; tasks -> ???


--     ExternalSystemId
-- Date1
-- Date2
--
-- У кого больше правильный ответов - тот и лидер
--
-- Select game_id FROM games
-- WHERE games.start_date >= Date1 OR games.start_date <= Date2
--     INNER JOIN answers
-- ON answers.game_id = game_id
-- WHERE answers.external_system_id = ExternalSystemId

-- games
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
INSERT INTO external_systems (external_system_id)
VALUES ('ex-1');
-- external_systems end

-- users
INSERT INTO users (user_id)
VALUES ('vasya_pup123');
-- users end
