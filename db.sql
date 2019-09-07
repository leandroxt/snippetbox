CREATE DATABASE
IF NOT EXISTS snippetbox
WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

CREATE TABLE
IF NOT EXISTS snippets
(
	id SERIAL PRIMARY KEY,
	title VARCHAR
(100) NOT NULL,
	content TEXT NOT NULL,
	created DATE NOT NULL,
	expires DATE NOT NULL
);

-- Add an index on the created column.
CREATE INDEX
IF NOT EXISTS idx_snippets_created ON snippets
(created);


INSERT INTO snippets
	(title, content, created, expires)
VALUES
	(
		'An old silent pond',
		'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
		now(),
		now() + INTERVAL
'365' DAY
);

INSERT INTO snippets
	(title, content, created, expires)
VALUES
	(
		'Over the wintry forest',
		'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
		now(),
		now() + INTERVAL
'365' DAY
);

INSERT INTO snippets
	(title, content, created, expires)
VALUES
	(
		'First autumn morning',
		'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
		now(),
		now() + INTERVAL
'7' DAY
);

CREATE ROLE snippetbox
with PASSWORD 'Password' LOGIN;

GRANT CONNECT ON DATABASE snippetbox TO snippetbox;
GRANT USAGE ON SCHEMA public TO snippetbox;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO snippetbox;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO snippetbox;

