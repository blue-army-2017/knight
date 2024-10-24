CREATE TABLE present_members (
  season_game_id text NOT NULL,
  member_id text NOT NULL,
  PRIMARY KEY (season_game_id,member_id),
  CONSTRAINT fk_present_members_member FOREIGN KEY (member_id) REFERENCES members(id)
  ,CONSTRAINT fk_present_members_season_game FOREIGN KEY (season_game_id) REFERENCES season_games(id)
);
