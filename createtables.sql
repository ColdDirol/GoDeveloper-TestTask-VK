-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(100),
    role VARCHAR(100)
);

-- Таблица фильмов
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150),
    rating FLOAT,
    date DATE,
    user_id INT REFERENCES users(id)
);

-- Таблица актеров
CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    sex VARCHAR(10),
    date_of_birth DATE
);

-- Таблица связи актеров и фильмов
CREATE TABLE actor_movies (
    actor_id INT REFERENCES actors(id),
    movie_id INT REFERENCES movies(id)
);
