package main

import (
	"fmt"
	"sort"
)

// 1. Определите структуру BrainrotMeme
type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64 // в миллионах
}

// 3. Реализуйте три аналитические функции

// FindTopTrending возвращает массив мемов, у которых количество просмотров больше minViews,
// отсортированный по убыванию TrendLevel.
func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var result []BrainrotMeme

	// Фильтрация
	for _, m := range memes {
		if m.Views > minViews {
			result = append(result, m)
		}
	}

	// Сортировка по убыванию TrendLevel
	sort.Slice(result, func(i, j int) bool {
		return result[i].TrendLevel > result[j].TrendLevel
	})

	return result
}

// CalculateCategoryImpact возвращает map, где ключ — категория,
// а значение — суммарное количество просмотров всех мемов этой категории.
func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	impact := make(map[string]float64)

	for _, m := range memes {
		impact[m.Category] += m.Views
	}

	return impact
}

// FilterByComplexCondition возвращает массив названий мемов, которые удовлетворяют сложному условию:
// TrendLevel >= 7 ИЛИ (количество просмотров > 50 И категория "Sigma").
func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var names []string

	for _, m := range memes {
		if m.TrendLevel >= 7 || (m.Views > 50 && m.Category == "Sigma") {
			names = append(names, m.Name)
		}
	}

	return names
}

func main() {
	// 2. Создайте срез из минимум 7 различных Brainrot-мемов
	memes := []BrainrotMeme{
		{Name: "Skibidi Toilet Episode 1", TrendLevel: 10, Category: "Skibidi", Views: 150.5},
		{Name: "Sigma Face", TrendLevel: 9, Category: "Sigma", Views: 80.2},
		{Name: "Looksmaxxing Tutorial", TrendLevel: 8, Category: "Mewing", Views: 45.0},
		{Name: "Subo Prank", TrendLevel: 5, Category: "Subo Bratik", Views: 20.1},
		{Name: "Sahur Sahur", TrendLevel: 6, Category: "TUNTUNTUNSAHUR", Views: 15.5},
		{Name: "Random Cat", TrendLevel: 3, Category: "Other", Views: 5.0},
		{Name: "Alpha Male Grindset", TrendLevel: 6, Category: "Sigma", Views: 60.0},
		{Name: "Mewing Streak", TrendLevel: 7, Category: "Mewing", Views: 30.5},
	}

	fmt.Println("=== Исходный список мемов ===")
	for _, m := range memes {
		fmt.Printf("- %s (Trend: %d, Views: %.1fM, Cat: %s)\n", m.Name, m.TrendLevel, m.Views, m.Category)
	}
	fmt.Println()

	// Демонстрация FindTopTrending
	fmt.Println("=== Top Trending (Views > 40M, Sorted by Trend) ===")
	topTrending := FindTopTrending(memes, 40.0)
	for _, m := range topTrending {
		fmt.Printf("%s: Trend %d, Views %.1fM\n", m.Name, m.TrendLevel, m.Views)
	}
	fmt.Println()

	// Демонстрация CalculateCategoryImpact
	fmt.Println("=== Category Impact (Total Views) ===")
	impact := CalculateCategoryImpact(memes)
	for cat, views := range impact {
		fmt.Printf("%s: %.1fM\n", cat, views)
	}
	fmt.Println()

	// Демонстрация FilterByComplexCondition
	fmt.Println("=== Complex Condition Filter (Trend >= 7 OR (Views > 50 AND Category 'Sigma')) ===")
	filteredNames := FilterByComplexCondition(memes)
	for _, name := range filteredNames {
		fmt.Println("-", name)
	}
}
