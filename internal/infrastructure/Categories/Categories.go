package Categories

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Features"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Parsers"
)

func getCategories() ([]string, error) {
	categories, err := Parsers.ParseCategoriesFile("Categories.json")
	if err != nil {
		fmt.Println("Ошибка при парсинге файла:", err)
		nilCategories := Parsers.Categories{
			Arr: []string{}, // Явное указание пустого среза
		}
		return nilCategories.Arr, err
	}
	//id := rand.Intn(len(riddles))
	return categories.Arr, nil
}

func GetCategory() (string, error) {
	categoriesArray, err := getCategories()
	if err != nil {
		fmt.Println("Ошибка парсинга категорий", err)
		return "", err
	}
	categoriesSet := Features.GenStringSet()
	var category string
	//fmt.Println(err)
	//fmt.Println(categoriesArray)
	for _, val := range categoriesArray {
		categoriesSet[val] = true
		fmt.Println(val)
	}
	fmt.Println("Выберите категорию")
	fmt.Scanln(&category)
	_, exists := categoriesSet[category]
	for exists == false {
		fmt.Println("Данной категории нет в списке категорий")
		fmt.Println("Выберите категорию")
		fmt.Scanln(&category)
		_, exists = categoriesSet[category]
	}
	return category, nil
}
