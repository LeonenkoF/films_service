-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
       id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       username UUID NOT NULL UNIQUE,
       role BOOLEAN NOT NULL,
       password VARCHAR(32)
    );

CREATE TABLE IF NOT EXISTS actors(
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	name TEXT NOT NULL UNIQUE,
	gender VARCHAR(1) NOT NULL,
	birth TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS movie(
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	title TEXT NOT NULL,
	description TEXT,
	release_date TEXT NOT NULL,
	rating INT NOT NULL,
	CHECK (length(title) >= 1 AND length(title) <= 150),
	CHECK (length(description) <= 1000),
	CHECK (rating >= 0 AND rating <= 10)
);
CREATE INDEX IF NOT EXISTS idx_id ON users (id);

CREATE INDEX IF NOT EXISTS idx_actors_id ON actors (id);
CREATE INDEX IF NOT EXISTS idx_actors_name ON actors (name);

CREATE INDEX IF NOT EXISTS idx_movie_id ON movie (id);
CREATE INDEX IF NOT EXISTS idx_movies_title ON movie (title);
CREATE INDEX IF NOT EXISTS idx_movie_rating ON movie (rating);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS actors CASCADE;
DROP TABLE IF EXISTS movie CASCADE;
-- +goose StatementEnd
