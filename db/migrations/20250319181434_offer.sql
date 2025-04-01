-- +goose Up
-- Таблица направлений
create table direction (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT NOT NULL,
	image TEXT NOT NULL,
	description TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT now()
);
-- Таблица объявлений препода
CREATE table offers (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	price INT NOT NULL,
	direction_id UUID NOT NULL REFERENCES direction(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    is_online BOOLEAN NOT NULL,
	created_at TIMESTAMP DEFAULT now()
);

-- Таблица образования
CREATE TABLE educations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    institution TEXT NOT NULL, 
    degree TEXT NOT NULL,     
    start_year INT NOT NULL,
    end_year INT NOT NULL,        
	city TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
-- Таблица опыта
CREATE TABLE experiences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    company TEXT NOT NULL,  
    role TEXT NOT NULL,     
    years INT NOT NULL,    
    created_at TIMESTAMP DEFAULT now()
);

-- Таблица связей объявлений с образованием
CREATE TABLE offer_educations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_id UUID NOT NULL REFERENCES offers(id) ON DELETE CASCADE,
    education_id UUID NOT NULL REFERENCES educations(id) ON DELETE CASCADE,
    UNIQUE (offer_id, education_id)
    
);

-- Таблица умений

CREATE TABLE skills (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    image TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
-- Таблица связей объявлений с опытом работы
CREATE TABLE offer_experiences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_id UUID NOT NULL REFERENCES offers(id) ON DELETE CASCADE,
    experience_id UUID NOT NULL REFERENCES experiences(id) ON DELETE CASCADE,
    UNIQUE (offer_id, experience_id)
);

-- Таблица связей объявлений с умениями
CREATE TABLE offer_skills (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_id UUID NOT NULL REFERENCES offers(id) ON DELETE CASCADE,
    skill_id UUID NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    UNIQUE (offer_id, skill_id)
);



-- +goose Down
DROP TABLE offer_experiences;
DROP TABLE offer_educations;
DROP TABLE offer_skills;

DROP TABLE offers;
DROP TABLE experiences;
DROP TABLE educations;
DROP TABLE direction;
DROP TABLE skills;