CREATE TABLE IF NOT EXISTS authors_rating (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    COMMENT VARCHAR(50),
    rating INT NOT NULL,
    user_id BIGINT REFERENCES student(id)
);
