CREATE DATABASE godb0;
CREATE USER alex WITH ENCRYPTED PASSWORD 'admin';
GRANT ALL PRIVILEGES ON DATABASE godb0 TO alex;

-- postgres://{user}:{password}@{hostname}:{port}/{database-name}
-- postgres://postgres:postgres@localhost:5432/testik

-- INFO: show all database tables
SELECT table_name FROM information_schema.tables
WHERE table_schema = 'public';

-- INFO: uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- INFO: test uuid extension
SELECT uuid_generate_v1();

-- games
-- create game
INSERT INTO games ("ext_system_id", "name", "start_date", "end_date", "answer_type", "question", "options_csv")
VALUES
(
    'ext-system-id-123',
    'first game!',
    '1601510400000',
    '1601683200000',
    '2',
    'Есть ли нарушение?',
    'Да,Нет'
)
RETURNING "game_id";
-- ##############################################

-- ext_systems
INSERT INTO external_systems (external_system_id)
VALUES ('ex-1');
-- ##############################################

-- sources
-- insert schedule (schedule has own ID)
INSERT INTO sources ("source_id", "source_type", "created_at", "game_id")
VALUES ('some-source-id', '1', '1601683200000', 'some-game-id')
RETURNING "source_id";

-- insert zip-archive (generate ID for archive source)
INSERT INTO sources ("source_type", "created_at", "game_id")
VALUES ('1', '1601683200000', 'some-game-id')
RETURNING "source_id";
-- ##############################################

-- screenshots
-- insert without expert answer
INSERT INTO screenshots ("game_id", "source_id", "filename")
VALUES ("some-game-id", "some-source-id", "filename.jpg");

-- insert with expert answer
INSERT INTO screenshots ("game_id", "source_id", "filename", "expert_answer")
VALUES ("some-game-id", "some-source-id", "filename.jpg", "1");
-- ##############################################

-- users
INSERT INTO users (user_id)
VALUES ('vasya_pup123');
-- ##############################################

-- TODO: into docs
--     ExternalSystemId
-- Date1
-- Date2
--

