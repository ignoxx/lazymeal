package main

import (
	"context"
	"database/sql"
	"fmt"
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"

	"github.com/gosimple/slug"
)

func main() {
	ctx := context.Background()
	d := db.Get()

	meals, err := d.GetAllMeals(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("slugging START")

	slug.AppendTimestamp = true

	for _, meal := range meals {
		err := d.UpdateMeal(ctx, sqlc.UpdateMealParams{
			ID:                       meal.ID,
			Slug:                     sql.NullString{String: slug.Make(meal.Name), Valid: true},
			Category:                 meal.Category,
			Ingredients:              meal.Ingredients,
			Instructions:             meal.Instructions,
			Name:                     meal.Name,
			Description:              meal.Description,
			LightVersionInstructions: meal.LightVersionInstructions,
			ImageUrl:                 meal.ImageUrl,
			Calories:                 meal.Calories,
			Protein:                  meal.Protein,
			CookTime:                 meal.CookTime,
			PrepTime:                 meal.PrepTime,
			TotalTime:                meal.TotalTime,
			WashingEffort:            meal.WashingEffort,
			PeelingEffort:            meal.PeelingEffort,
			CuttingEffort:            meal.CuttingEffort,
			ItemsRequired:            meal.ItemsRequired,
			TotalEffort:              meal.TotalEffort,
			Servings:                 meal.Servings,
		})

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	fmt.Println("slugging DONE")
}
