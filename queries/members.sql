-- name: FindAllMembers :many
SELECT
  *
FROM
  members
ORDER BY
  last_name,
  first_name;

-- name: FindMemberById :one
SELECT
  *
FROM
  members
WHERE
  id = ?;

-- name: SaveMember :exec
INSERT OR REPLACE INTO
  members (id, first_name, last_name, active)
VALUES
  (?, ?, ?, ?);
