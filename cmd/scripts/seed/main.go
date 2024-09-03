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
			Instructions: `Slice the white and green parts of the spring onion separately into thin rings. Quarter the lime. Drain the black beans in a sieve.
In a large pan, heat 1 tablespoon of oil. Press the garlic into the pan and sauté for 1 minute. Remove and mix with sour cream and juice from 1 lime quarter. Season with salt and pepper.
In the same pan, without adding extra oil, brown the ground beef and the white part of the spring onion for 2-3 minutes. Add the "Hello Fiesta" spice mix and black beans, cooking for another minute. Deglaze with diced tomatoes, 50 ml of water, and 1 teaspoon of sugar. Simmer for 3-4 minutes, then season with salt and pepper.
Serve the chili in deep bowls, topped with tortilla chips, lime crema, and tomato salsa. Garnish with the green part of the spring onion and serve with the remaining lime wedges. Enjoy!`,
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
