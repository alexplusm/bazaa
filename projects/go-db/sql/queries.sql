CREATE DATABASE godb0;
CREATE USER alex WITH ENCRYPTED PASSWORD 'admin';
GRANT ALL PRIVILEGES ON DATABASE godb0 TO alex;

-- postgres://{user}:{password}@{hostname}:{port}/{database-name}
-- postgres://postgres:postgres@localhost:5432/testik

-- INFO: uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- INFO: test uuid extension
SELECT uuid_generate_v1();

-- games
-- create game
INSERT INTO games ("name", "start_date", "end_date", "answer_type", "question", "options_csv")
VALUES
(
    'first game!',
    '1601510400000',
    '1601683200000',
    '2',
    'Есть ли нарушение?',
    'Да,Нет'
)
RETURNING "game_id";
-- end games

-- sources
-- INSERT schedule (schedule has own ID)
INSERT INTO sources ("source_id", "source_type", "created_at", "game_id")
VALUES ('some-source-id', '1', '1601683200000', 'some-game-id')
RETURNING "source_id";
-- INSERT zip-archive (generate ID for archive source)
INSERT INTO sources ("source_type", "created_at", "game_id")
VALUES ('1', '1601683200000', 'some-game-id')
RETURNING "source_id";
-- end sources

-- external_systems
INSERT INTO external_systems (external_system_id)
VALUES ('ex-1');
-- external_systems end

-- users
INSERT INTO users (user_id)
VALUES ('vasya_pup123');
-- users end

-- TODO: into docs

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
