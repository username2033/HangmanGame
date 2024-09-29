package Features

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Parsers"
	"math/rand"
	"time"
)

var HangmanStages = []string{
	`
      ------
      |    |
      |    
      |   
      |   
      |   
    --------
    `,
	`
      ------
      |    |
      |    O
      |   
      |   
      |   
    --------
    `,
	`
      ------
      |    |
      |    O
      |    |
      |   
      |   
    --------
    `,
	`
      ------
      |    |
      |    O
      |   /|
      |   
      |   
    --------
    `,
	`
      ------
      |    |
      |    O
      |   /|\
      |   
      |   
    --------
    `,
	`
      ------
      |    |
      |    O
      |   /|\
      |   / 
      |   
    --------
    `,
	`
      ------
      |    |
      |    O
      |   /|\
      |   / \
      |   
    --------
    `,
}

func StringToSet(s string) map[string]bool {
	set := GenStringSet() // Создаем множество

	for _, char := range s { // Итерируем по символам строки
		set[string(char)] = true // Добавляем символ в множество
	}

	return set
}

func GenIntSet() map[int]bool {
	return make(map[int]bool)
}

func GenStringSet() map[string]bool {
	return make(map[string]bool)
}

func GenStringSlice(lenght int) []string {

	stringSlice := make([]string, lenght)
	for i := range stringSlice {
		stringSlice[i] = "_" // Заполняем значением по умолчанию
	}
	return stringSlice
}

func EnterALetter(arr *[]string, word *string, letter string) int {
	counter := 0
	indexesSet := GenIntSet()
	for id, char := range *word {
		if string(char) == letter {
			indexesSet[id] = true
			counter++
		}
	}
	for id, _ := range indexesSet {
		(*arr)[id] = letter
	}
	return counter
}

func PrintSlice(arr *[]string) {
	for _, elem := range *arr {
		fmt.Print(elem)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func PrintStringSet(set *map[string]bool) {
	for elem, _ := range *set {
		fmt.Print(elem)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func PrintHangman(id int) {
	fmt.Println(HangmanStages[id])
}

func PrintDict(set *map[string][]Parsers.Riddle) {
	for elem, _ := range *set {
		fmt.Print(elem)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func GetRandomindex(lenght int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(lenght)
}
