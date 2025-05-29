--create_table.sql

DROP TABLE IF EXISTS teams;
CREATE TABLE IF NOT EXISTS teams (
    name TEXT PRIMARY KEY,
    defence INTEGER,
    midfield INTEGER,
    forward INTEGER,
    hometeammodifier INTEGER,
    played INTEGER,
    wins INTEGER,
    draws INTEGER,
    losses INTEGER,
    goals_for INTEGER,
    goals_against INTEGER,
    points INTEGER
);

DROP TABLE  IF EXISTS matches;
CREATE TABLE IF NOT EXISTS matches (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    home_team TEXT,
    away_team TEXT,
    home_goals INTEGER,
    away_goals INTEGER,
    played BOOLEAN,
    FOREIGN KEY(home_team) REFERENCES teams(name),
    FOREIGN KEY(away_team) REFERENCES teams(name)
);

DROP TABLE  IF EXISTS scoreboard;
CREATE TABLE IF NOT EXISTS scoreboard (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    team_name TEXT,
    matches_played INTEGER,
    wins INTEGER,
    draws INTEGER,
    losses INTEGER,
    goals_scored INTEGER,
    goals_against INTEGER,
    goals_difference INTEGER,
    points INTEGER,
    FOREIGN KEY(team_name) REFERENCES teams(name)
);

