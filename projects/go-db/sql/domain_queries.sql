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
INSERT INTO external_systems ("external_system_id", "description", "post_results_url")
VALUES ('ex-1', 'Активный гражданин', 'https://hello.php');
--
INSERT INTO external_systems ("description", "post_results_url")
VALUES ('Активный гражданин', 'https://hello.php');
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
INSERT INTO users ("user_id")
VALUES ('vasya_pup123');

-- answers
INSERT INTO answers ("screenshot_id", "game_id", "user_id", "value")
VALUES ('screenshot-id-1', 'game-id-1', 'user-id-1', 'answer-value');

-- count of occurrences each screenshots in answers
SELECT screenshot_id, COUNT(screenshot_id) occurrences
FROM answers GROUP BY screenshot_id
-- ##############################################
