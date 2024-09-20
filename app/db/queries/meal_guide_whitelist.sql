-- name: InsertEmail :one
INSERT INTO meal_guide_whitelists (email,discount) VALUES (?1, ?2)
RETURNING *;

-- name: GetWaitlistCount :one
SELECT COUNT(*) FROM meal_guide_whitelists;
