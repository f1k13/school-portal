-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
		middle_name TEXT,
		refresh_token TEXT NOT NULL,
		role TEXT NOT NULL,
		phone_number TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE users;
jet -dsn=postgres://postgres:1234@localhost:5432/school-portal?sslmode=disable -schema=dvds -path=./.internal/storage 