package main

import (
	"context"
	"database/sql"
	"fmt"
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"
)

func main() {
	ctx := context.Background()
	d := db.Get()

	mealsParams := []sqlc.InsertMealParams{
		// {
		// 	Name:        "Vegetable Stir-Fry",
		// 	Category:    "Dinner",
		// 	Description: "A quick and healthy vegetable stir-fry with a savory sauce.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Skip the sauce and simply season with salt and pepper.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Heat oil in a pan. \n2. Add sliced vegetables and stir-fry until tender. \n3. Add soy sauce, ginger, and garlic, and stir-fry for another 2 minutes.",
		// 	ImageUrl:      "https://example.com/vegetable_stir_fry.jpg",
		// 	Calories:      250,
		// 	Protein:       8,
		// 	CookTime:      10,
		// 	PrepTime:      5,
		// 	WashingEffort: 3,
		// 	PeelingEffort: 2,
		// 	CuttingEffort: 4,
		// 	ItemsRequired: "Pan, Knife, Cutting Board",
		// 	Ingredients:   "1 cup broccoli, 1 bell pepper, 1 carrot, 2 tbsp soy sauce, 1 tsp ginger, 1 garlic clove, 1 tbsp oil",
		// },
		// {
		// 	Name:        "Chicken Wrap",
		// 	Category:    "Lunch",
		// 	Description: "A healthy chicken wrap with vegetables and a tangy sauce.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Skip the sauce and use pre-cooked chicken for faster preparation.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Grill the chicken until fully cooked. \n2. Slice the vegetables. \n3. Place the chicken and vegetables into a wrap, and add sauce. \n4. Roll up and serve.",
		// 	ImageUrl:      "https://example.com/chicken_wrap.jpg",
		// 	Calories:      400,
		// 	Protein:       30,
		// 	CookTime:      10,
		// 	PrepTime:      5,
		// 	WashingEffort: 2,
		// 	PeelingEffort: 0,
		// 	CuttingEffort: 2,
		// 	ItemsRequired: "Grill, Knife, Cutting Board",
		// 	Ingredients:   "1 chicken breast, 1 wrap, 1/2 cup lettuce, 1/4 cup tomato, 1/4 cup cucumber, 2 tbsp sauce",
		// },
		// {
		// 	Name:        "Tomato Basil Pasta",
		// 	Category:    "Dinner",
		// 	Description: "A simple and fresh pasta dish with tomatoes and basil.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Use canned tomatoes and skip the fresh basil to reduce prep time.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Cook the pasta according to package instructions. \n2. Heat oil in a pan, and add chopped tomatoes and garlic. Cook for 5 minutes. \n3. Add cooked pasta to the pan, toss with the sauce, and top with basil.",
		// 	ImageUrl:      "https://example.com/tomato_basil_pasta.jpg",
		// 	Calories:      350,
		// 	Protein:       12,
		// 	CookTime:      15,
		// 	PrepTime:      5,
		// 	WashingEffort: 2,
		// 	PeelingEffort: 0,
		// 	CuttingEffort: 3,
		// 	ItemsRequired: "Pot, Pan, Knife",
		// 	Ingredients:   "2 cups pasta, 2 tomatoes, 1 garlic clove, 2 tbsp olive oil, 1/4 cup basil",
		// },
		// {
		// 	Name:        "Egg Salad Sandwich",
		// 	Category:    "Lunch",
		// 	Description: "A creamy egg salad sandwich with a hint of mustard.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Use fewer eggs and skip the mustard for a lighter version.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Boil the eggs for 10 minutes. \n2. Peel and chop the eggs, then mix with mayo, mustard, and seasoning. \n3. Spread the egg mixture on bread and serve.",
		// 	ImageUrl:      "https://example.com/egg_salad_sandwich.jpg",
		// 	Calories:      300,
		// 	Protein:       12,
		// 	CookTime:      10,
		// 	PrepTime:      5,
		// 	WashingEffort: 2,
		// 	PeelingEffort: 5,
		// 	CuttingEffort: 1,
		// 	ItemsRequired: "Pot, Knife",
		// 	Ingredients:   "3 eggs, 2 tbsp mayo, 1 tsp mustard, salt, pepper, 2 slices bread",
		// },
		// {
		// 	Name:        "Tuna Salad",
		// 	Category:    "Lunch",
		// 	Description: "A light and refreshing tuna salad with a citrus dressing.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Use only tuna and dressing, skipping the additional vegetables.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Drain the tuna. \n2. Chop the vegetables. \n3. Mix tuna, vegetables, and dressing in a bowl and serve.",
		// 	ImageUrl:      "https://example.com/tuna_salad.jpg",
		// 	Calories:      250,
		// 	Protein:       25,
		// 	CookTime:      0,
		// 	PrepTime:      5,
		// 	WashingEffort: 1,
		// 	PeelingEffort: 0,
		// 	CuttingEffort: 2,
		// 	ItemsRequired: "Bowl, Knife",
		// 	Ingredients:   "1 can tuna, 1/2 cup lettuce, 1/4 cup cucumber, 1/4 cup tomato, 2 tbsp dressing",
		// },
		// {
		// 	Name:        "Mushroom Risotto",
		// 	Category:    "Dinner",
		// 	Description: "A creamy mushroom risotto with Parmesan cheese.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Use pre-sliced mushrooms and skip the Parmesan for a faster version.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Heat oil in a pot, and add rice. \n2. Slowly add broth while stirring. \n3. Cook mushrooms in a separate pan, then add to the rice. \n4. Stir in Parmesan and serve.",
		// 	ImageUrl:      "https://example.com/mushroom_risotto.jpg",
		// 	Calories:      400,
		// 	Protein:       15,
		// 	CookTime:      20,
		// 	PrepTime:      5,
		// 	WashingEffort: 3,
		// 	PeelingEffort: 0,
		// 	CuttingEffort: 2,
		// 	ItemsRequired: "Pot, Pan, Spoon",
		// 	Ingredients:   "1 cup rice, 2 cups broth, 1 cup mushrooms, 1/4 cup Parmesan, 2 tbsp oil",
		// },
		// {
		// 	Name:        "Grilled Cheese Sandwich",
		// 	Category:    "Snack",
		// 	Description: "A classic grilled cheese sandwich with melted cheese.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Use less cheese and butter for a lighter version.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Butter one side of each slice of bread. \n2. Place cheese between the slices. \n3. Grill on a pan until golden brown and cheese is melted.",
		// 	ImageUrl:      "https://example.com/grilled_cheese.jpg",
		// 	Calories:      300,
		// 	Protein:       10,
		// 	CookTime:      5,
		// 	PrepTime:      2,
		// 	WashingEffort: 1,
		// 	PeelingEffort: 0,
		// 	CuttingEffort: 0,
		// 	ItemsRequired: "Pan, Spatula",
		// 	Ingredients:   "2 slices bread, 1/2 cup cheese, 1 tbsp butter",
		// },
		// {
		// 	Name:        "Greek Yogurt Parfait",
		// 	Category:    "Breakfast",
		// 	Description: "A healthy and quick Greek yogurt parfait with fruit and granola.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Use fewer toppings for a quicker version.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Layer Greek yogurt, granola, and fruit in a bowl or glass. \n2. Serve immediately.",
		// 	ImageUrl:      "https://example.com/greek_yogurt_parfait.jpg",
		// 	Calories:      200,
		// 	Protein:       15,
		// 	CookTime:      0,
		// 	PrepTime:      3,
		// 	WashingEffort: 1,
		// 	PeelingEffort: 1,
		// 	CuttingEffort: 1,
		// 	ItemsRequired: "Bowl, Spoon",
		// 	Ingredients:   "1 cup Greek yogurt, 1/2 cup granola, 1/2 cup berries",
		// },
		// {
		// 	Name:        "Shrimp Tacos",
		// 	Category:    "Dinner",
		// 	Description: "Spicy shrimp tacos with a refreshing lime slaw.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Use pre-cooked shrimp and skip the slaw for a quicker version.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Cook shrimp with seasoning. \n2. Prepare the slaw by mixing cabbage and lime juice. \n3. Assemble tacos with shrimp, slaw, and toppings.",
		// 	ImageUrl:      "https://example.com/shrimp_tacos.jpg",
		// 	Calories:      350,
		// 	Protein:       20,
		// 	CookTime:      10,
		// 	PrepTime:      5,
		// 	WashingEffort: 2,
		// 	PeelingEffort: 0,
		// 	CuttingEffort: 2,
		// 	ItemsRequired: "Pan, Knife",
		// 	Ingredients:   "1/2 lb shrimp, 4 tortillas, 1/2 cup cabbage, 1 lime, 1/4 cup salsa",
		// },
		// {
		// 	Name:        "Fruit Smoothie",
		// 	Category:    "Snack",
		// 	Description: "A quick and healthy fruit smoothie.",
		// 	LightVersionInstructions: sql.NullString{
		// 		String: "Use fewer fruits for a quicker version.",
		// 		Valid:  true,
		// 	},
		// 	Instructions:  "1. Add fruit, yogurt, and juice to a blender. \n2. Blend until smooth and serve.",
		// 	ImageUrl:      "https://example.com/fruit_smoothie.jpg",
		// 	Calories:      150,
		// 	Protein:       5,
		// 	CookTime:      0,
		// 	PrepTime:      3,
		// 	WashingEffort: 1,
		// 	PeelingEffort: 1,
		// 	CuttingEffort: 1,
		// 	ItemsRequired: "Blender",
		// 	Ingredients:   "1 banana, 1/2 cup berries, 1/2 cup yogurt, 1/2 cup juice",
		// },
		{
			Name:        "Spicy Fusilli Calabrese with Pine Nuts",
			Category:    "vegetarian",
			Servings:    2,
			Description: "A simple and delicious vegetarian Italian dish featuring spicy fusilli, baby spinach, cherry tomatoes, and pine nuts.",
			LightVersionInstructions: sql.NullString{
				String: "For a lighter version, reduce the amount of butter and cheese, and skip the chili if you prefer it less spicy.",
				Valid:  true,
			},
			Instructions: `1. Boil water in a kettle. Finely dice the onion and peel the garlic. Roast pine nuts in a large pot for 2-3 minutes without oil until they begin to brown. Remove and set aside.
2. In the same pot, heat 1 tbsp olive oil. Add onions and pressed garlic, and sauté for 1-2 minutes until translucent.
3. Add 500 ml hot water, fusilli, vegetable broth powder, and 0.5 tsp salt to the pot. Bring to a boil, cover, and simmer for 9-11 minutes until pasta is al dente, stirring occasionally.
4. Halve the cherry tomatoes and slice the chili pepper. When the pasta is al dente, stir in cream cheese, Calabrese pesto, 1 tbsp butter, and chili mix. Stir in cherry tomatoes and baby spinach, cooking for 1-2 minutes until warmed through. Season with salt and pepper.
5. Plate the pasta and top with roasted pine nuts, grated Italian hard cheese, and chili slices if desired.`,
			ImageUrl:      "meal3.png", // Replace with actual image URL
			Calories:      861,
			Protein:       31,
			CookTime:      11,
			PrepTime:      10,
			TotalTime:     25,
			WashingEffort: 3, // Washing vegetables (medium effort)
			PeelingEffort: 2, // Peeling garlic and onions (low effort)
			CuttingEffort: 4, // Cutting onions, tomatoes, and chili (medium effort)
			ItemsRequired: "Pot with Lid, Large Pot, Plastic Wrap",
			Ingredients:   "1 Onion, 1 Garlic Clove, 50g Calabrese Pesto, 270g Fusilli, 10g Pine Nuts, 4g Vegetable Broth Powder, 40g Grated Italian Hard Cheese, 100g Baby Spinach, 100g Cream Cheese, 125g Cherry Tomatoes, 1 Red Chili Pepper, 1g Mild Chili Mix, Salt, Pepper, 1 tbsp Olive Oil, 1 tbsp Butter, 500 ml Water",
			TotalEffort:   3, // Averaged overall effort
		},
	}

	for _, mealParams := range mealsParams {
		err := d.InsertMeal(ctx, mealParams)
		if err != nil {
			fmt.Printf("Error creating meal: %v\n", err)
			return
		}
		fmt.Printf("Created meal\n")
	}

}
