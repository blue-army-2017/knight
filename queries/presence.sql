-- name: FindSeasonPresence :many
SELECT
  s.name,
  SUM(g.home) AS home_games,
  COUNT(g.id) AS total_games
FROM
  seasons s
  INNER JOIN season_games g ON g.season_id = s.id
GROUP BY
  s.id
ORDER BY
  s.created DESC;

-- name: FindMemberPresence :many
SELECT
  s.name AS season,
  m.last_name,
  m.first_name,
  SUM(g.home) AS home_games,
  COUNT(g.id) AS total_games
FROM
  present_members p
  INNER JOIN members m ON m.id = p.member_id
  INNER JOIN season_games g ON g.id = p.season_game_id
  INNER JOIN seasons s ON s.id = g.season_id
WHERE
  m.active = TRUE
GROUP BY
  s.id,
  m.id
ORDER BY
  s.created DESC,
  total_games DESC,
  m.last_name,
  m.first_name;

-- name: FindPresentMembersForGame :many
SELECT
  m.id
FROM
  present_members p
  INNER JOIN members m ON m.id = p.member_id
  INNER JOIN season_games g ON g.id = p.season_game_id
WHERE
  g.id = ?;

-- name: DeletePresentMembersForGame :exec
DELETE FROM present_members
WHERE
  season_game_id = ?;

-- name: SavePresentMemberForGame :exec
INSERT INTO
  present_members (season_game_id, member_id)
VALUES
  (?, ?);
