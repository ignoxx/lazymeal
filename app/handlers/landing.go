package handlers

import (
	"lazymeal/app/types"
	"lazymeal/app/views/home"

	"github.com/anthdm/superkit/kit"
)

func HandleLandingIndex(kit *kit.Kit) error {
	trendingMeals := []types.Meal{}
	meals := []types.Meal{
		{Name: "Spaghetti Carbonara", Description: "A delicious pasta dish with bacon and eggs.", ShortDescription: "A delicious pasta dish with bacon and eggs. A delicious pasta dish with bacon and eggs.A delicious pasta dish with bacon and eggs. ", Ingredients: []types.Ingredient{{Name: "Spaghetti", Unit: types.GRAMS, Amount: 200, WashingEffort: 1, PeelingEffort: 0}, {Name: "Bacon", Unit: types.GRAMS, Amount: 100, WashingEffort: 1, PeelingEffort: 0}, {Name: "Eggs", Unit: types.PIECE, Amount: 2, WashingEffort: 1, PeelingEffort: 0}}, Items: []types.CookItem{{Name: "Pan", WashingEffort: 1}, {Name: "Pot", WashingEffort: 1}}, InstructionSteps: []string{"Cook the spaghetti", "Fry the bacon", "Mix the eggs", "Mix everything together"}, Nutritions: []types.Nutrition{{Name: "Calories", Unit: types.GRAMS, Amount: 500}}, TotalEffort: 5, Calories: 500, Likes: 420, Comments: []string{}, ImageUrl: "https://i.pinimg.com/564x/66/fd/76/66fd766d2f75f4fb791d248087864ef1.jpg"},
	}

	return kit.Render(home.Index(trendingMeals, meals))
}
