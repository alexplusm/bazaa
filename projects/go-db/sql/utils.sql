CREATE DATABASE godb0;
CREATE USER alex WITH ENCRYPTED PASSWORD 'admin';
GRANT ALL PRIVILEGES ON DATABASE godb0 TO alex;

-- INFO: show all public database tables
SELECT table_name FROM information_schema.tables
WHERE table_schema = 'public';

-- INFO: drop tables
DROP TABLE games;
DROP TABLE ext_systems;
DROP TABLE sources;
DROP TABLE screenshots;
DROP TABLE users;
DROP TABLE users__ext_systems;
DROP TABLE answers;