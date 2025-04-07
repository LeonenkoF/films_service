-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS actors(
	id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
	id_actor INT NOT NULL,
    id_movie INT NOT NULL,
    FOREIGN KEY (id_actor) REFERENCES actors (id),
    FOREIGN KEY (movie) REFERENCES movie (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
