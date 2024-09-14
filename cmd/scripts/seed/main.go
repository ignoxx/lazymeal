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

	mealsParams := []sqlc.InsertMealParams{}
	mealsParams = append(mealsParams, mealWeek1()...)
	mealsParams = append(mealsParams, mealWeek2()...)
	mealsParams = append(mealsParams, mealWeek2Part2()...)

	fmt.Println("Total Meals inserting..", len(mealsParams))

	for _, mealParams := range mealsParams {
		_, err := d.InsertMeal(ctx, mealParams)
		if err != nil {
			fmt.Printf("Error creating meal: %v\n", err)
			return
		}
		fmt.Printf("Created meal\n")
	}
}

func mealWeek2Part2() []sqlc.InsertMealParams {
	return []sqlc.InsertMealParams{
		{
			Name:        "Chicken Salad (Salpicão)",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    4,
			Description: "A traditional and flavorful chicken salad made with shredded chicken, vegetables, raisins, and mayonnaise, topped with crunchy potato sticks. This dish is perfect for a light lunch or dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Use low-fat mayonnaise or reduce the amount to make a lighter version. You can also skip the potato sticks for fewer calories.",
				Valid:  true,
			},
			Instructions:  "In a large bowl, mix the cooked and shredded chicken, grated carrots, peas, corn, raisins, diced apple, mayonnaise, and fresh parsley. Stir until well combined. Once mixed, top the salad with potato sticks. Serve immediately and enjoy!",
			ImageUrl:      "chicken-salad-salpicao.jpg",
			Calories:      350, // Approximate per serving
			Protein:       25,
			CookTime:      0,
			PrepTime:      10,
			TotalTime:     10,
			WashingEffort: 2,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Large mixing bowl, spoon, grater, knife, cutting board",
			Ingredients:   "500g cooked shredded chicken breast,2 large carrots grated,1 can peas,1 can corn,200g raisins,1 apple diced,mayonnaise to taste,fresh parsley to taste,potato sticks to taste",
			TotalEffort:   3,
		},
		{
			Name:        "Apple or Banana with Water",
			Category:    "Snacks, Healthy",
			Servings:    1,
			Description: "A simple and satisfying snack of either an apple or banana, followed by a glass of water. Perfect for a quick, healthy boost that keeps you feeling full.",
			LightVersionInstructions: sql.NullString{
				String: "No simplification needed. Simply eat the fruit and drink a glass of water.",
				Valid:  true,
			},
			Instructions:  "Pick either an apple or banana. Wash the apple or peel the banana, and enjoy it as a snack. Follow with a glass of water to feel satisfied.",
			ImageUrl:      "apple-banana-water.png",
			Calories:      100, // Approx. for one medium fruit
			Protein:       1,
			CookTime:      0,
			PrepTime:      1,
			TotalTime:     1,
			WashingEffort: 1,
			PeelingEffort: 1,
			CuttingEffort: 0,
			ItemsRequired: "Glass, knife (optional for apple)",
			Ingredients:   "1 apple or 1 banana,1 glass of water",
			TotalEffort:   1,
		},

		{
			Name:        "Creamy Pasta",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    1,
			Description: "A rich and delicious creamy pasta dish made with chicken breast, tomato sauce, and mozzarella cheese. Cooked quickly in a pressure cooker, it's a perfect hearty meal for busy days.",
			LightVersionInstructions: sql.NullString{
				String: "For a lighter version, reduce the amount of heavy cream and cheese, or substitute with low-fat alternatives. You can also use pre-cooked chicken to save time.",
				Valid:  true,
			},
			Instructions:  "In a pressure cooker, sauté the garlic and onion with diced chicken until lightly browned. Add your preferred seasonings. Pour in the tomato sauce, penne pasta, and enough water to cover the ingredients. Seal the pressure cooker and cook under pressure for 10 minutes (or less). After releasing the pressure, open the cooker and stir in the heavy cream and mozzarella cheese. Adjust the salt and serve warm.",
			ImageUrl:      "creamy-pasta.jpg",
			Calories:      750,
			Protein:       40,
			CookTime:      10,
			PrepTime:      10,
			TotalTime:     20,
			WashingEffort: 4,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Pressure cooker, knife, cutting board, spoon",
			Ingredients:   "100g penne pasta,1 clove garlic,1/2 onion,1/2 cup tomato sauce,150g diced chicken breast,1/4 cup heavy cream,50g mozzarella cheese,water,salt,pepper,Optional: other seasonings (paprika, Italian herbs)",
			TotalEffort:   3,
		},
		{
			Name:        "Chicken Salad",
			Category:    "Lunch, Dinner, Healthy",
			Servings:    1,
			Description: "A delicious and healthy chicken salad with mixed vegetables and a tangy lemon dressing. A perfect light meal with high protein content.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-made salad mix or eat raw vegetables like sweet pepper, cucumber, or tomato with a slice of bread for a quick sandwich.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Chop vegetables (red beans, corn, cucumber, carrot, spinach) and mix with olive oil, lemon juice, salt, and pepper. Optional: Add tuna for extra protein.",
			ImageUrl:      "chicken-salad.png",
			Calories:      400,
			Protein:       45,
			CookTime:      10,
			PrepTime:      10,
			TotalTime:     20,
			WashingEffort: 4,
			PeelingEffort: 2,
			CuttingEffort: 5,
			ItemsRequired: "Air fryer or oven, knife, cutting board, salad bowl",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,1/4 cup red beans,1/4 cup corn,1/2 cucumber,1 carrot,1 cup spinach,1 lemon,2 tbsp olive oil,salt,pepper,Optional: 1 can tuna",
			TotalEffort:   4,
		},
		{
			Name:        "Chicken Rice",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    1,
			Description: "A satisfying chicken rice meal with sautéed onions and carrots, perfect for a hearty and comforting lunch or dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Use bag rice and skip sautéing the onion and carrot for a quicker version.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Sauté onion and grated carrot in olive oil, then add rice and water. Cook for 11 minutes over medium heat with the pot covered.",
			ImageUrl:      "chicken-rice.png",
			Calories:      550,
			Protein:       42,
			CookTime:      15,
			PrepTime:      10,
			TotalTime:     25,
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Air fryer or oven, knife, cutting board, pot with lid, grater",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,1/2 cup basmati rice,1/2 onion,1-2 carrots,1 tbsp olive oil",
			TotalEffort:   3,
		},
		{
			Name:        "Chicken Pasta",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    1,
			Description: "A simple and hearty chicken pasta dish, perfect for a filling meal. Combine chicken with your choice of pasta and sauce.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-cooked pasta and store-bought sauce to save time.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Boil pasta according to package instructions. Toss cooked pasta with olive oil and your choice of sauce (tomato or cream).",
			ImageUrl:      "chicken-pasta.png",
			Calories:      600,
			Protein:       43,
			CookTime:      15,
			PrepTime:      10,
			TotalTime:     25,
			WashingEffort: 4,
			PeelingEffort: 1,
			CuttingEffort: 3,
			ItemsRequired: "Air fryer or oven, pot, knife, cutting board",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,80g pasta,1 tbsp olive oil,1/2 cup tomato sauce or cream sauce",
			TotalEffort:   3,
		},
		{
			Name:        "Chicken Mashed Potatoes",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    1,
			Description: "A classic combination of chicken and mashed potatoes, with creamy buttered potatoes making it a comforting and filling meal.",
			LightVersionInstructions: sql.NullString{
				String: "Use instant mashed potatoes to reduce cooking time.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Boil peeled and cubed potatoes, then mash with butter, milk, salt, and pepper.",
			ImageUrl:      "chicken-mashed-potatoes.png",
			Calories:      550,
			Protein:       42,
			CookTime:      20,
			PrepTime:      10,
			TotalTime:     30,
			WashingEffort: 4,
			PeelingEffort: 5,
			CuttingEffort: 3,
			ItemsRequired: "Air fryer or oven, pot, potato masher, knife, cutting board",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,2 medium potatoes,1 tbsp butter,1/4 cup milk,salt,pepper",
			TotalEffort:   4,
		},
	}
}

