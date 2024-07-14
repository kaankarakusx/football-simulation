CREATE TABLE league (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    current_week INT DEFAULT 0,
    total_weeks INT DEFAULT 0,
    champion_team_name VARCHAR(255)
);