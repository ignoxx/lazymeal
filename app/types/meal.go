package types

import (
	"lazymeal/app/db/sqlc"
	"time"
)

type Unit string

const (
	GRAMS Unit = "g"
	ML    Unit = "ml"
	PIECE Unit = "pcs"
)

type MealFilter struct {
	ID    string
	Name  string
	Emoji string
}

const (
	MEAL_FILTER_FASTEST         = "fastest"
	MEAL_FILTER_LOWEST_EFFORT   = "lowest-effort"
	MEAL_FILTER_HIGH_PROTEIN    = "high-protein"
	MEAL_FILTER_LOW_CALORIE     = "low-calorie"
	MEAL_FILTER_NO_CUTTING      = "no-cutting"
	MEAL_FILTER_NO_PEELING      = "no-peeling"
	MEAL_FILTER_MIN_INGREDIENTS = "min-ingredients"
	MEAL_FILTER_MIN_WASHING     = "min-washing"
	MEAL_FILTER_ONE_SERVING     = "one-serving"
)

var (
	MEAL_FILTERS []MealFilter = []MealFilter{
		{ID: MEAL_FILTER_FASTEST, Name: "Fastest", Emoji: "🚀"},
		{ID: MEAL_FILTER_LOWEST_EFFORT, Name: "Lowest Effort", Emoji: "👌"},
		{ID: MEAL_FILTER_HIGH_PROTEIN, Name: "High Protein", Emoji: "💪"},
		{ID: MEAL_FILTER_LOW_CALORIE, Name: "Low-Calorie", Emoji: "🥗"},
		{ID: MEAL_FILTER_ONE_SERVING, Name: "Lazy dinner for one", Emoji: "🍽"},
		// {ID: MEAL_FILTER_NO_CUTTING, Name: "No Cutting", Emoji: "🔪"},
		// {ID: MEAL_FILTER_NO_PEELING, Name: "No Peeling", Emoji: "🥕"},
		{ID: MEAL_FILTER_MIN_INGREDIENTS, Name: "Minimal Ingredients", Emoji: "🥄"},
		{ID: MEAL_FILTER_MIN_WASHING, Name: "Minimal Washing", Emoji: "🧼"},
	}

	TRENDING_MEALS sqlc.GetMealByIDsParams = sqlc.GetMealByIDsParams{
		ID:   15,
		ID_2: 12,
		ID_3: 17,
	}
)

type Meal struct {
	Name             string
	Description      string
	ShortDescription string
	ImageUrl         string
	ImagePreview     string
	TotalEffort      int
	Calories         int
	CookTime         time.Time
	PrepTime         time.Time
	TotalTime        time.Time
	InstructionSteps []string

	Ingredients []Ingredient
	Nutritions  []Nutrition
	Items       []CookItem
	Comments    []string
	Likes       int
}

type Like struct {
	Source    string
	CreatedAt time.Time
}

type Ingredient struct {
	Name          string
	Unit          Unit
	Amount        float32
	WashingEffort int
	PeelingEffort int
}

type CookItem struct {
	Name          string
	WashingEffort int
}

type Nutrition struct {
	Name   string
	Unit   Unit
	Amount float32
}

type Comment struct {
	Author    string
	Content   string
	Approved  bool
	CreatedAt time.Time
}
