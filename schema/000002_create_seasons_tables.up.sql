CREATE TABLE seasons (
  id text NOT NULL,
  name text NOT NULL,
  created text NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE season_games (
  id text NOT NULL,
  opponent text NOT NULL,
  home numeric NOT NULL,
  mode text NOT NULL,
  date text NOT NULL,
  season_id text NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_seasons_games FOREIGN KEY (season_id) REFERENCES seasons(id),
  CONSTRAINT fk_season_games_season FOREIGN KEY (season_id) REFERENCES seasons(id)
);
