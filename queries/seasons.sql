-- name: FindAllSeasons :many
SELECT *
FROM seasons
ORDER BY created DESC;

-- name: FindSeasonById :one
SELECT *
FROM seasons
WHERE id = ?;

-- name: SaveSeason :exec
INSERT OR REPLACE
INTO seasons (id, name, created)
VALUES (?, ?, ?);

-- name: DeleteSeason :exec
DELETE
FROM seasons
WHERE id = ?;
