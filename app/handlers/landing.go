package handlers

import (
	"context"
	"fmt"
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"
	"lazymeal/app/types"
	"lazymeal/app/views/home"
	"log/slog"
	"strings"

	"github.com/anthdm/superkit/kit"
)

func HandleMealGuideWaitlistIndex(kit *kit.Kit) error {
	kit.Response.Header().Set("Content-Type", "text/html; charset=utf-8; image/jpeg; image/png; text/css; application/javascript")
	return kit.Render(home.MealPlanWaitlistIndex(kit.Auth().Check()))
}

func HandleMealGuideWaitlistSubmit(kit *kit.Kit) error {
	email := kit.Request.FormValue("email")
	email = kit.Request.PostFormValue("email")
	fmt.Println(email)
	if email == "" || len(email) < 6 {
		errors := map[string]string{"invalidEmail": "please enter a valid email and try again!"}
		return kit.Render(home.MealPlanWaitlistSubmitForm(errors))
	}

	slog.Info("received waitlist signup", slog.String("email", email))
	count, err := db.Get().GetWaitlistCount(kit.Request.Context())

	if err != nil {
		errors := map[string]string{"invalidEmail": "we fucked something up, sorry. please try again later!"}
		return kit.Render(home.MealPlanWaitlistSubmitForm(errors))
	}

	var discount float64
	switch {
	case count <= 200:
		discount = 0.80 // 80% discount
	case count <= 400:
		discount = 0.50 // 50% discount
	case count <= 600:
		discount = 0.30 // 30% discount
	case count <= 800:
		discount = 0.15 // 15% discount
	case count <= 1000:
		discount = 0.10 // 10% discount
	default:
		discount = 0.05 // 5% discount
	}

	waitlist, err := db.Get().InsertEmail(kit.Request.Context(), sqlc.InsertEmailParams{
		Email:    email,
		Discount: discount,
	})

	if err != nil {
		slog.Warn("failed to insert email into waitlist", slog.String("err", err.Error()), slog.String("email", email))

		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return kit.Render(home.MealPlanWaitlistSubmitSuccess(int(waitlist.ID)))
		}
	}

	return kit.Render(home.MealPlanWaitlistSubmitSuccess(int(count + 1)))
}

func HandleMealGuideIndex(kit *kit.Kit) error {
	kit.Response.Header().Set("Content-Type", "text/html; charset=utf-8; image/jpeg; image/png; text/css; application/javascript")

	return kit.Render(home.MealPlanIndex(kit.Auth().Check()))
}

func HandlePrivacyPolicyIndex(kit *kit.Kit) error {
	kit.Response.Header().Set("Content-Type", "text/html; charset=utf-8; image/jpeg; image/png; text/css; application/javascript")

	return kit.Render(home.PrivacyPolicyIndex(kit.Auth().Check()))
}

func HandleLandingIndex(kit *kit.Kit) error {
	activeFilter := getActiveFilter(kit)
	meals, err := getMealsFiltered(kit.Request.Context(), activeFilter)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}

	trendingMeals, err := db.Get().GetMealByIDs(kit.Request.Context(), types.TRENDING_MEALS)

	kit.Response.Header().Set("Content-Type", "text/html; charset=utf-8; image/jpeg; image/png; text/css; application/javascript")

	return kit.Render(home.Index(trendingMeals, meals, activeFilter, kit.Auth().Check()))
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
		meals, err = db.Get().GetMealsByCalories(ctx)
	case types.MEAL_FILTER_NO_CUTTING:
		meals, err = db.Get().GetMealsWithNoCutting(ctx)
	case types.MEAL_FILTER_NO_PEELING:
		meals, err = db.Get().GetMealsWithNoPeeling(ctx)
	case types.MEAL_FILTER_MIN_INGREDIENTS:
		meals, err = db.Get().GetMealsWithMinimumIngredients(ctx)
	case types.MEAL_FILTER_MIN_WASHING:
		meals, err = db.Get().GetMealsWithMinimumWashing(ctx)
	case types.MEAL_FILTER_ONE_SERVING:
		meals, err = db.Get().GetMealsForOnePerson(ctx)
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
