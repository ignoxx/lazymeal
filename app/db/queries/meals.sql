-- name: GetAllMeals :many
SELECT * FROM meals
ORDER BY category DESC;

-- name: GetAllMealsPaginated :many
SELECT * FROM meals
ORDER BY created_at DESC
LIMIT ?1 OFFSET ?2;

-- name: GetMealByID :one
SELECT * FROM meals
WHERE id = ?1;

-- name: GetMealByIDs :many
SELECT * FROM meals
WHERE id IN (?1, ?2, ?3);

-- name: GetMealsByCategory :many
SELECT * FROM meals
WHERE category = ?1;

-- name: GetMealsByCalories :many
SELECT * FROM meals
ORDER BY calories ASC;

-- name: GetMealsForOnePerson :many
SELECT * FROM meals
WHERE servings = 1 AND category LIKE '%dinner%'
ORDER BY total_time ASC;

-- name: GetHighProteinMeals :many
SELECT * FROM meals
ORDER BY protein DESC;

-- name: GetMealsByEffort :many
SELECT * FROM meals
ORDER BY total_effort ASC;

-- name: GetFastestMeals :many
SELECT * FROM meals
ORDER BY total_time ASC;

-- name: GetNewestMeals :many
SELECT * FROM meals
ORDER BY created_at DESC
LIMIT ?1;

-- name: GetMostLikedMeals :many
SELECT * FROM meals
ORDER BY likes DESC
LIMIT ?1;

-- name: GetMealsWithNoCutting :many
SELECT * FROM meals
ORDER BY cutting_effort ASC;

-- name: GetMealsWithNoPeeling :many
SELECT * FROM meals
ORDER BY peeling_effort ASC;

-- name: GetMealsWithMinimumWashing :many
SELECT * FROM meals
ORDER BY washing_effort ASC;

-- name: GetMealsWithMinimumIngredients :many
SELECT * FROM meals
ORDER BY LENGTH(ingredients) ASC;

-- name: InsertMeal :one
INSERT INTO meals (
    name, category, description, light_version_instructions, instructions, image_url, calories, protein,
    cook_time, prep_time, total_time, washing_effort, peeling_effort, cutting_effort, items_required, ingredients, total_effort, servings, updated_at
) VALUES (
    ?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8, ?9, ?10, ?11, ?12, ?13, ?14, ?15, ?16, ?17, ?18, CURRENT_TIMESTAMP
)
RETURNING *;

-- name: UpdateMeal :exec
UPDATE meals
SET
    name = ?2,
    category = ?3,
    description = ?4,
    light_version_instructions = ?5,
    instructions = ?6,
    image_url = ?7,
    calories = ?8,
    protein = ?9,
    cook_time = ?10,
    prep_time = ?11,
    total_time = ?12,
    washing_effort = ?13,
    peeling_effort = ?14,
    cutting_effort = ?15,
    items_required = ?16,
    ingredients = ?17,
    total_effort = ?18,
    servings = ?19,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?1;

-- name: UpdateMealLikes :exec
UPDATE meals
SET
    likes = likes + 1,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?1;

-- name: DeleteMeal :exec
DELETE FROM meals
WHERE id = ?1;
