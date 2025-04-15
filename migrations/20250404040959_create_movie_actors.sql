-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movie_actors (
    actor_id INT NOT NULL,
    movie_id INT NOT NULL,
    FOREIGN KEY (actor_id) REFERENCES actors(id),
    FOREIGN KEY (movie_id) REFERENCES movie(id),
    PRIMARY KEY (actor_id, movie_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_actors CASCADE;
-- +goose StatementEnd
