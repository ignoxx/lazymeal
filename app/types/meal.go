package types

import "time"

type Unit string

const (
	GRAMS Unit = "g"
	ML    Unit = "ml"
	PIECE Unit = "pcs"
)

type Meal struct {
	Name             string
	Description      string
	ShortDescription string
	ImageUrl         string
	ImagePreview     string
	Ingredients      []Ingredient
	Items            []CookItem
	InstructionSteps []string
	Nutritions       []Nutrition
	TotalEffort      int
	Calories         int
	CookTime         time.Time
	PrepTime         time.Time
	TotalTime        time.Time
	Likes            int
	Comments         []string
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
