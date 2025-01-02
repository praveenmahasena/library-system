CREATE TABLE IF NOT EXISTS book_rating(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    comment VARCHAR(50),
    rating INT NOT NULL,
    user_id BIGINT REFERENCES student(id)
);
