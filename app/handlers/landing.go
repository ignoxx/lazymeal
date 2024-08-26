package handlers

import (
	"lazymeal/app/db"
	"lazymeal/app/types"
	"lazymeal/app/views/home"

	"github.com/anthdm/superkit/kit"
)

func HandleLandingIndex(kit *kit.Kit) error {
	activeFilters := getActiveFilters(kit)

	meals, err := db.Get().GetAllMeals(kit.Request.Context())
	if err != nil {
		return err
	}

	// TODO: apply filters
	// and fetch 200 meals

	return kit.Render(home.Index(meals[:3], meals, activeFilters))
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
