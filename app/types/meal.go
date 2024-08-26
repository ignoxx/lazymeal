package types

import "time"

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

var (
	MEAL_FILTERS []MealFilter = []MealFilter{
		{ID: "fastest", Name: "Fastest", Emoji: "🚀"},
		{ID: "high-protein", Name: "High Protein", Emoji: "💪"},
		{ID: "low-calorie", Name: "Low-Calorie", Emoji: "🥗"},
		{ID: "no-cutting", Name: "No Cutting", Emoji: "🔪"},
		{ID: "no-peeling", Name: "No Peeling", Emoji: "🥕"},
		{ID: "min-ingredients", Name: "Minimal Ingredients", Emoji: "🥄"},
		{ID: "min-washing", Name: "Minimal Washing", Emoji: "🧼"},
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
