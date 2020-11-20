CREATE DATABASE godb0;
CREATE USER alex WITH ENCRYPTED PASSWORD 'admin';
GRANT ALL PRIVILEGES ON DATABASE godb0 TO alex;
-- ##############################################

-- INFO: show all public database tables
SELECT table_name FROM information_schema.tables
WHERE table_schema = 'public';
-- ##############################################

-- INFO: test uuid extension
SELECT uuid_generate_v1();
-- ##############################################

-- INFO: drop tables
DROP TABLE games CASCADE;
DROP TABLE ext_systems CASCADE;
DROP TABLE sources CASCADE;
DROP TABLE screenshots CASCADE;
DROP TABLE users CASCADE;
DROP TABLE users__ext_systems CASCADE;
DROP TABLE answers CASCADE;
-- ##############################################