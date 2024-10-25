-- name: FindAllSeasonGames :many
SELECT
  /* sql-formatter-disable */
  sqlc.embed(g),
  /* sql-formatter-enable */
  s.name AS 'season_name'
FROM
  season_games g
  INNER JOIN seasons s ON s.id = g.season_id
ORDER BY
  DATE DESC;

-- name: FindAllSeasonGamesBySeason :many
SELECT
  /* sql-formatter-disable */
  sqlc.embed(g),
  /* sql-formatter-enable */
  s.name AS 'season_name'
FROM
  season_games g
  INNER JOIN seasons s ON s.id = g.season_id
WHERE
  g.season_id = ?
ORDER BY
  g.date DESC;

-- name: FindSeasonGameById :one
SELECT
  /* sql-formatter-disable */
  sqlc.embed(g),
  /* sql-formatter-enable */
  s.name AS 'season_name'
FROM
  season_games g
  INNER JOIN seasons s ON s.id = g.season_id
WHERE
  g.id = ?;

-- name: SaveSeasonGame :exec
INSERT OR REPLACE INTO
  season_games (id, opponent, home, mode, DATE, season_id)
VALUES
  (?, ?, ?, ?, ?, ?);

-- name: DeleteSeasonGame :exec
DELETE FROM season_games
WHERE
  id = ?;
