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

-- rename table column
ALTER TABLE images 
RENAME COLUMN category TO init_category;




-- Finish system

-- Эта таблица заполняется "вручную" / хардкодится заранее?
-- Если в запросе отправляется id внешней системы, который нет у нас в таблице, 
-- 		то мы возвращаем ошибку?
CREATE TABLE IF NOT EXISTS eternal_systems (
	id uuid NOT NULL, -- id VARCHAR NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
	-- if request with new user - create new user and user_tasks instances
	-- use user_id in request query_params
	id VARCHAR NOT NULL, -- ? id uuid NOT NULL,
	-- tasks (O2M)
	-- ? eternal system (O2M)
	
	-- O2M eternal_systems
	
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_tasks (
	id SERIAL,
	user_id uuid NOT NULL,
	task_id uuid NOT NULL,
	
	-- or not need, if we will store each user answer and, if need, will culc total_earned
	total_earned INTEGER DEFAULT 0, -- after answering undate (increase if right answer)
	
	PRIMARY KEY (id),
	FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS user_answers (
	id SERIAL,
	user_id uuid NOT NULL,
	task_id uuid NOT NULL,
	
	-- user_answer
	
	PRIMARY KEY (id),
	FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)

CREATE TYPE image_category AS ENUM (
	'no_violation',
	'violation',
	'undefined'
)

CREATE TABLE IF NOT EXISTS tasks (
	id uuid DEFAULT uuid_generate_v4 (),
	"image_url" VARCHAR NOT NULL,
	
	init_category image_category, -- ?
	users_answer_category image_category DEFAULT 'undefined', -- ?

	job_id uuid NOT NULL, -- (O2M)

	PRIMARY KEY (id),
	FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE,
);


CREATE TABLE IF NOT EXISTS jobs (
	id uuid DEFAULT uuid_generate_v4 (),
	load_timestamp, -- create_at
	job_date, -- TIMESTAMP

	PRIMARY KEY (id)
)

-- todo: rename jobs -> tasks; tasks -> ???
