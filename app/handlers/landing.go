package handlers

import (
	"context"
	"fmt"
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"
	"lazymeal/app/types"
	"lazymeal/app/views/home"

	"github.com/anthdm/superkit/kit"
)

func HandleLandingIndex(kit *kit.Kit) error {
	activeFilter := getActiveFilter(kit)
	meals, err := getMealsFiltered(kit.Request.Context(), activeFilter)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}

	trendingMeals, err := db.Get().GetMealByIDs(kit.Request.Context(), sqlc.GetMealByIDsParams{
		ID:   7,
		ID_2: 5,
		ID_3: 3,
	})

	return kit.Render(home.Index(trendingMeals, meals, activeFilter))
}

func getActiveFilter(kit *kit.Kit) string {
	activeFilter := ""
	filters, ok := kit.Request.URL.Query()["filter"]
	if ok && len(filters) > 0 {
		activeFilter = filters[0]
	}

	return activeFilter
}

func getMealsFiltered(ctx context.Context, filter string) ([]sqlc.Meal, error) {
	var meals []sqlc.Meal
	var err error
	if filter == "" {
		meals, err = db.Get().GetAllMeals(ctx)
	}

	switch filter {
	case types.MEAL_FILTER_FASTEST:
		meals, err = db.Get().GetFastestMeals(ctx)
	case types.MEAL_FILTER_HIGH_PROTEIN:
		meals, err = db.Get().GetHighProteinMeals(ctx)
	case types.MEAL_FILTER_LOW_CALORIE:
		meals, err = db.Get().GetMealsByCalories(ctx, 400)
	case types.MEAL_FILTER_NO_CUTTING:
		meals, err = db.Get().GetMealsWithNoCutting(ctx)
	case types.MEAL_FILTER_NO_PEELING:
		meals, err = db.Get().GetMealsWithNoPeeling(ctx)
	case types.MEAL_FILTER_MIN_INGREDIENTS:
		meals, err = db.Get().GetMealsWithMinimumIngredients(ctx)
	case types.MEAL_FILTER_MIN_WASHING:
		meals, err = db.Get().GetMealsWithMinimumWashing(ctx)
	default:
		meals, err = db.Get().GetAllMeals(ctx)
	}

	if err != nil {
		return nil, err
	}

	return meals, nil
}

func getActiveFilters(kit *kit.Kit) map[string]bool {
	queryValues := kit.Request.URL.Query()

	if _, ok := queryValues["filter"]; !ok {
		return map[string]bool{}
	}

	var filters = map[string]bool{}

	for _, filter := range queryValues["filter"] {
		for _, f := range types.MEAL_FILTERS {
			if f.ID == filter {
				filters[filter] = true
			}
		}
	}

	return filters
}
