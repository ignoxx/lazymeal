-- name: GetAllMeals :many
SELECT * FROM meals;

-- name: CreateMeal :one
INSERT INTO meals (
    name, category, description, short_description,
    image_url, image_preview, calories, likes, total_effort,
    instruction_steps, cook_time, prep_time, total_time, created_at, updated_at
) VALUES (
    ?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8, ?9, ?10,
    ?11, ?12, ?13, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING *;

-- name: GetAllMealsPaginated :many
-- SELECT * FROM meals
-- WHERE created_at > ?1
-- ORDER BY created_at DESC
-- LIMIT ?2;


-- name: GetAllMealsPaginated :many
SELECT * FROM meals
ORDER BY created_at DESC
LIMIT ?1 OFFSET ?2;

