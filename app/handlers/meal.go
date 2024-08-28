package handlers

import (
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"
	"lazymeal/app/views/errors"
	mealView "lazymeal/app/views/meal"
	"strconv"

	"github.com/anthdm/superkit/kit"
)

func HandleMealIndex(kit *kit.Kit) error {
	mealIDStr := kit.Request.PathValue("mealID")
	mealID, err := strconv.ParseInt(mealIDStr, 10, 64)
	if err != nil {
		return kit.Render(errors.Error404())
	}

	trendingMeals, err := db.Get().GetMealByIDs(kit.Request.Context(), sqlc.GetMealByIDsParams{
		ID:   7,
		ID_2: 5,
		ID_3: 3,
	})

	meal, err := db.Get().GetMealByID(kit.Request.Context(), mealID)

	return kit.Render(mealView.Index(trendingMeals, meal))
}

func HandleMealLike(kit *kit.Kit) error {
	mealIDStr := kit.Request.PathValue("mealID")
	mealID, err := strconv.ParseInt(mealIDStr, 10, 64)
	if err != nil {
		return kit.Render(errors.Error404())
	}

	err = db.Get().UpdateMealLikes(kit.Request.Context(), mealID)
	if err != nil {
		return kit.Render(errors.Error500())
	}

	// 200 OK
	return kit.Render(mealView.LikeButtonClicked())
}
