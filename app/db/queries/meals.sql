-- name: GetAllMeals :many
SELECT * FROM meals;

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
WHERE calories < ?1;

-- name: GetHighProteinMeals :many
SELECT * FROM meals
WHERE protein >= 30;

-- name: GetMealsByEffort :many
SELECT * FROM meals
WHERE total_effort <= ?1;

-- name: GetFastestMeals :many
SELECT * FROM meals
ORDER BY total_time ASC;

-- name: GetMealsWithNoCutting :many
SELECT * FROM meals
WHERE cutting_effort = 0;

-- name: GetMealsWithNoPeeling :many
SELECT * FROM meals
WHERE peeling_effort = 0;

-- name: GetMealsWithMinimumWashing :many
SELECT * FROM meals
ORDER BY washing_effort ASC;

-- name: GetMealsWithMinimumIngredients :many
SELECT * FROM meals
ORDER BY LENGTH(ingredients) - LENGTH(REPLACE(ingredients, ',', '')) ASC;

-- name: InsertMeal :exec
INSERT INTO meals (
    name, category, description, light_version_instructions, instructions, image_url, calories, protein,
    cook_time, prep_time, total_time, washing_effort, peeling_effort, cutting_effort, items_required, ingredients, total_effort, servings, updated_at
) VALUES (
    ?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8, ?9, ?10, ?11, ?12, ?13, ?14, ?15, ?16, ?17, ?18, CURRENT_TIMESTAMP
);

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
	likes = ?19,
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
