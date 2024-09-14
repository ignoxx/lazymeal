package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"
	"os"
	"strings"
)

const SITE_BASE_URL = "https://lazy-meal.com/"
const SITE_BASE_URL_IMAGE_FORMAT = SITE_BASE_URL + "public/assets/%s"

type RecipeSchema struct {
	Context         string        `json:"@context"`
	Type            string        `json:"@type"`
	Name            string        `json:"name"`
	Image           string        `json:"image"`
	Author          AuthorSchema  `json:"author"`
	Description     string        `json:"description"`
	PrepTime        string        `json:"prepTime"`
	CookTime        string        `json:"cookTime"`
	TotalTime       string        `json:"totalTime"`
	RecipeYield     string        `json:"recipeYield"`
	RecipeCategory  string        `json:"recipeCategory"`
	Nutrition       NutritionInfo `json:"nutrition"`
	Ingredients     []string      `json:"recipeIngredient"`
	Instructions    []string      `json:"recipeInstructions"`
	AggregateRating RatingSchema  `json:"aggregateRating"`
}

type AuthorSchema struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type NutritionInfo struct {
	Type           string `json:"@type"`
	Calories       string `json:"calories"`
	ProteinContent string `json:"proteinContent"`
}

type RatingSchema struct {
	Type        string `json:"@type"`
	RatingValue string `json:"ratingValue"`
	RatingCount int64  `json:"ratingCount"`
}

func main() {
	ctx := context.Background()
	d := db.Get()

	meals := []sqlc.Meal{}

	mealSlugs := []string{
		"healthy-baked-salmon-with-fennel-and-roasted-cherry-tomatoes-1726349012",
		"morning-whey-protein-coffee-shake-1726349012",
		"overnight-oats-with-yogurt-and-fruits-1726349012",
		"berry-kefir-smoothie-1726349012",
		"traditional-beetroot-soup-borscht-healthy-and-nutritious-1726349012",
		"hearty-croatian-bean-stew-with-cannellini-beans-and-chorizo-1726349012",
		"moroccan-carrot-soup-healthy-and-flavorful-1726349012",
		"spicy-arrabiata-penne-1726349012",
	}

	for _, mSlug := range mealSlugs {
		meal, err := d.GetMealBySlug(ctx, sql.NullString{String: mSlug, Valid: true})

		if err != nil {
			panic(err)
		}

		meals = append(meals, meal)
	}

	structuredData := generateRecipesSchema(meals)

	// write to json file
	jsonData, err := json.Marshal(structuredData)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("recipes.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Recipes structured data generated successfully")
}

func generateRecipesSchema(meals []sqlc.Meal) []RecipeSchema {
	var recipes []RecipeSchema

	for _, meal := range meals {
		// Convert times to ISO 8601 duration format (PT#M for minutes)
		prepTime := fmt.Sprintf("PT%dM", meal.PrepTime)
		cookTime := fmt.Sprintf("PT%dM", meal.CookTime)
		totalTime := fmt.Sprintf("PT%dM", meal.TotalTime)

		if meal.Likes == 0 {
			meal.Likes++
		}

		recipe := RecipeSchema{
			Context:        "https://schema.org",
			Type:           "Recipe",
			Name:           meal.Name,
			Image:          fmt.Sprintf(SITE_BASE_URL_IMAGE_FORMAT, meal.ImageUrl),
			Author:         AuthorSchema{Type: "Organization", Name: "Lazy Meal"},
			Description:    meal.Description,
			PrepTime:       prepTime,
			CookTime:       cookTime,
			TotalTime:      totalTime,
			RecipeYield:    fmt.Sprintf("%d servings", meal.Servings),
			RecipeCategory: meal.Category,
			Nutrition: NutritionInfo{
				Type:           "NutritionInformation",
				Calories:       fmt.Sprintf("%d calories", meal.Calories),
				ProteinContent: fmt.Sprintf("%d grams", meal.Protein),
			},
			Ingredients:  strings.Split(meal.Ingredients, "\n"),  // Assuming ingredients are newline-separated
			Instructions: strings.Split(meal.Instructions, "\n"), // Assuming instructions are newline-separated
			AggregateRating: RatingSchema{
				Type:        "AggregateRating",
				RatingValue: "4.5", // Placeholder, replace with real data if available
				RatingCount: meal.Likes,
			},
		}

		recipes = append(recipes, recipe)
	}

	return recipes
}
