CREATE DATABASE godb0;
CREATE USER alex WITH ENCRYPTED PASSWORD 'admin';
GRANT ALL PRIVILEGES ON DATABASE godb0 TO alex;

-- postgres://{user}:{password}@{hostname}:{port}/{database-name}
-- postgres://postgres:postgres@localhost:5432/testik

-- to use uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- to test extension
SELECT uuid_generate_v1();


-- create table and use its PK in images__category:
-- image_category
-- -- no_violation
-- -- violation
-- -- undefined
-- //////
-- create table images
CREATE TABLE IF NOT EXISTS images (
	image_id uuid DEFAULT uuid_generate_v4 (),
    "url" VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    -- tasks (задание и чет еще)

    -- был ли дан ответ (+ поле самого ответа (мб использовать enum category))
    resolved BOOLEAN DEFAULT false, 
	PRIMARY KEY (image_id)
);

-- insert image
INSERT INTO images ("url", category)
VALUES ('/path/kek/123123123.png', 'no_violation');
