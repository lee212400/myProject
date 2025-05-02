-- name: GetSample :one
SELECT * FROM samples
WHERE id = $1 LIMIT 1;

-- name: ListSamples :many
SELECT * FROM samples
ORDER BY sample_title;

-- name: CreateSample :one
INSERT INTO samples (
  sample_title, sample_memo
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateSample :exec
UPDATE samples
  set sample_title = $2,
  sample_memo = $3
WHERE id = $1;

-- name: DeleteSample :exec
DELETE FROM samples
WHERE id = $1;