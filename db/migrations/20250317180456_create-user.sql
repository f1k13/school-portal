-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
		auth_code TEXT UNIQUE NOT NULL DEFAULT '',
		verified BOOLEAN DEFAULT FALSE NOT NULL,
		refresh_token TEXT UNIQUE NOT NULL DEFAULT ''
);

CREATE TABLE profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    first_name TEXT,
    last_name TEXT,
    phone_number TEXT,
    avatar_url TEXT,
    dob TEXT,
    created_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE profiles;
DROP TABLE users;
