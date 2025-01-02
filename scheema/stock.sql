CREATE TABLE IF NOT EXISTS stock (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    book_id BIGINT REFERENCES books(id),
    author_id BIGINT REFERENCES authors(id)
);