func mealWeek2() []sqlc.InsertMealParams {
	return []sqlc.InsertMealParams{
		{
			Name:        "Tapioca or Crepioca with Cheese, Turkey, and Veggies",
			Category:    "Breakfast, Lunch, Healthy",
			Servings:    1,
			Description: "A quick and healthy mini pancake made from tapioca or a mix of tapioca and egg (Crepioca), filled with cheese, turkey slices, and veggies. Perfect for a fast, nutritious meal.",
			LightVersionInstructions: sql.NullString{
				String: "If you don't have tapioca, you can use chickpea flour to make a quick flatbread. You can also skip the veggies and turkey for an even simpler version.",
				Valid:  true,
			},
			Instructions:  "In a bowl, mix 2 tbsp of tapioca flour (or chickpea flour) with 1 egg to make the crepioca. Heat a non-stick pan over medium heat and pour in the mixture, cooking it like a pancake for 2-3 minutes on each side until firm. Once cooked, add slices of cheese, turkey, and any sautéed or raw veggies you like (e.g., spinach, tomatoes, or peppers). Fold the crepioca in half and serve warm. If using just tapioca, simply heat the tapioca flour in a pan to form the flatbread, and then fill it with the toppings.",
			ImageUrl:      "tapioca-crepioca.jpg",
			Calories:      350,
			Protein:       20,
			CookTime:      5,
			PrepTime:      5,
			TotalTime:     10,
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Non-stick pan, bowl, spatula",
			Ingredients:   "2 tbsp tapioca flour or chickpea flour,1 egg,2 slices cheese,2 slices turkey,Optional: veggies (spinach,tomatoes,peppers),salt,pepper",
			TotalEffort:   2,
		},
		{
			Name:        "Weekly Meal Prep",
			Category:    "Meal Prep, Healthy, Comfort Food",
			Servings:    0, // Number of servings will vary based on meal type
			Description: "A method of batch cooking meals once a week and freezing them in small portions for easy reheating throughout the week. Ideal for people with busy schedules or those looking to save time in the kitchen.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought frozen meals or prepare simpler dishes that require less cooking time. Focus on easy-to-reheat ingredients like pre-cooked grains, roasted vegetables, and proteins like chicken or beans.",
				Valid:  true,
			},
			Instructions:  "Set aside one day each week (e.g., Sunday) to prepare meals in large quantities. Cook multiple servings of different ingredients like grains (rice, quinoa, pasta), proteins (chicken, beef, tofu), and vegetables. Once the food is cooked, divide it into small, individual portions and store them in airtight containers or freezer bags. Label each container with the date and contents, then freeze. Throughout the week, simply reheat the frozen meals in the microwave or on the stovetop. You can also combine different portions to create variety in your meals. This method ensures you always have ready-to-eat meals with minimal daily effort.",
			ImageUrl:      "weekly-meal-prep.jpg",
			Calories:      0,
			Protein:       0,
			CookTime:      0, // Total time varies by recipe
			PrepTime:      0,
			TotalTime:     0,
			WashingEffort: 6, // Average effort for meal prep
			PeelingEffort: 5,
			CuttingEffort: 6,
			ItemsRequired: "Large pots, pans, containers for storage, freezer bags, knife, cutting board, oven, stove, microwave for reheating",
			Ingredients:   "Various ingredients like rice, quinoa, pasta, vegetables, chicken, beef, tofu, beans, herbs, spices",
			TotalEffort:   5,
		},
		{
			Name:        "Overnight Oats with Yogurt and Fruits",
			Category:    "Breakfast, Healthy, Meal Prep",
			Servings:    1,
			Description: "A quick and nutritious overnight oats recipe made with milk, yogurt, and fresh fruits. Perfect for a healthy breakfast or snack that's ready to grab and go the next morning.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought fruit yogurt or pre-cut fruits to save time. You can also use plant-based milk and yogurt for a dairy-free version.",
				Valid:  true,
			},
			Instructions:  "The night before, combine 1/2 cup of rolled oats with 1 cup of milk in a jar or bowl. Cover and let it rest in the refrigerator overnight. The next morning, take out the oats and layer them with yogurt and your choice of fresh fruits such as berries, bananas, or apples. Serve immediately or store in the fridge for later.",
			ImageUrl:      "overnight-oats.png",
			Calories:      350,
			Protein:       15,
			CookTime:      0,
			PrepTime:      5,
			TotalTime:     5,
			WashingEffort: 1,
			PeelingEffort: 2,
			CuttingEffort: 2,
			ItemsRequired: "Bowl or jar, spoon, knife",
			Ingredients:   "1/2 cup rolled oats,1 cup milk,1/2 cup yogurt,Optional: fresh fruits (berries,banana,apple),honey or sweetener to taste",
			TotalEffort:   2,
		},
		{
			Name:        "Stuffed Zucchini with Meat and Cheese",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    1,
			Description: "A simple and quick meal of stuffed zucchini filled with ground meat and cheese, baked in the oven for a hearty and satisfying dish. Perfect for lazy cooking when you need to be done fast.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-cooked or leftover meat, and skip the oven by microwaving for a few minutes to melt the cheese.",
				Valid:  true,
			},
			Instructions:  "Preheat the oven to 200°C (400°F). Cut the zucchini in half lengthwise and scoop out some of the flesh to create a hollow center. Fill the zucchini halves with cooked ground meat and top with shredded cheese. Bake in the oven for 10 minutes or until the cheese is melted and bubbly. Serve warm.",
			ImageUrl:      "stuffed-zucchini.png",
			Calories:      400,
			Protein:       25,
			CookTime:      10,
			PrepTime:      5,
			TotalTime:     15,
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 3,
			ItemsRequired: "Oven, knife, spoon, baking sheet",
			Ingredients:   "1 medium zucchini,100g ground meat (beef or chicken),50g shredded cheese (cheddar or mozzarella),salt,pepper",
			TotalEffort:   2,
		},
		{
			Name:        "Pasta Salad",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    1,
			Description: "A quick and easy pasta salad with tuna, corn, peas, and olives, mixed with mayonnaise for a creamy, delicious meal. Perfect for a light and satisfying lunch or dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought pasta salad or skip the tuna for a vegetarian version. Alternatively, you can use a lighter dressing like yogurt instead of mayonnaise.",
				Valid:  true,
			},
			Instructions:  "Cook the penne pasta in boiling salted water until al dente. Drain the pasta and add mayonnaise, tuna, corn, peas, olives, and fresh parsley or chives. Mix well, chill in the refrigerator for a few minutes, and serve.",
			ImageUrl:      "pasta-salad.png",
			Calories:      600,
			Protein:       25,
			CookTime:      10,
			PrepTime:      5,
			TotalTime:     15,
			WashingEffort: 3,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Pot, strainer, mixing bowl, spoon",
			Ingredients:   "150g penne pasta,2 tbsp mayonnaise,1 can tuna,1/4 cup corn,1/4 cup peas,1/4 cup olives,1 tbsp fresh parsley or chives",
			TotalEffort:   2,
		},
		{
			Name:        "Chicken Couscous",
			Category:    "Lunch, Dinner, Healthy",
			Servings:    1,
			Description: "A quick and easy chicken and couscous meal, perfect for a light and healthy lunch or dinner with Mediterranean flavors.",
			LightVersionInstructions: sql.NullString{
				String: "Use instant couscous and skip adding vegetables for a quicker version.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Prepare couscous by adding boiling water, covering, and letting it steam for 5 minutes. Mix in olive oil and your choice of chopped vegetables like tomatoes, cucumbers, and bell peppers.",
			ImageUrl:      "chicken-couscous.png",
			Calories:      450,
			Protein:       40,
			CookTime:      10,
			PrepTime:      10,
			TotalTime:     20,
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 4,
			ItemsRequired: "Air fryer or oven, pot, knife, cutting board",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,1/2 cup couscous,1 tbsp olive oil,Optional: tomatoes,cucumbers,bell peppers",
			TotalEffort:   3,
		},
		{
			Name:        "Chicken Roasted Vegetables",
			Category:    "Lunch, Dinner, Healthy",
			Servings:    1,
			Description: "A healthy and flavorful chicken dish served with a variety of roasted vegetables like bell peppers, zucchini, and carrots, perfect for a balanced meal.",
			LightVersionInstructions: sql.NullString{
				String: "Use frozen roasted vegetables to save time, or choose just one vegetable like zucchini for a simpler version.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Chop vegetables (bell peppers, zucchini, carrots) and toss with olive oil, salt, and pepper. Roast in oven at 200°C for 20-25 minutes.",
			ImageUrl:      "chicken-roasted-vegetables.png",
			Calories:      400,
			Protein:       40,
			CookTime:      25,
			PrepTime:      10,
			TotalTime:     35,
			WashingEffort: 4,
			PeelingEffort: 2,
			CuttingEffort: 4,
			ItemsRequired: "Air fryer or oven, baking sheet, knife, cutting board",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,1 bell pepper,1 zucchini,2 carrots,1 tbsp olive oil,salt,pepper",
			TotalEffort:   4,
		},
		{
			Name:        "Chicken Pita Bread with Hummus",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    1,
			Description: "A delicious Mediterranean-inspired meal with grilled chicken, pita bread, and creamy hummus, ideal for a flavorful and satisfying dish.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought pita bread and hummus to simplify the meal.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Warm pita bread in the oven or microwave. Serve chicken with pita bread and hummus on the side. Optionally, add fresh vegetables like cucumber and tomato for extra flavor.",
			ImageUrl:      "chicken-pita-hummus.png",
			Calories:      500,
			Protein:       45,
			CookTime:      10,
			PrepTime:      5,
			TotalTime:     15,
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Air fryer or oven, microwave, knife, cutting board",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,1 pita bread,1/4 cup hummus,Optional: cucumber,tomato",
			TotalEffort:   2,
		},
		{
			Name:        "Chicken Feta and Tomato Salad",
			Category:    "Lunch, Dinner, Healthy",
			Servings:    1,
			Description: "A fresh and tangy Greek-inspired chicken salad with feta cheese, tomatoes, cucumbers, and a drizzle of olive oil, perfect for a light and refreshing meal.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-cut salad mix and feta crumbles for a quicker version.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Chop tomatoes, cucumbers, and red onion. Mix with crumbled feta cheese and olive oil. Serve with chicken on the side or mixed in.",
			ImageUrl:      "chicken-feta-tomato-salad.png",
			Calories:      450,
			Protein:       43,
			CookTime:      10,
			PrepTime:      10,
			TotalTime:     20,
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 4,
			ItemsRequired: "Air fryer or oven, knife, cutting board, salad bowl",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,2 tomatoes,1 cucumber,1/4 red onion,50g feta cheese,2 tbsp olive oil,salt,pepper",
			TotalEffort:   3,
		},
		{
			Name:        "Chicken Garlic Bread",
			Category:    "Lunch, Dinner, Comfort Food",
			Servings:    1,
			Description: "A comforting meal with crispy garlic bread served alongside juicy chicken, perfect for a satisfying lunch or dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-made garlic bread to save time.",
				Valid:  true,
			},
			Instructions:  "Marinade chicken with sweet pepper, garlic, black pepper, soy sauce, and optional cayenne pepper. Cook in air fryer or oven. Prepare garlic bread by spreading butter and minced garlic on slices of bread, then toasting in the oven. Serve chicken with garlic bread on the side.",
			ImageUrl:      "chicken-garlic-bread.png",
			Calories:      600,
			Protein:       45,
			CookTime:      10,
			PrepTime:      10,
			TotalTime:     20,
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Air fryer or oven, knife, cutting board, toaster",
			Ingredients:   "200g chicken breast,1 tsp sweet pepper,1 clove garlic,1 tsp black pepper,1 tsp soy sauce,Optional: 1/4 tsp cayenne pepper,2 slices bread,1 tbsp butter,1 clove garlic",
			TotalEffort:   2,
		},
	}
}

