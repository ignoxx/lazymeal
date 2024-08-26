package handlers

import (
	"lazymeal/app/db"
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

	meal, err := db.Get().GetMealByID(kit.Request.Context(), mealID)

	return kit.Render(mealView.Index(meal))
}
