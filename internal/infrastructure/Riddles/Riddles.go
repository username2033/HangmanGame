package Riddles

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Features"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Parsers"
)

func printDict(set *map[string][]Parsers.Riddle) {
	for elem, _ := range *set {
		fmt.Print(elem)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func getDict(filename string) ([]Parsers.Level, error) {
	//filename := category + ".json"
	levels, err := Parsers.ParseLevelsFile(filename)
	//fmt.Println(filename)
	if err != nil {
		fmt.Println("Ошибка при парсинге файла:", err, filename)
		nilLevels := make([]Parsers.Level, 0)
		return nilLevels, err
	}
	//id := rand.Intn(len(riddles))
	return levels, nil
}

func getRiddles(dict map[string][]Parsers.Riddle) []Parsers.Riddle {
	fmt.Println("Выберите сложность игры из предложенных")
	printDict(&dict)
	var difficalty string
	fmt.Scanln(&difficalty)
	riddles, exists := dict[difficalty]
	for exists == false {
		fmt.Println("Загадок данной сложности нет в списке загадок")
		fmt.Println("Выберите сложность игры из предложенных")
		fmt.Scanln(&difficalty)
		riddles, exists = dict[difficalty]
	}
	return riddles
}

func GetWord(category string) (Parsers.Riddle, error) {

	filename := category + ".json"
	lvl, err := getDict(filename)
	dict := make(map[string][]Parsers.Riddle)

	if err != nil {
		fmt.Println("Ошибка при парсинге файла:", err)
		nilriddle := Parsers.Riddle{
			Word:  "",         // Явное указание пустой строки
			Clues: []string{}, // Явное указание пустого среза
		}
		return nilriddle, err
	}

	for _, level := range lvl {
		dict[level.Difficalty] = level.Words
	}

	riddles := getRiddles(dict)
	return riddles[Features.GetRandomindex(len(riddles))], nil

}
