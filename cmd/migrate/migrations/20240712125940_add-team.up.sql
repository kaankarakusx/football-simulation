CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    strength INT DEFAULT 0,
    points INT DEFAULT 0,
    matches INT DEFAULT 0,
    wins INT DEFAULT 0,
    draws INT DEFAULT 0,
    losses INT DEFAULT 0,
    goals_for INT DEFAULT 0,
    goals_against INT DEFAULT 0,
    goals_difference INT DEFAULT 0,
    temporary_drop INT DEFAULT 0  
);