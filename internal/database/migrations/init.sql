-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    image VARCHAR(255) DEFAULT '',
    valorant_id INT DEFAULT 0,
    user_rating INT DEFAULT 0,
    token VARCHAR(255) DEFAULT '',
    token_expire TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- CREATE TABLE valorant(
--     id SERIAL PRIMARY KEY,
--     user_id INT REFERENCES users(id) ON DELETE CASCADE,
--     riot_id INT NOT NULL,
--     username VARCHAR(255) NOT NULL,
--     rank VARCHAR(50) NOT NULL,
--     wins INT,
--     losses INT,
--     kd_ratio FLOAT,
--     headshot_percentage FLOAT,
--     matches_played INT,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );