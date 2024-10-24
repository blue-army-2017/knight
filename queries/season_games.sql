-- name: FindAllSeasonGames :many
SELECT sqlc.embed(g), s.name AS 'season_name'
FROM season_games g
INNER JOIN seasons s ON s.id = g.season_id
ORDER BY date DESC;

-- name: FindAllSeasonGamesBySeason :many
SELECT sqlc.embed(g), s.name AS 'season_name'
FROM season_games g
INNER JOIN seasons s ON s.id = g.season_id
WHERE g.season_id = ?
ORDER BY g.date DESC;

-- name: FindSeasonGameById :one
SELECT sqlc.embed(g), s.name AS 'season_name'
FROM season_games g
INNER JOIN seasons s ON s.id = g.season_id
WHERE g.id = ?;

-- name: SaveSeasonGame :exec
INSERT OR REPLACE
INTO season_games (id, opponent, home, mode, date, season_id)
VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteSeasonGame :exec
DELETE
FROM season_games
WHERE id = ?;
