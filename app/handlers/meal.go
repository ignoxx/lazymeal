package handlers

import (
	"database/sql"
	"fmt"
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"
	"lazymeal/app/views/errors"
	mealView "lazymeal/app/views/meal"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/anthdm/superkit/kit"
)

func HandleMealIndex(kit *kit.Kit) error {
	mealIDStr := kit.Request.PathValue("mealID")
	mealID, err := strconv.ParseInt(mealIDStr, 10, 64)
	if err != nil {
		return kit.Render(errors.Error404())
	}

	trendingMeals, err := db.Get().GetMostLikedMeals(kit.Request.Context(), 3)

	meal, err := db.Get().GetMealByID(kit.Request.Context(), mealID)

	if err != nil {
		slog.Warn("failed to get meal by id", slog.String("err", err.Error()))
		kit.Redirect(http.StatusSeeOther, "/")
	}

	return kit.Render(mealView.Index(trendingMeals, meal, kit.Auth().Check()))
}

func HandleMealLike(kit *kit.Kit) error {
	mealIDStr := kit.Request.PathValue("mealID")
	mealID, err := strconv.ParseInt(mealIDStr, 10, 64)
	if err != nil {
		return kit.Render(errors.Error404())
	}

	err = db.Get().UpdateMealLikes(kit.Request.Context(), mealID)
	if err != nil {
		slog.Warn("failed to update meal likes", slog.String("err", err.Error()))
	}

	// 200 OK
	return kit.Render(mealView.LikeButtonClicked())
}
func HandleMealEdit(kit *kit.Kit) error {
	mealIDStr := kit.Request.PathValue("mealID")
	mealID, err := strconv.ParseInt(mealIDStr, 10, 64)
	if err != nil {
		return kit.Render(errors.Error404())
	}

	lightVersionInstructionsValue := strings.TrimSpace(kit.Request.FormValue("light_version_instructions"))
	lightVersionInstructions := sql.NullString{
		String: lightVersionInstructionsValue,
		Valid:  len(lightVersionInstructionsValue) > 0,
	}

	imageAltValue := strings.TrimSpace(kit.Request.FormValue("image_alt"))
	imageAlt := sql.NullString{
		String: imageAltValue,
		Valid:  len(imageAltValue) > 0,
	}

	var (
		servings, _       = strconv.Atoi(kit.Request.FormValue("servings"))
		washingEffort, _  = strconv.Atoi(kit.Request.FormValue("washing_effort"))
		peeelingEffort, _ = strconv.Atoi(kit.Request.FormValue("peeling_effort"))
		cuttingEffort, _  = strconv.Atoi(kit.Request.FormValue("cutting_effort"))
		calories, _       = strconv.Atoi(kit.Request.FormValue("calories"))
		protein, _        = strconv.Atoi(kit.Request.FormValue("protein"))
		cookTime, _       = strconv.Atoi(kit.Request.FormValue("cook_time"))
		prepTime, _       = strconv.Atoi(kit.Request.FormValue("prep_time"))
		totalTime         = cookTime + prepTime
		totalEffort       = math.Floor(float64(washingEffort)+float64(peeelingEffort)+float64(cuttingEffort)) / 3
	)

	updateParams := sqlc.UpdateMealParams{
		ID:                       mealID,
		Name:                     kit.Request.FormValue("name"),
		Description:              kit.Request.FormValue("description"),
		Category:                 kit.Request.FormValue("category"),
		LightVersionInstructions: lightVersionInstructions,
		Instructions:             kit.Request.FormValue("instructions"),
		ImageUrl:                 kit.Request.FormValue("image_url"),
		ImageAlt:                 imageAlt,
		Calories:                 int64(calories),
		Protein:                  int64(protein),
		CookTime:                 int64(cookTime),
		PrepTime:                 int64(prepTime),
		TotalTime:                int64(totalTime),
		WashingEffort:            int64(washingEffort),
		PeelingEffort:            int64(peeelingEffort),
		CuttingEffort:            int64(cuttingEffort),
		ItemsRequired:            kit.Request.FormValue("items_required"),
		Ingredients:              kit.Request.FormValue("ingredients"),
		TotalEffort:              int64(totalEffort),
		Servings:                 int64(servings),
	}

	err = db.Get().UpdateMeal(kit.Request.Context(), updateParams)
	if err != nil {
		return kit.Render(errors.Error500())
	}

	// 200 OK
	return kit.Redirect(http.StatusSeeOther, fmt.Sprintf("/%d", mealID))
}

func HandleMealEditIndex(kit *kit.Kit) error {
	mealIDStr := kit.Request.PathValue("mealID")
	mealID, err := strconv.ParseInt(mealIDStr, 10, 64)
	if err != nil {
		return kit.Render(errors.Error404())
	}
	meal, err := db.Get().GetMealByID(kit.Request.Context(), mealID)

	// 200 OK
	return kit.Render(mealView.Edit(meal, kit.Auth().Check()))
}

func HandleMealCreateIndex(kit *kit.Kit) error {
	// 200 OK
	return kit.Render(mealView.Create(kit.Auth().Check()))
}

func HandleMealCreate(kit *kit.Kit) error {
	lightVersionInstructions := sql.NullString{}

	if kit.Request.FormValue("light_version_instructions") != "" {
		lightVersionInstructions = sql.NullString{
			String: kit.Request.FormValue("light_version_instructions"),
			Valid:  true,
		}
	}

	var (
		servings, _       = strconv.Atoi(kit.Request.FormValue("servings"))
		washingEffort, _  = strconv.Atoi(kit.Request.FormValue("washing_effort"))
		peeelingEffort, _ = strconv.Atoi(kit.Request.FormValue("peeling_effort"))
		cuttingEffort, _  = strconv.Atoi(kit.Request.FormValue("cutting_effort"))
		calories, _       = strconv.Atoi(kit.Request.FormValue("calories"))
		protein, _        = strconv.Atoi(kit.Request.FormValue("protein"))
		cookTime, _       = strconv.Atoi(kit.Request.FormValue("cook_time"))
		prepTime, _       = strconv.Atoi(kit.Request.FormValue("prep_time"))
		totalTime         = cookTime + prepTime
		totalEffort       = math.Floor(float64(washingEffort)+float64(peeelingEffort)+float64(cuttingEffort)) / 3
	)

	updateParams := sqlc.InsertMealParams{
		Name:                     kit.Request.FormValue("name"),
		Description:              kit.Request.FormValue("description"),
		Category:                 kit.Request.FormValue("category"),
		LightVersionInstructions: lightVersionInstructions,
		Instructions:             kit.Request.FormValue("instructions"),
		ImageUrl:                 kit.Request.FormValue("image_url"),
		Calories:                 int64(calories),
		Protein:                  int64(protein),
		CookTime:                 int64(cookTime),
		PrepTime:                 int64(prepTime),
		TotalTime:                int64(totalTime),
		WashingEffort:            int64(washingEffort),
		PeelingEffort:            int64(peeelingEffort),
		CuttingEffort:            int64(cuttingEffort),
		ItemsRequired:            kit.Request.FormValue("items_required"),
		Ingredients:              kit.Request.FormValue("ingredients"),
		TotalEffort:              int64(totalEffort),
		Servings:                 int64(servings),
	}

	meal, err := db.Get().InsertMeal(kit.Request.Context(), updateParams)
	if err != nil {
		return kit.Render(errors.Error500())
	}

	return kit.Redirect(http.StatusSeeOther, fmt.Sprintf("/%d", meal.ID))
}
