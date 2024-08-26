package main

import (
	"context"
	"fmt"
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"
	"time"
)

func main() {
	ctx := context.Background()
	d := db.Get()

	mealsParams := []sqlc.CreateMealParams{
		{
			Name:             "Spaghetti Carbonara",
			Category:         "Lunch,Dinner",
			Description:      "Spaghetti Carbonara is a classic pasta dish that everyone should know how to make. It's a simple dish that uses few ingredients, but they're all high quality, so it's important to use the best you can find.",
			ShortDescription: "Spaghetti Carbonara is a classic pasta dish that everyone should know how to make.",
			ImageUrl:         "https://www.themealdb.com/images/media/meals/xxpqsy1511452222.jpg",
			// base64
			ImagePreview: "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxMTEhUTExMWFhUXGBgYGBgYGBgYGBgYGBgYFxgYFxgYHSggGBolHRgYITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGxAQGy0lICUt",
			Calories:     500,
			Likes:        0,
			// effort between 1-5, consider the washing, cutting, and peeling effort of ingredients
			TotalEffort:      3,
			InstructionSteps: "1. Bring a large pot of salted water to a boil. Add spaghetti and cook until al dente, 8 to 10 minutes. Drain and set aside.\n2. In a large skillet, cook pancetta until crispy. Remove pancetta and all but 1 tablespoon of fat. Add garlic and cook until fragrant, 1 minute. Add cooked spaghetti and toss until fully coated in pancetta fat.\n3. Remove skillet from heat and add eggs and Parmesan. Toss until creamy, then season generously with black pepper. Serve immediately with more cheese.",
			PrepTime:         int64(time.Duration(10) * time.Minute),
			CookTime:         int64(time.Duration(20) * time.Minute),
			TotalTime:        int64(time.Duration(30) * time.Minute),
		},
		{
			Name:             "Chicken Alfredo",
			Category:         "Lunch,Dinner",
			Description:      "Chicken Alfredo is a classic Italian-American dish that's perfect for a cozy night in. It's a simple dish that combines tender chicken, creamy Alfredo sauce, and perfectly cooked pasta.",
			ShortDescription: "Chicken Alfredo is a classic Italian-American dish that's perfect for a cozy night in.",
			ImageUrl:         "https://www.themealdb.com/images/media/meals/syqypv1486981727.jpg",
			// base64
			ImagePreview: "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxMTEhUTExMWFhUXGBgYGBgYGBgYGBgYGBgYFxgYFxgYHSggGBolHRgYITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGxAQGy0lICUt",
			Calories:     700,
			Likes:        0,
			// effort between 1-5, consider the washing, cutting, and peeling effort of ingredients
			TotalEffort:      4,
			InstructionSteps: "1. Bring a large pot of salted water to a boil. Add fettuccine and cook until al dente, 8 to 10 minutes. Drain and set aside.\n2. In a large skillet, cook chicken until golden and no longer pink, 8 minutes per side. Remove chicken and set aside. Add garlic and cook until fragrant, 1 minute. Add heavy cream and bring to a simmer.\n3. Add cooked fettuccine and toss until fully coated in sauce. Add Parmesan and toss until creamy. Season with salt and pepper. Serve immediately with more cheese.",
			PrepTime:         int64(time.Duration(15) * time.Minute),
			CookTime:         int64(time.Duration(25) * time.Minute),
			TotalTime:        int64(time.Duration(40) * time.Minute),
		},
		{
			Name:             "Chicken Parmesan",
			Category:         "Lunch,Dinner",
			Description:      "Chicken Parmesan is a classic Italian-American dish that's perfect for a cozy night in. It's a simple dish that combines tender chicken, marinara sauce, and melted cheese.",
			ShortDescription: "Chicken Parmesan is a classic Italian-American dish that's perfect for a cozy night in.",
			ImageUrl:         "https://www.themealdb.com/images/media/meals/vvpprx1511794188.jpg",
			// base64
			ImagePreview: "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxMTEhUTExMWFhUXGBgYGBgYGBgYGBgYGBgYFxgYFxgYHSggGBolHRgYITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0OGxAQGy0lICUt",
			Calories:     800,
			Likes:        0,
			// effort between 1-5, consider the washing, cutting, and peeling effort of ingredients
			TotalEffort:      5,
			InstructionSteps: "1. Preheat oven to 400°. In a large resealable plastic bag, combine flour, 1 teaspoon salt, and 1/2 teaspoon pepper. Add chicken, seal bag, and shake to coat.\n2. In a large skillet over medium heat, heat 2 tablespoons oil. Add chicken and cook until golden and no longer pink, 8 minutes per side. Transfer to a plate.\n3. In the same skillet, heat remaining oil. Add garlic and cook until fragrant, 1 minute. Add marinara and bring to a simmer. Return chicken to skillet and top with mozzarella and Parmesan.\n4. Bake until cheese is melty, 10 minutes. Garnish with basil and serve.",
			PrepTime:         int64(time.Duration(20) * time.Minute),
			CookTime:         int64(time.Duration(30) * time.Minute),
			TotalTime:        int64(time.Duration(50) * time.Minute),
		},
	}

	for _, mealParams := range mealsParams {
		m, err := d.CreateMeal(ctx, mealParams)
		if err != nil {
			fmt.Printf("Error creating meal: %v\n", err)
			return
		}
		fmt.Printf("Created meal: %v\n", m)
	}

}