func mealWeek1() []sqlc.InsertMealParams {
	return []sqlc.InsertMealParams{
		{
			Name:        "Chicken Fajita Mac and Cheese",
			Category:    "Main Course, Comfort Food, Dinner",
			Servings:    4,
			Description: "This Chicken Fajita Mac and Cheese combines the spicy flavors of fajitas with creamy macaroni and cheese for a comforting and satisfying dish. Perfect for a weeknight meal that the whole family will love.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-cooked chicken and pre-shredded cheese to reduce prep time.",
				Valid:  true,
			},
			Instructions:  "Fry onion, peppers, and garlic in olive oil until translucent. Add chicken, fajita seasoning, and salt, cooking until chicken is slightly browned. Add cream, chicken stock, and macaroni, cooking on low for 20 minutes. Stir in cheddar cheese and top with roasted peppers and parsley.",
			ImageUrl:      "chicken-fajita-mac-and-cheese.jpg",
			Calories:      650, // Estimated per serving
			Protein:       30,  // Estimated per serving
			CookTime:      20,  // Time to cook the pasta and finish dish
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     30,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 1,
			CuttingEffort: 3,
			ItemsRequired: "Skillet, Pot, Wooden spoon, Knife",
			Ingredients:   "500g macaroni, 2 cups chicken stock, 1/2 cup heavy cream, 1 packet fajita seasoning, 1 tsp salt, 3 diced chicken breasts, 2 tbsp olive oil, 1 small diced onion, 2 diced red peppers, 2 minced garlic cloves, 1 cup cheddar cheese, Parsley for garnish",
			TotalEffort:   2,
		},

		{
			Name:        "Chicken Enchilada Casserole",
			Category:    "Main Course, Dinner, Comfort Food",
			Servings:    4,
			Description: "This Chicken Enchilada Casserole is a cheesy, hearty, and delicious dish that's easy to prepare. Layered with enchilada sauce, shredded chicken, corn tortillas, and Monterey Jack cheese, it is perfect for a quick and satisfying dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Use rotisserie chicken and canned enchilada sauce to save time.",
				Valid:  true,
			},
			Instructions:  "Cook chicken breasts in enchilada sauce for 20 minutes. Shred the chicken and layer with tortillas, sauce, and cheese in a casserole dish. Repeat layers and bake at 375F for 20-30 minutes until bubbly and browned on top.",
			ImageUrl:      "chicken-enchilada-casserole.jpg",
			Calories:      600, // Estimated per serving
			Protein:       35,  // Estimated per serving
			CookTime:      30,  // Time to bake the casserole
			PrepTime:      15,  // Time to cook chicken and assemble
			TotalTime:     45,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Casserole dish, Pot, Forks",
			Ingredients:   "14 oz enchilada sauce, 3 cups shredded Monterey Jack cheese, 6 corn tortillas, 2 chicken breasts",
			TotalEffort:   2,
		},

		{
			Name:        "Bread Omelette",
			Category:    "Breakfast, Snacks",
			Servings:    1,
			Description: "A quick and easy bread omelette, perfect for breakfast or a light snack. This dish combines eggs and bread for a hearty, filling meal in minutes.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-sliced bread and pre-beaten eggs to speed up the process.",
				Valid:  true,
			},
			Instructions:  "Make and enjoy.",
			ImageUrl:      "bread-omelette.jpg",
			Calories:      250, // Estimated per serving
			Protein:       10,  // Estimated per serving
			CookTime:      5,   // Quick cook time for the omelette
			PrepTime:      2,   // Minimal preparation required
			TotalTime:     7,   // Total time: prep + cook
			WashingEffort: 1,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Skillet, Spatula",
			Ingredients:   "2 slices of bread, 2 eggs, 1/2 tsp salt",
			TotalEffort:   1,
		},

		{
			Name:        "Beetroot Soup (Borscht)",
			Category:    "Soups, Vegetarian",
			Servings:    4,
			Description: "A traditional Eastern European dish, Beetroot Soup (Borscht) is a vibrant and earthy soup made with fresh beetroot, potatoes, and beans. It's a warm, nutritious, and flavorful dish that's perfect for a light yet hearty meal.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the beans and use pre-cooked beetroot to reduce cooking time.",
				Valid:  true,
			},
			Instructions:  "Chop the beetroot and boil in water with a stock cube for 15 minutes. Add diced potatoes and cook until soft. Finally, add beans and cook for 5 more minutes. Serve hot, garnished with dill.",
			ImageUrl:      "beetroot-soup-borscht.jpg",
			Calories:      200, // Estimated per serving
			Protein:       6,   // Estimated per serving
			CookTime:      20,  // Time to cook the soup
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     30,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 3,
			CuttingEffort: 3,
			ItemsRequired: "Soup pot, Knife, Cutting board",
			Ingredients:   "3 beetroots, 4 tbsp olive oil, 1 chicken stock cube, 6 cups water, 3 potatoes, 1 can cannellini beans, Fresh dill for garnish",
			TotalEffort:   3,
		},

		{
			Name:        "Blini Pancakes",
			Category:    "Breakfast, Appetizers",
			Servings:    4,
			Description: "Blini pancakes are small, fluffy pancakes made with buckwheat flour and yeast, traditionally served with toppings like sour cream, caviar, or smoked salmon. Perfect for breakfast or as a savory appetizer.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-made pancake batter for a quick version. Skip the yeast and let the batter rest for just 10 minutes before cooking.",
				Valid:  true,
			},
			Instructions:  "In a bowl, whisk buckwheat flour, all-purpose flour, salt, and yeast. Add warm milk and mix until smooth. Let rise for 1 hour. Stir in melted butter and egg yolk. Fold in whisked egg white. Let batter rest for 20 minutes. Heat butter in a skillet, drop batter, cook for 1 minute, flip, and cook for 30 seconds more. Repeat with remaining batter.",
			ImageUrl:      "blini-pancakes.jpg",
			Calories:      200, // Estimated per serving
			Protein:       6,   // Estimated per serving
			CookTime:      15,  // Time to cook blini pancakes
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     25,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Mixing bowl, Skillet, Whisk",
			Ingredients:   "1/2 cup buckwheat flour, 2/3 cup flour, 1/2 tsp salt, 1 tsp yeast, 1 cup milk, 2 tbsp butter, 1 egg (separated)",
			TotalEffort:   2,
		},

		{
			Name:        "Bigos (Hunters Stew)",
			Category:    "Main Course, Dinner, Comfort Food",
			Servings:    6,
			Description: "A hearty Polish dish, Bigos (Hunters Stew) is a flavorful combination of meats, cabbage, and sauerkraut, simmered with mushrooms, garlic, and spices. This traditional stew is perfect for a cold day, offering deep, rich flavors.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the initial browning steps by using pre-cooked meats, and opt for pre-shredded cabbage and sauerkraut to save time.",
				Valid:  true,
			},
			Instructions:  "Preheat the oven to 350F (175C). Cook bacon and kielbasa in a large pot until browned. Remove and transfer to a casserole dish. Lightly flour and brown pork in the same pot, then transfer to the casserole. Cook garlic, onion, carrots, mushrooms, cabbage, and sauerkraut. Deglaze with red wine and add spices. Combine with the meat, cover, and bake for 2 1/2 to 3 hours.",
			ImageUrl:      "bigos-hunters-stew.jpg",
			Calories:      500, // Estimated per serving
			Protein:       25,  // Estimated per serving
			CookTime:      180, // Time to bake the stew
			PrepTime:      20,  // Time to prepare ingredients
			TotalTime:     200, // Total time: prep + cook
			WashingEffort: 4,
			PeelingEffort: 3,
			CuttingEffort: 3,
			ItemsRequired: "Casserole dish, Large pot, Slotted spoon, Knife",
			Ingredients:   "2 slices bacon, 1 lb kielbasa, 1 lb pork, 1/4 cup flour, 3 garlic cloves, 1 diced onion, 1 1/2 cup mushrooms, 4 cups cabbage, 1 jar sauerkraut, 1/4 cup red wine, 1 bay leaf, 1 tsp basil, 1 tsp marjoram, 1 tbsp paprika, 1/8 tsp caraway seeds",
			TotalEffort:   3,
		},

		{
			Name:        "Breakfast Potatoes",
			Category:    "Breakfast, Brunch",
			Servings:    2,
			Description: "A hearty and savory breakfast dish, these crispy breakfast potatoes are tossed with bacon, garlic, and a touch of maple syrup for a perfect balance of flavors. Serve with a sunny-side-up egg for a complete meal.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the bacon and use pre-chopped potatoes to save time. You can also use pre-minced garlic to reduce effort.",
				Valid:  true,
			},
			Instructions:  "Freeze bacon for easy chopping. Wash and dice potatoes, then soak in water to prevent browning. Heat oil in a skillet over medium-high, drain potatoes, and cook with seasoning for 10 minutes. Add chopped bacon, cook for 5-6 minutes until crispy, then add garlic. Drizzle with maple syrup before serving, and toss for a caramelized finish.",
			ImageUrl:      "breakfast-potatoes.jpg",
			Calories:      350, // Estimated per serving
			Protein:       10,  // Estimated per serving
			CookTime:      15,  // Time to cook potatoes and bacon
			PrepTime:      10,  // Time to chop and prepare ingredients
			TotalTime:     25,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Skillet, Knife, Bowl",
			Ingredients:   "3 medium potatoes, 1 tbsp olive oil, 2 strips bacon, 1 garlic clove, 1 tbsp maple syrup, Parsley for garnish, Salt, Pepper, Allspice",
			TotalEffort:   3,
		},

		{
			Name:        "Baked Salmon with Fennel & Tomatoes",
			Category:    "Main Course, Dinner, Healthy",
			Servings:    2,
			Description: "A healthy and flavorful baked salmon dish with fennel and cherry tomatoes. This dish is packed with fresh flavors, combining tender salmon with the aromatic fennel and sweet roasted tomatoes. Perfect for a nutritious and simple dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Skip pre-cooking the fennel in boiling water and directly roast it with the tomatoes and salmon to save time.",
				Valid:  true,
			},
			Instructions:  "Heat oven to 180C/fan 160C/gas 4. Trim the fronds from the fennel and set aside. Cut the fennel bulbs in half, then cut each half into 3 wedges. Cook in boiling salted water for 10 mins, then drain well. Chop the fennel fronds roughly, then mix with the parsley and lemon zest. Spread the drained fennel over a shallow ovenproof dish, then add the tomatoes. Drizzle with olive oil, then bake for 10 mins. Nestle the salmon among the veg, sprinkle with lemon juice, then bake 15 mins more until the fish is just cooked. Scatter over the parsley and serve.",
			ImageUrl:      "baked-salmon-with-fennel-tomatoes.jpg",
			Calories:      400, // Estimated per serving
			Protein:       30,  // Estimated per serving
			CookTime:      25,  // Time to bake salmon and vegetables
			PrepTime:      10,  // Time to prepare vegetables and salmon
			TotalTime:     35,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Ovenproof dish, Knife, Saucepan",
			Ingredients:   "2 medium fennel, 2 tbsp chopped parsley, 1 lemon (juiced), 175g cherry tomatoes, 1 tbsp olive oil, 350g salmon, Black olives (optional)",
			TotalEffort:   3,
		},

		{
			Name:        "Blackberry Fool",
			Category:    "Desserts, Comfort Food",
			Servings:    4,
			Description: "A classic British dessert, Blackberry Fool is a light and creamy treat made with fresh blackberries and a mix of cream and yogurt. Served with hazelnut biscuits, this dish is both fruity and indulgent.",
			LightVersionInstructions: sql.NullString{
				String: "Skip making the hazelnut biscuits and serve the fool with store-bought cookies. You can also use pre-made blackberry puree to save time.",
				Valid:  true,
			},
			Instructions:  "For the biscuits, preheat the oven to 200C/180C (fan)/Gas 6 and line two large baking trays with baking parchment. Scatter the hazelnuts over a baking tray and roast for 6-8 minutes, or until golden-brown. Remove and cool. In a bowl, beat butter and sugar until light and creamy. Add chopped nuts, lemon zest, flour, and baking powder, and mix until a dough forms. Shape into 24 balls, flatten them, and bake for 12 minutes. For the fool, cook blackberries with sugar and lemon juice for 15 minutes. Sieve to remove seeds. Whip cream and yogurt until soft peaks form. Gently fold in the blackberry purée. Serve in glass bowls with biscuits.",
			ImageUrl:      "blackberry-fool.jpg",
			Calories:      320, // Estimated per serving
			Protein:       4,   // Estimated per serving
			CookTime:      15,  // Time to prepare the blackberry fool
			PrepTime:      10,  // Time to prepare the biscuits and ingredients
			TotalTime:     25,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Baking trays, Mixing bowl, Saucepan, Electric whisk, Sieve, Glass bowls",
			Ingredients:   "50g hazelnuts, 125g butter, 150g caster sugar, 1 lemon (zested), 150g plain flour, 1/2 tsp baking powder, 600g blackberries, 75g sugar, 2 tbsp caster sugar, 1 tbsp lemon juice, 300ml double cream, 100ml yogurt, Mint leaves for garnish",
			TotalEffort:   2,
		},

		{
			Name:        "Apple Frangipane Tart",
			Category:    "Desserts, Comfort Food",
			Servings:    6,
			Description: "A delicious Apple Frangipane Tart featuring a biscuit base with a rich almond-flavored filling and thinly sliced apples. A perfect dessert for special occasions, best served warm with cream or ice cream.",
			LightVersionInstructions: sql.NullString{
				String: "Use a store-bought tart base to skip the biscuit base preparation and save time. Skip peeling the apples if using softer varieties like Golden Delicious.",
				Valid:  true,
			},
			Instructions:  "Preheat the oven to 200C/180C Fan/Gas 6. Put the digestive biscuits in a resealable freezer bag and crush with a rolling pin into fine crumbs. Melt butter in a pan, mix with biscuit crumbs, and press the mixture evenly over the base and sides of a tart tin. Chill in the fridge while preparing the filling. Cream butter and sugar until light and fluffy, then mix in eggs, ground almonds, and almond extract until combined. Peel and thinly slice the apples. Arrange apple slices over the biscuit base. Spread the frangipane filling over the apples, level the surface, and sprinkle with flaked almonds. Bake for 20-25 minutes until golden-brown and set. Let the tart cool for 15 minutes before removing from the tin. Serve warm with cream, crème fraiche, or ice cream.",
			ImageUrl:      "apple-frangipane-tart.jpg",
			Calories:      450, // Estimated per serving
			Protein:       6,   // Estimated per serving
			CookTime:      25,  // Time to bake the tart
			PrepTime:      15,  // Time to prepare base and filling
			TotalTime:     40,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 4,
			CuttingEffort: 3,
			ItemsRequired: "Tart tin, Rolling pin, Mixing bowl, Wooden spoon, Saucepan, Knife",
			Ingredients:   "175g digestive biscuits, 75g butter, 200g Bramley apples, 75g caster sugar, 75g ground almonds, 1 tsp almond extract, 50g flaked almonds, 2 eggs",
			TotalEffort:   3,
		},

		{
			Name:        "Apple & Blackberry Crumble",
			Category:    "Desserts, Comfort Food",
			Servings:    4,
			Description: "A warm and comforting Apple & Blackberry Crumble, combining tart apples and sweet blackberries with a buttery crumble topping. Perfect for a cozy dessert, served with ice cream.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the pre-baking of the crumble topping and directly bake it with the fruit filling to save time. Use store-bought crumble mix for convenience.",
				Valid:  true,
			},
			Instructions:  "Preheat oven to 190C (170C fan/gas 5). In a bowl, mix the flour and caster sugar. Add butter and rub it into the flour to form a light crumb texture. Spread the mixture on a baking sheet and bake for 15 minutes until lightly golden. For the compote, peel, core, and dice the apples. In a saucepan, melt butter and demerara sugar over medium heat until it turns to light caramel. Add apples and cook for 3 minutes, then add blackberries and cinnamon, cooking for 3 more minutes. Remove from heat and let sit for 2-3 minutes. Spoon the fruit mixture into an ovenproof dish, top with the crumble, and bake for another 5-10 minutes. Serve warm with ice cream.",
			ImageUrl:      "apple-blackberry-crumble.jpg",
			Calories:      380, // Estimated per serving
			Protein:       3,   // Estimated per serving
			CookTime:      15,  // Time to bake the crumble topping and finish dish
			PrepTime:      10,  // Time to prepare fruit and crumble mixture
			TotalTime:     25,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 4,
			CuttingEffort: 3,
			ItemsRequired: "Baking sheet, Ovenproof dish, Saucepan, Mixing bowl, Wooden spoon, Knife",
			Ingredients:   "120g plain flour, 60g caster sugar, 60g butter, 300g Braeburn apples, 30g butter, 30g demerara sugar, 120g blackberries, 1/4 tsp cinnamon, Ice cream for serving",
			TotalEffort:   3,
		},
		{
			Name:        "Spicy Arrabiata Penne",
			Category:    "Main Course, Dinner, Vegetarian",
			Servings:    2,
			Description: "This quick and spicy Arrabiata Penne is perfect for a simple yet flavorful dinner. Made with garlic, red chili flakes, and fresh basil, it brings a classic Italian dish to your table in under 30 minutes.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-minced garlic and canned tomato sauce to speed up the process. Skip the Parmigiano-Reggiano topping for a quicker, dairy-free version.",
				Valid:  true,
			},
			Instructions:  "Bring a large pot of water to a boil. Add kosher salt and the penne pasta. Cook according to package instructions, about 9 minutes. Meanwhile, heat olive oil in a large skillet over medium-high heat until shimmering. Add minced garlic and sauté for 1 to 2 minutes until fragrant. Add chopped tomatoes, red chili flakes, Italian seasoning, and salt and pepper. Bring to a boil and cook for 5 minutes. Stir in chopped basil. Drain the pasta and mix it with the sauce. Garnish with Parmigiano-Reggiano and basil. Serve warm.",
			ImageUrl:      "spicy-arrabiata-penne.jpg",
			Calories:      500, // Estimated per serving
			Protein:       12,  // Estimated per serving
			CookTime:      9,   // Time to cook the pasta
			PrepTime:      10,  // Time for sauce preparation and gathering ingredients
			TotalTime:     19,  // Total time: prep + cook
			WashingEffort: 4,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Pot, Skillet, Strainer, Wooden Spoon, Knife",
			Ingredients:   "1 pound penne rigate, 1/4 cup olive oil, 3 cloves garlic, 1 tin chopped tomatoes, 1/2 teaspoon red chili flakes, 1/2 teaspoon Italian seasoning, 6 basil leaves, Parmigiano-Reggiano flakes",
			TotalEffort:   3,
		},
		{
			Name:        "Chicken in Curry Cream Sauce with Carrot Rice",
			Category:    "Dinner, High-Protein, Main Course",
			Servings:    2,
			Description: "A comforting and protein-rich meal featuring tender chicken in a creamy curry sauce, served with carrot-infused basmati rice. This dish is both flavorful and satisfying.",
			LightVersionInstructions: sql.NullString{
				String: "Use light cream instead of full-fat cream, reduce the amount of butter, and omit the green beans.",
				Valid:  true,
			},
			Instructions:  "Prepare the rice with carrots and cook the chicken until golden brown. Sauté the beans (if using), then prepare the curry cream sauce by combining the sautéed onions, tomato paste, and spices. Simmer the sauce until thickened, then serve over the rice with the chicken. Top with toasted almond flakes.",
			ImageUrl:      "chicken-in-curry-cream-sauce-with-carrot-rice.png",
			Calories:      710,
			Protein:       41,
			CookTime:      25,
			PrepTime:      20,
			TotalTime:     45,
			WashingEffort: 5,
			PeelingEffort: 3,
			CuttingEffort: 4,
			ItemsRequired: "Large pan, small pot, spatula, knife",
			Ingredients:   "250g chicken breast, 150g basmati rice, 1 carrot, 150g green beans, 1 onion, 150g cooking cream, 35g tomato paste, 10g almond flakes, 2g curry spice mix, 4g chicken broth, 400ml water, oil, salt, pepper, butter",
			TotalEffort:   4,
		},
		{
			Name:        "Glazed Mongolian Chicken with Ginger on Sesame Rice",
			Category:    "Dinner, High-Protein, Main Course",
			Servings:    2,
			Description: "A flavorful and protein-packed dish featuring glazed chicken with ginger, served on sesame rice with carrots and spring onions. Perfect for a nutritious and satisfying dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Use less oil and skip the sesame seeds for a lighter version.",
				Valid:  true,
			},
			Instructions:  "Cook the rice by boiling in water and letting it sit. Slice the vegetables and set them aside. Coat the chicken with cornstarch and fry until golden brown. Sauté the vegetables with ginger, then add soy sauce, sesame oil, sesame seeds, and chicken broth. Simmer until the carrots are tender. Mix cornstarch with water, add to the sauce, and cook until thickened. Serve the chicken and sauce over the rice.",
			ImageUrl:      "glazed-mongolian-chicken-with-ginger-on-sesame-rice.png",
			Calories:      664,
			Protein:       41,
			CookTime:      25,
			PrepTime:      25,
			TotalTime:     25,
			WashingEffort: 5,
			PeelingEffort: 3,
			CuttingEffort: 4,
			ItemsRequired: "Small pot, peeler, large bowl, large pan, small bowl",
			Ingredients:   "250g chicken strips, 150g jasmine rice, 2 carrots, 1 spring onion, 2 garlic cloves, 5g ginger paste, 8g cornstarch, 50ml soy sauce, 10g sesame seeds, 10ml sesame oil, 4g chicken broth, salt, sugar, water, oil, pepper",
			TotalEffort:   4,
		},
		{
			Name:        "Morning Whey Protein Coffee Shake",
			Category:    "Breakfast, High-Protein",
			Servings:    1,
			Description: "A quick and energizing shake combining protein and caffeine to kickstart your morning. This shake is perfect for those on the go, providing the necessary fuel to start the day.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the coffee if you prefer a caffeine-free shake.",
				Valid:  true,
			},
			Instructions:  "In a shaker or blender, combine the protein powder, water, and coffee. Mix well until smooth. Enjoy immediately.",
			ImageUrl:      "morning-whey-protein-coffee-shake.png",
			Calories:      150,
			Protein:       30,
			CookTime:      0,
			PrepTime:      2,
			TotalTime:     2,
			WashingEffort: 1,
			PeelingEffort: 0,
			CuttingEffort: 0,
			ItemsRequired: "Shaker or blender",
			Ingredients:   "30g protein powder, 200ml water, 2 cups of coffee",
			TotalEffort:   0,
		},
		{
			Name:        "Berry-Kefir Smoothie",
			Category:    "Breakfast, Healthy, High-Protein",
			Servings:    1,
			Description: "A quick and healthy smoothie that combines the probiotic benefits of kefir with the antioxidants of mixed berries. Ideal for breakfast or a snack.",
			LightVersionInstructions: sql.NullString{
				String: "Use fewer berries and skip the almond butter to reduce calories.",
				Valid:  true,
			},
			Instructions:  "Combine frozen berries, kefir, banana, almond butter, and vanilla extract in a blender. Blend until smooth.",
			ImageUrl:      "berry-kefir-smoothie.png",
			Calories:      304,
			Protein:       15,
			CookTime:      0,
			PrepTime:      5,
			TotalTime:     5,
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Blender",
			Ingredients:   "1 ½ cups frozen mixed berries, 1 cup plain kefir, ½ medium banana, 2 tsp almond butter, ½ tsp vanilla extract",
			TotalEffort:   1,
		},
		{
			Name:        "Healthy Pesto Eggs on Toast",
			Category:    "Breakfast, Vegetarian, Healthy",
			Servings:    2,
			Description: "A quick, low-fat brunch or lunch option where pesto is used to fry eggs, adding delicious flavor to this healthy meal.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the homemade pesto and use store-bought pesto for a faster preparation.",
				Valid:  true,
			},
			Instructions:  "To make the pesto, blend the garlic, basil, pine nuts, oil, and water until smooth. Stir in the cheese. Toast the bread and cook the pesto in a frying pan over medium heat. Fry the eggs and tomatoes in the pesto. Place the eggs on toast. Wilt the spinach in the pan with tomatoes and serve on toast with chili flakes if desired.",
			ImageUrl:      "healthy-pesto-eggs-on-toast.png",
			Calories:      287,
			Protein:       15,
			CookTime:      15,
			PrepTime:      5,
			TotalTime:     20,
			WashingEffort: 4,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Small food processor, frying pan, toaster",
			Ingredients:   "2-4 slices rye sourdough, 2 eggs, 170g tomatoes, 160g baby spinach, pinch of chili flakes, 1 garlic clove, 10g basil, 1 tbsp pine nuts, 1 tbsp rapeseed oil, 1 tbsp finely grated parmesan",
			TotalEffort:   3,
		},

		{
			Name:        "Stuffed Chicken Breast with Feta Cheese on Cannellini Carrot Salad and Lemon Dressing",
			Category:    "Dinner, Low-Carb, Main Course",
			Servings:    2,
			Description: "This delicious low-carb meal features tender chicken breast stuffed with feta cheese, served on a bed of cannellini bean and carrot salad with a refreshing lemon dressing. A satisfying dish packed with flavor but light on carbs, perfect for a healthy dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the step of slicing and marinating the vegetables separately, and simply mix them all together with the dressing after baking. Use pre-crumbled feta cheese to save time.",
				Valid:  true,
			},
			Instructions: `Preheat the oven to 220°C (200°C fan).

Peel the carrots if desired and slice them diagonally into 0.5 cm thick slices. Drain the cannellini beans and rinse them with water. Halve the red onion and cut it into 1 cm wedges. Wash the lemon with hot water and grate 1 teaspoon of the zest. Cut the lemon into 6 wedges.

In a large bowl, mix the rosemary sprig, a third of the "Hello Paprika" spice mix, juice from 2 lemon wedges, lemon zest, 1 tablespoon of olive oil, salt, and pepper. Add the carrots, beans, and onions, mix well, and spread on a baking sheet lined with parchment paper. Roast the vegetables in the oven for 20-25 minutes.

Meanwhile, cut the chicken breast horizontally but not all the way through, opening it like a book. Gently pound the chicken to make it slightly thinner. Season the chicken on all sides with the remaining "Hello Paprika," salt, pepper, and a little olive oil. Crumble half of the feta cheese into the center of the fillets, then fold them back over to close.

Place the chicken fillets on the baking sheet with the vegetables for the last 14-16 minutes of baking time until the chicken is cooked through and no longer pink in the middle.

For the dressing, in the large bowl from step 2, mix mustard, juice from 2 more lemon wedges, 2 tablespoons of olive oil, 1 teaspoon of honey, 1 tablespoon of water, salt, and pepper.

Cut the romaine lettuce into bite-sized pieces. After baking, toss the roasted vegetables in the bowl with the dressing, removing the rosemary sprig. Add the lettuce and mix again.

To serve, arrange the salad on plates, crumble the remaining feta cheese on top, place the chicken breast on top, and drizzle with lemon juice and a little olive oil.

Enjoy your meal!`,
			ImageUrl:      "stuffed-chicken-breast-with-feta-cheese-on-cannellini-carrot-salad-and-lemon-dressing.png",
			Calories:      662,
			Protein:       55,
			CookTime:      20,
			PrepTime:      10,
			TotalTime:     30,
			WashingEffort: 5,
			PeelingEffort: 3,
			CuttingEffort: 5,
			ItemsRequired: "Sieve, Grater, Large Bowl, Frying Pan",
			Ingredients:   "250g chicken breast fillet in brine,100g feta cheese,4g Hello Paprika spice mix,10g rosemary sprig,2 carrots,1 red onion,10ml medium mustard,380g cannellini beans,1 lemon,120g romaine lettuce,1 tsp honey,3 tbsp olive oil,salt,pepper",
			TotalEffort:   4,
		},

		{
			Name:        "Beef and Broccoli Rice Skillet with Sweet Sauce and Sesame Topping",
			Category:    "Main Course, One-Pot Dish, Dairy-Free",
			Servings:    2,
			Description: "This Beef and Broccoli Rice Skillet is a flavorful one-pot dish combining tender ground beef, fresh broccoli, and aromatic basmati rice, all brought together with a sweet soy sauce and topped with toasted sesame seeds. It's a perfect meal for a quick and easy dinner with minimal cleanup.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the step of toasting the sesame seeds and use pre-made garlic and ginger paste. Reduce the simmering time by cooking the broccoli separately in the microwave and adding it at the end.",
				Valid:  true,
			},
			Instructions: `Heat 350 ml of water in a kettle. Cut the broccoli into bite-sized florets. Halve the onion and slice it into strips.
In a large pan with a lid, toast the sesame seeds for 1-2 minutes without any oil until they are browned. Remove and set aside.
In the same pan, crumble and brown the ground beef for 2-3 minutes. Add cornstarch, garlic & ginger mix, soy sauce, 2 teaspoons of sugar, 1 teaspoon of balsamic vinegar, and pepper. Let it thicken for 1 minute, then transfer to a bowl and set aside.
In the pan, heat 1 tablespoon of oil. Sauté the broccoli and onion for 2-3 minutes. Add the rice, beef broth, 0.25 teaspoon of salt, and 350 ml of water. Stir well, bring to a boil, then cover and simmer on medium heat for 12 minutes.
Stir the rice, add the beef with the sauce on top, remove the pan from heat, and let it sit covered for at least 10 minutes. Season with salt and pepper if needed.
Serve the Beef & Broccoli Rice Skillet on deep plates, garnished with toasted sesame seeds. Enjoy your meal!`,
			ImageUrl:      "beef-and-broccoli-rice-skillet.png",
			Calories:      737,
			Protein:       34,
			CookTime:      20,
			PrepTime:      25,
			TotalTime:     45,
			WashingEffort: 4,
			PeelingEffort: 3,
			CuttingEffort: 4,
			ItemsRequired: "Soup pan or large pan with lid, Small Bowl",
			Ingredients:   "150g basmati rice,200g ground beef,1 broccoli,1 onion,50ml soy sauce,10g sesame seeds,30g chopped garlic & ginger in oil,4g beef broth,4g cornstarch,1/4 teaspoon salt,350ml water,1 teaspoon balsamic vinegar,2 teaspoons sugar,1 tablespoon oil,pepper to taste",
		},

		{
			Name:        "Spicy Fusilli Calabrese with Spinach and Pine Nuts",
			Category:    "Lunch, Dinner, Vegetarian",
			Servings:    2,
			Description: "This spicy Fusilli Calabrese dish with spinach and pine nuts brings the flavors of Italy to your home in just 30 minutes. The combination of creamy cheese, fresh spinach, and crunchy pine nuts with a hint of chili makes this a perfect quick meal.",
			LightVersionInstructions: sql.NullString{
				String: "Skip roasting the pine nuts to save time. Use pre-minced garlic instead of fresh. You can also skip the fresh chili and use only chili flakes for less preparation.",
				Valid:  true,
			},
			Instructions: `Heat a large amount of water in a kettle. Peel and finely chop the onion. Peel the garlic.
            Roast pine nuts in a large pot without adding any fat for 30 seconds to 1 minute, stirring until they brown and become fragrant. Remove and set aside.
            In the same pot, heat 1 tablespoon of olive oil. Add the onions and press in the garlic, sauté both for 1 to 2 minutes over medium heat until translucent.
            Add 500 ml of hot water, Fusilli, vegetable broth powder, and 0.5 teaspoon of salt to the pot. Bring to a boil and cover, simmer for about 9 to 11 minutes until the pasta is al dente. Stir several times, especially towards the end when the water has almost evaporated, to prevent the pasta from sticking.
            Meanwhile, halve the chili lengthwise, deseed it, and cut it into thin strips (be careful: it's hot!). Halve the cherry tomatoes. When the pasta is al dente, stir in cream cheese, Calabrese pesto, and chili flakes to taste (be careful: it's hot!), creating a smooth sauce. Stir in the cherry tomatoes and baby spinach, heating for 1 to 2 minutes. Season with salt and pepper to taste.
            Distribute the one-pot pasta onto plates. Top with roasted pine nuts, grated hard cheese, and optionally, chili strips (be careful: it's hot!). Enjoy your meal!`,
			ImageUrl:      "spicy-fusilli-calabrese-with-spinach-and-pine-nuts.png",
			Calories:      852,
			Protein:       32,
			CookTime:      11,
			PrepTime:      19,
			TotalTime:     30,
			WashingEffort: 6,
			PeelingEffort: 3,
			CuttingEffort: 5,
			ItemsRequired: "Large pot, Kettle, Bowl",
			Ingredients:   "1 Onion, 1 Garlic clove, 10g Pine nuts, 4g Vegetable broth, 40g Grated hard cheese, 270g Fusilli pasta, 100g Baby spinach, 100g Cream cheese, 50g Calabrese pesto, 125g Cherry tomatoes, 1 Red chili, 1g Mild chili flakes, 500ml Water, 1 tablespoon Olive oil, 0.5 teaspoon Salt, Pepper to taste",
			TotalEffort:   5,
		},
		{
			Name:        "Tex-Mex Chili Bowl with Tortilla Chips, Lime Crema, and Tomato Salsa",
			Category:    "Main Course, Comfort Food, Family",
			Servings:    2,
			Description: "This Tex-Mex Chili Bowl is a quick and delicious meal featuring ground beef, black beans, and tomatoes, topped with crispy tortilla chips, lime crema, and fresh tomato salsa. Perfect for a comforting yet fresh dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the step of making the lime crema and serve with plain sour cream. Use store-bought salsa instead of making it fresh. Reduce the garnish to save time.",
				Valid:  true,
			},
			Instructions: `1. Slice the white and green parts of the spring onion separately into thin rings. Quarter the lime. Drain the black beans in a sieve.
In a large pan, heat 1 tablespoon of oil. Press the garlic into the pan and sauté for 1 minute. Remove and mix with sour cream and juice from 1 lime quarter. Season with salt and pepper.
2. In the same pan, without adding extra oil, brown the ground beef and the white part of the spring onion for 2-3 minutes. Add the "Hello Fiesta" spice mix and black beans, cooking for another minute. Deglaze with diced tomatoes, 50 ml of water, and 1 teaspoon of sugar. Simmer for 3-4 minutes, then season with salt and pepper.
3. Serve the chili in deep bowls, topped with tortilla chips, lime crema, and tomato salsa. Garnish with the green part of the spring onion and serve with the remaining lime wedges. Enjoy!`,
			ImageUrl:      "tex-mex-chili-bowl-with-tortilla-chips.png",
			Calories:      855,
			Protein:       40,
			CookTime:      10,
			PrepTime:      5,
			TotalTime:     15,
			WashingEffort: 5,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Sieve, Large Pan, Small Bowl, Large Frying Pan",
			Ingredients:   "380g black beans,75g tortilla chips,100g tomato salsa,1 garlic clove,6g 'Hello Fiesta' spice mix,100g sour cream,390g diced tomatoes,200g ground beef,1 spring onion,1 lime,50ml water,salt to taste,pepper to taste,1.5 tablespoons oil",
		},
		{
			Name:        "Chicken Handi",
			Category:    "Main Course, Dinner",
			Servings:    4,
			Description: "A flavorful and creamy Chicken Handi cooked with aromatic spices like cumin, coriander, and garam masala, blended with yogurt and cream. This popular dish is perfect when served with naan or roti.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-fried onions and canned tomato sauce to save time. You can also use boneless chicken pieces for quicker cooking.",
				Valid:  true,
			},
			Instructions:  "Heat oil in a large pot, fry onions until golden brown, and set aside. Sauté garlic, add tomatoes, and cook until soft. Add the fried onions, ginger, cumin, coriander, green chilies, turmeric, and red chili powder. Add chicken and cook for 15 minutes. Add yogurt, fenugreek, and cream. Finish with garam masala and fenugreek leaves, and serve hot with naan or roti.",
			ImageUrl:      "chicken-handi.jpg",
			Calories:      450, // Estimated per serving
			Protein:       40,  // Estimated per serving
			CookTime:      30,  // Time to cook the chicken
			PrepTime:      15,  // Time to prepare ingredients
			TotalTime:     45,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Large pot, Knife, Wooden spoon",
			Ingredients:   "1.2 kg chicken, 5 onions, 2 tomatoes, 8 garlic cloves, 1 tbsp ginger paste, 1/4 cup vegetable oil, 2 tsp cumin seeds, 3 tsp coriander seeds, 1 tsp turmeric, 1 tsp chili powder, 2 green chilies, 1 cup yogurt, 3/4 cup cream, 3 tsp dried fenugreek, 1 tsp garam masala",
			TotalEffort:   3,
		},
		{
			Name:        "Croatian Bean Stew",
			Category:    "Soups, Main Course, Comfort Food",
			Servings:    4,
			Description: "A hearty and flavorful Croatian Bean Stew made with cannellini beans, vegetables, and chorizo. This stew is perfect for a cozy meal and gets even better when reheated the next day.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-cooked beans and pre-chopped vegetables to save time. Skip the oven and finish on the stovetop for faster cooking.",
				Valid:  true,
			},
			Instructions:  "Heat oil in a pan, sauté the chopped vegetables until tender. Add beans and vegetables into a pot, along with chorizo. Cook for 20 minutes on low heat or bake at 180C/350F for 30 minutes. Serve warm.",
			ImageUrl:      "croatian-bean-stew.jpg",
			Calories:      500, // Estimated per serving
			Protein:       20,  // Estimated per serving
			CookTime:      30,  // Time to cook the stew
			PrepTime:      15,  // Time to prepare ingredients
			TotalTime:     45,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Pan, Pot, Ovenproof dish",
			Ingredients:   "2 cans cannellini beans, 3 tbsp vegetable oil, 2 cups tomatoes, 5 shallots, 2 garlic cloves, 1/2 kg chorizo, Pinch parsley",
			TotalEffort:   3,
		},
		{
			Name:        "Chivito uruguayo",
			Category:    "Main Course, Snacks",
			Servings:    1,
			Description: "Chivito uruguayo is a traditional Uruguayan sandwich made with beef, bacon, ham, mozzarella, fried egg, and fresh vegetables. It’s a delicious, hearty sandwich perfect for a quick and filling meal.",
			LightVersionInstructions: sql.NullString{
				String: "Skip frying the bacon and ham if you want to reduce time and effort. You can also use pre-cooked or thinly sliced beef.",
				Valid:  true,
			},
			Instructions:  "Flatten the beef and cook it on a griddle. Fry the bacon, ham, and eggs. Assemble the sandwich by layering the beef, bacon, ham, egg, mozzarella, tomato, and lettuce on bread. Serve immediately.",
			ImageUrl:      "chivito-uruguayo.jpg",
			Calories:      700, // Estimated per serving
			Protein:       40,  // Estimated per serving
			CookTime:      15,  // Time to cook the meat and eggs
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     25,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Griddle, Knife, Frying pan",
			Ingredients:   "2 beef brisket slices, 2 bread rolls, 1 tomato, 1 lettuce, 100g ham, 100g mozzarella, 100g bacon, 1 egg, 1 onion, 1 pepper",
			TotalEffort:   2,
		},
		{
			Name:        "Fettucine Alfredo",
			Category:    "Main Course, Dinner",
			Servings:    2,
			Description: "A rich and creamy Fettucine Alfredo made with clotted cream, Parmesan cheese, and butter. This quick and easy pasta dish is perfect for a comforting meal, topped with freshly chopped parsley for added flavor.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought Alfredo sauce to save time, or skip the nutmeg for a simpler flavor profile.",
				Valid:  true,
			},
			Instructions:  "In a medium saucepan, stir the clotted cream, butter, and cornflour over low heat until it simmers. Cook the pasta in salted boiling water, then drain and mix with the cream sauce. Add Parmesan and stir over low heat until the sauce is glossy. Serve with parsley.",
			ImageUrl:      "fettucine-alfredo.jpg",
			Calories:      650, // Estimated per serving
			Protein:       15,  // Estimated per serving
			CookTime:      10,  // Time to cook pasta and sauce
			PrepTime:      5,   // Time to prepare ingredients
			TotalTime:     15,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Saucepan, Rubber spatula, Pot",
			Ingredients:   "227g clotted cream, 25g butter, 1 tsp cornflour, 100g Parmesan, Nutmeg, 250g fettucine, Parsley",
			TotalEffort:   2,
		},
		{
			Name:        "Grilled Mac and Cheese Sandwich",
			Category:    "Main Course, Snacks, Comfort Food",
			Servings:    4,
			Description: "A comforting and indulgent Grilled Mac and Cheese Sandwich, combining creamy macaroni and cheese with toasted bread for a satisfying meal or snack. Perfect for cheese lovers!",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-made mac and cheese and skip refrigerating it to save time.",
				Valid:  true,
			},
			Instructions:  "Prepare mac and cheese by cooking pasta, then making a cheese sauce with butter, flour, milk, cream, and cheese. Chill the mac and cheese, then assemble sandwiches with cheddar cheese and garlic butter. Cook on a skillet until golden brown.",
			ImageUrl:      "grilled-mac-and-cheese-sandwich.jpg",
			Calories:      700, // Estimated per serving
			Protein:       25,  // Estimated per serving
			CookTime:      20,  // Time to cook the sandwiches
			PrepTime:      30,  // Time to prepare mac and cheese
			TotalTime:     50,  // Total time: prep + cook
			WashingEffort: 4,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Skillet, Spatula, Saucepan, Baking sheet",
			Ingredients:   "230g macaroni, 1/3 cup flour, 1 1/2 cups milk, 1 cup heavy cream, 455g Monterey Jack cheese, 16 slices bread, 8 slices cheddar cheese, 6 tbsp butter, 1 tsp garlic powder",
			TotalEffort:   3,
		},
		{
			Name:        "General Tso's Chicken",
			Category:    "Main Course, Dinner",
			Servings:    4,
			Description: "General Tso's Chicken is a popular Chinese-American dish featuring crispy fried chicken in a sweet, tangy sauce. It's flavorful and perfect for dinner served over rice.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the deep frying and stir-fry the chicken instead to reduce oil usage.",
				Valid:  true,
			},
			Instructions:  "Mix water, soy sauce, vinegar, garlic, ginger, and sauces for the sauce. Marinate chicken in egg whites and cornstarch, then deep fry until golden. Stir-fry the sauce, add broccoli, thicken with cornstarch, and combine with chicken. Serve hot over rice.",
			ImageUrl:      "general-tsos-chicken.jpg",
			Calories:      550, // Estimated per serving
			Protein:       35,  // Estimated per serving
			CookTime:      30,  // Time to fry and stir-fry the chicken
			PrepTime:      15,  // Time to marinate and prepare ingredients
			TotalTime:     45,  // Total time: prep + cook
			WashingEffort: 4,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Wok, Deep fryer, Knife",
			Ingredients:   "2 cups water, 2 tbsp soy sauce, 2 tbsp rice vinegar, 2 egg whites, 3 tbsp cornstarch, 2 chicken breasts, 1/4 cup hoisin sauce, 1 tbsp ketchup, 1 tsp ginger, 2 garlic cloves, 1 tsp sugar",
			TotalEffort:   3,
		},
		{
			Name:        "Honey Teriyaki Salmon",
			Category:    "Main Course, Dinner, Healthy",
			Servings:    2,
			Description: "A delicious and simple Honey Teriyaki Salmon with a rich glaze, perfect for a quick and healthy dinner. Pan-fried to perfection, this dish is garnished with sesame seeds and pairs wonderfully with steamed vegetables or rice.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought teriyaki sauce to save time on preparing the glaze.",
				Valid:  true,
			},
			Instructions:  "Mix all ingredients for the honey teriyaki glaze. Coat the salmon with the glaze. Heat oil in a skillet over medium heat and pan-fry the salmon on both sides until cooked through and the glaze thickens. Garnish with sesame seeds and serve.",
			ImageUrl:      "honey-teriyaki-salmon.jpg",
			Calories:      400, // Estimated per serving
			Protein:       35,  // Estimated per serving
			CookTime:      10,  // Time to pan-fry the salmon
			PrepTime:      5,   // Time to prepare the glaze
			TotalTime:     15,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Skillet, Whisk",
			Ingredients:   "1 lb salmon, 1 tbsp olive oil, 2 tbsp soy sauce, 2 tbsp sake, 4 tbsp sesame seeds",
			TotalEffort:   2,
		},
		{
			Name:        "Honey Balsamic Chicken with Crispy Broccoli & Potatoes",
			Category:    "Main Course, Dinner, Healthy",
			Servings:    2,
			Description: "A healthy and flavorful Honey Balsamic Chicken with crispy roasted broccoli and potatoes. This dish combines tender chicken with a sweet balsamic glaze and crispy vegetables for a well-balanced meal.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the glaze and use a store-bought balsamic vinaigrette to save time.",
				Valid:  true,
			},
			Instructions:  "Preheat oven to 425°F. Toss potatoes with oil, salt, and pepper, roast for 5 minutes. Add broccoli with garlic oil and roast until crispy. Meanwhile, cook chicken in a pan until browned, set aside. Make a balsamic glaze with vinegar, honey, and stock, and coat the chicken. Serve with crispy potatoes and broccoli.",
			ImageUrl:      "honey-balsamic-chicken-with-crispy-broccoli-potatoes.jpg",
			Calories:      600, // Estimated per serving
			Protein:       35,  // Estimated per serving
			CookTime:      30,  // Time to roast veggies and cook chicken
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     40,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 2,
			ItemsRequired: "Baking sheet, Pan, Knife",
			Ingredients:   "5 potatoes, 1 broccoli, 2 garlic cloves, 2 chicken breasts, Balsamic vinegar, Honey, Chicken stock, 1 tbsp butter, 1 tbsp vegetable oil, 1 tbsp olive oil",
			TotalEffort:   3,
		},
		{
			Name:        "Key Lime Pie",
			Category:    "Desserts, Comfort Food",
			Servings:    8,
			Description: "A classic Key Lime Pie with a crunchy biscuit base, filled with a zesty and creamy lime filling. It's topped with softly whipped cream, making it the perfect refreshing dessert.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought graham cracker crust and ready-made whipped cream to save time.",
				Valid:  true,
			},
			Instructions:  "Crush digestive biscuits and mix with melted butter, then press into a tart tin. Bake the base and let it cool. Whisk egg yolks, condensed milk, lime zest, and juice, then pour into the base. Bake for 15 minutes and chill for 3 hours. Serve topped with whipped cream and lime zest.",
			ImageUrl:      "key-lime-pie.jpg",
			Calories:      450, // Estimated per serving
			Protein:       8,   // Estimated per serving
			CookTime:      25,  // Total baking time
			PrepTime:      15,  // Time to prepare the base and filling
			TotalTime:     40,  // Total time: prep + bake + chill
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Tart tin, Food processor, Whisk, Electric beaters",
			Ingredients:   "300g digestive biscuits, 150g butter, 400g condensed milk, 3 egg yolks, 4 limes, 300ml double cream, 1 tbsp icing sugar",
			TotalEffort:   2,
		},
		{
			Name:        "Kung Pao Chicken",
			Category:    "Main Course, Dinner",
			Servings:    4,
			Description: "A spicy and flavorful Kung Pao Chicken dish made with marinated chicken, peanuts, and water chestnuts, tossed in a rich soy-based sauce. This popular Chinese dish is perfect for a quick and satisfying dinner.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought Kung Pao sauce to save time on preparing the sauce from scratch.",
				Valid:  true,
			},
			Instructions:  "Marinate the chicken with sake, soy sauce, sesame oil, and cornflour for 30 minutes. Prepare the sauce with the remaining marinade, chillies, vinegar, and sugar. Sauté the chicken and combine with the sauce, spring onions, garlic, water chestnuts, and peanuts. Simmer until thickened.",
			ImageUrl:      "kung-pao-chicken.jpg",
			Calories:      550, // Estimated per serving
			Protein:       35,  // Estimated per serving
			CookTime:      20,  // Time to sauté and cook chicken
			PrepTime:      30,  // Time for marinating and preparing ingredients
			TotalTime:     50,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Frying pan, Knife, Glass bowl",
			Ingredients:   "500g chicken, 2 tbsp sake, 2 tbsp soy sauce, 2 tbsp sesame oil, 2 tbsp cornflour, 1 tsp rice vinegar, 1 tbsp brown sugar, 4 spring onions, 6 garlic cloves, 220g water chestnuts, 100g peanuts",
			TotalEffort:   3,
		},
		{
			Name:        "Lasagne",
			Category:    "Main Course, Dinner, Comfort Food",
			Servings:    4,
			Description: "A classic lasagne layered with rich beef ragu, creamy crème fraîche, and mozzarella, topped with Parmesan for a perfect golden finish. This hearty dish is ideal for a comforting family meal.",
			LightVersionInstructions: sql.NullString{
				String: "Use store-bought lasagne sheets and ready-made ragu sauce to save time.",
				Valid:  true,
			},
			Instructions:  "Cook bacon, onions, celery, and carrots until softened. Add minced beef, cook until browned, then stir in tomato puree, chopped tomatoes, honey, and seasoning. Simmer for 20 minutes. Layer the ragu sauce and lasagne sheets in a baking dish, finish with crème fraîche and mozzarella topping. Bake for 25-30 minutes until golden and bubbling.",
			ImageUrl:      "lasagne.jpg",
			Calories:      750, // Estimated per serving
			Protein:       35,  // Estimated per serving
			CookTime:      30,  // Time to bake the lasagne
			PrepTime:      20,  // Time to prepare the ragu sauce
			TotalTime:     50,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Saucepan, Baking dish, Knife",
			Ingredients:   "500g minced beef, 2 bacon rashers, 1 onion, 1 carrot, 1 stick celery, 2 cloves garlic, 1 tbsp tomato puree, 800g chopped tomatoes, 1 tbsp honey, 400ml crème fraîche, 125g mozzarella, 50g Parmesan, 500g lasagne sheets",
			TotalEffort:   3,
		},
		{
			Name:        "Moroccan Carrot Soup",
			Category:    "Soups, Healthy",
			Servings:    4,
			Description: "A fragrant and warming Moroccan Carrot Soup made with roasted carrots, cumin, and coriander. This soup is both light and flavorful, perfect for a healthy meal.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the roasting step and boil the vegetables instead to reduce time.",
				Valid:  true,
			},
			Instructions:  "Preheat oven to 180°C. Mix carrots, onion, garlic, cumin, coriander, salt, and olive oil. Roast for 10-12 minutes. Blend the roasted vegetables with water into a smooth paste. Reheat in a pan, add garam masala and lemon juice. Serve hot.",
			ImageUrl:      "moroccan-carrot-soup.jpg",
			Calories:      150, // Estimated per serving
			Protein:       2,   // Estimated per serving
			CookTime:      12,  // Time to roast vegetables
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     22,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Baking tray, Blender, Non-stick pan",
			Ingredients:   "6 carrots, 1 onion, 4 garlic cloves, 1 tsp cumin, 1/2 tsp coriander, 1 tbsp olive oil, 1/4 tsp garam masala, 1 tsp lemon juice",
			TotalEffort:   2,
		},
		{
			Name:        "Pancakes",
			Category:    "Breakfast, Desserts",
			Servings:    4,
			Description: "Classic fluffy pancakes made with simple ingredients like flour, eggs, and milk. These are perfect for breakfast, served with fruit or syrup.",
			LightVersionInstructions: sql.NullString{
				String: "Skip the resting time for the batter and use pre-made pancake mix to save time.",
				Valid:  true,
			},
			Instructions:  "Whisk flour, eggs, milk, and oil into a smooth batter. Cook pancakes in a pan for 1 minute on each side until golden. Serve with lemon wedges, sugar, or your favorite topping.",
			ImageUrl:      "pancakes.jpg",
			Calories:      200, // Estimated per serving
			Protein:       6,   // Estimated per serving
			CookTime:      10,  // Time to cook the pancakes
			PrepTime:      5,   // Time to prepare batter
			TotalTime:     15,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 1,
			ItemsRequired: "Mixing bowl, Frying pan, Whisk",
			Ingredients:   "100g flour, 2 eggs, 300ml milk, 1 tbsp oil, Sugar for serving, Raspberries, Blueberries",
			TotalEffort:   2,
		},
		{
			Name:        "Stuffed Bell Peppers with Quinoa and Black Beans",
			Category:    "Main Course, Vegetarian",
			Servings:    4,
			Description: "These healthy stuffed bell peppers are filled with a flavorful mixture of quinoa, black beans, corn, and spices. Topped with melted cheese, they make for a satisfying vegetarian meal.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-cooked quinoa and canned beans to save time.",
				Valid:  true,
			},
			Instructions:  "Preheat oven to 375°F. Bake bell pepper halves for 15-20 minutes. Sauté onion and garlic, add quinoa, beans, corn, tomatoes, and spices. Stuff the peppers, top with cheese, and bake for another 15-20 minutes. Garnish with cilantro and serve.",
			ImageUrl:      "stuffed-bell-peppers-with-quinoa-and-black-beans.jpg",
			Calories:      350, // Estimated per serving
			Protein:       10,  // Estimated per serving
			CookTime:      40,  // Time to bake peppers and cook the filling
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     50,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 1,
			CuttingEffort: 3,
			ItemsRequired: "Baking dish, Skillet, Knife",
			Ingredients:   "4 bell peppers, 1 cup quinoa, 1 can black beans, 1 cup corn, 1 can diced tomatoes, 1 onion, 2 garlic cloves, 1 tsp cumin, 1/2 tsp chili powder, 1/2 tsp smoked paprika, 1 1/2 cup shredded cheese",
			TotalEffort:   3,
		},
		{
			Name:        "Shakshuka",
			Category:    "Breakfast, Vegetarian",
			Servings:    4,
			Description: "Shakshuka is a flavorful and comforting dish made of poached eggs in a rich tomato sauce, seasoned with spices and fresh herbs. It’s perfect for breakfast or brunch, served with crusty bread.",
			LightVersionInstructions: sql.NullString{
				String: "Use canned tomatoes and pre-chopped vegetables to save time.",
				Valid:  true,
			},
			Instructions:  "Heat oil in a pan, soften onions, chili, garlic, and coriander for 5 mins. Add tomatoes and sugar, simmer for 8-10 mins. Create dips in the sauce, crack eggs, and cook with the lid on for 6-8 mins. Garnish with coriander and serve.",
			ImageUrl:      "shakshuka.jpg",
			Calories:      300, // Estimated per serving
			Protein:       12,  // Estimated per serving
			CookTime:      15,  // Time to cook eggs and sauce
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     25,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Frying pan, Knife",
			Ingredients:   "1 tbsp olive oil, 2 onions, 1 red chili, 1 garlic clove, 1 tbsp caster sugar, 800g cherry tomatoes, 4 eggs, 1 tbsp chopped coriander, Feta cheese",
			TotalEffort:   2,
		},
		{
			Name:        "Tandoori Chicken",
			Category:    "Main Course, Dinner",
			Servings:    4,
			Description: "Tandoori Chicken is a popular Indian dish where chicken thighs are marinated in yogurt and spices, then grilled to perfection for a smoky, charred flavor. It’s perfect for a flavorful and healthy meal.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-marinated chicken thighs to save time.",
				Valid:  true,
			},
			Instructions:  "Mix lemon juice with paprika and onions, slash chicken thighs, and marinate. Prepare the yogurt marinade with ginger, garlic, cumin, garam masala, chili powder, and turmeric. Let marinate for at least 1 hour, then grill for 8 minutes per side until cooked through.",
			ImageUrl:      "tandoori-chicken.jpg",
			Calories:      400, // Estimated per serving
			Protein:       35,  // Estimated per serving
			CookTime:      16,  // Time to grill the chicken
			PrepTime:      10,  // Time to prepare the marinade
			TotalTime:     26,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 2,
			ItemsRequired: "Grill, Mixing bowl, Knife",
			Ingredients:   "16 chicken thighs, 2 lemons, 4 tsp paprika, 2 onions, 300ml yogurt, 4 garlic cloves, 1 tbsp ginger, 1 tsp garam masala, 1 tsp cumin, 1/2 tsp chili powder, 1/4 tsp turmeric",
			TotalEffort:   3,
		},
		{
			Name:        "Vegan Lasagna",
			Category:    "Main Course, Vegan",
			Servings:    4,
			Description: "A healthy and flavorful Vegan Lasagna made with lentils, vegetables, and a creamy vegan béchamel sauce. Perfect for a hearty and nutritious meal.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-cooked lentils and ready-made vegan béchamel sauce to save time.",
				Valid:  true,
			},
			Instructions:  "Boil vegetables and lentils for 20 minutes. Blanch spinach and cook lasagna sheets. Make the béchamel by combining vegan butter, flour, soya milk, mustard, and vinegar. Assemble the lasagna in a dish and bake at 180°C for 25 minutes.",
			ImageUrl:      "vegan-lasagna.jpg",
			Calories:      400, // Estimated per serving
			Protein:       12,  // Estimated per serving
			CookTime:      25,  // Time to bake lasagna
			PrepTime:      20,  // Time to prepare ingredients
			TotalTime:     45,  // Total time: prep + cook
			WashingEffort: 3,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Baking dish, Pan, Mixing bowl",
			Ingredients:   "1 cup red lentils, 1 carrot, 1 onion, 1 zucchini, 150g spinach, 10 lasagna sheets, 35g vegan butter, 4 tbsp flour, 300ml soya milk, 1.5 tsp mustard, 1 tsp vinegar",
			TotalEffort:   3,
		},
		{
			Name:        "Vegetarian Chilli",
			Category:    "Main Course, Vegetarian",
			Servings:    4,
			Description: "A hearty and healthy vegetarian chilli made with roasted vegetables, kidney beans, and mixed grains. It's a perfect, easy-to-make dish packed with flavor.",
			LightVersionInstructions: sql.NullString{
				String: "Use pre-cooked grains and canned vegetables to save time.",
				Valid:  true,
			},
			Instructions:  "Preheat oven to 200C. Roast vegetables for 15 minutes. Add beans and tomatoes, season, and cook for another 10-15 minutes. Heat the grains in the microwave and serve with the chilli.",
			ImageUrl:      "vegetarian-chilli.jpg",
			Calories:      350, // Estimated per serving
			Protein:       12,  // Estimated per serving
			CookTime:      30,  // Total cooking time
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     40,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 1,
			CuttingEffort: 2,
			ItemsRequired: "Casserole dish, Knife, Microwave",
			Ingredients:   "400g roasted vegetables, 1 can kidney beans, 1 can chopped tomatoes, 1 packet mixed grains",
			TotalEffort:   2,
		},
		{
			Name:        "Vegetarian Casserole",
			Category:    "Main Course, Vegetarian",
			Servings:    4,
			Description: "A healthy and comforting vegetarian casserole made with lentils, fresh vegetables, and herbs. This dish is perfect for a wholesome meal, served with rice or quinoa.",
			LightVersionInstructions: sql.NullString{
				String: "Use canned lentils and pre-cut vegetables to save time.",
				Valid:  true,
			},
			Instructions:  "Heat oil in a pan, cook onions until softened. Add garlic, spices, and vegetables. Cook for 5 minutes. Add tomatoes, stock, and thyme, simmer for 20-25 minutes. Stir in lentils and serve with rice or quinoa.",
			ImageUrl:      "vegetarian-casserole.jpg",
			Calories:      350, // Estimated per serving
			Protein:       15,  // Estimated per serving
			CookTime:      25,  // Time to cook casserole
			PrepTime:      10,  // Time to prepare ingredients
			TotalTime:     35,  // Total time: prep + cook
			WashingEffort: 2,
			PeelingEffort: 2,
			CuttingEffort: 3,
			ItemsRequired: "Pan, Knife, Spoon",
			Ingredients:   "1 tbsp rapeseed oil, 1 onion, 3 cloves garlic, 1 tsp paprika, 1/2 tsp cumin, 1 tbsp thyme, 3 carrots, 2 celery stalks, 1 red pepper, 1 yellow pepper, 2 tins tomatoes, 250ml vegetable stock, 2 courgettes, 250g lentils",
			TotalEffort:   3,
		},
	}

}
