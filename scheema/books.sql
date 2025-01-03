CREATE TABLE IF NOT EXISTS books (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50),
    catagory VARCHAR(50) NOT NULL,
    bar_code_id BIGINT REFERENCES bar_codes(id),
    author_id BIGINT REFERENCES authors(id)
);
