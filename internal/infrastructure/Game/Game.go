package Game

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Features"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Riddles"
	"strings"
)

func Play(category string) {

	word, _ := Riddles.GetWord(category)
	idSet := Features.GenIntSet()
	answeredLettres := Features.GenStringSet()
	numberOfUsedClues := 0
	allowedNumberOfClues := len(word.Clues)
	attempts := len(Features.HangmanStages) // Количество попыток
	badAttempts := 0                        // Количество попыток на предыдущей попытке
	numberOfCorrectLetters := 0
	outputstring := Features.GenStringSlice(len(word.Word))
	clueFlag := "Да"
	//fmt.Println(word.Word, make([]bool, len(word.Word)))
	//fmt.Println(word.Clues)
	Features.PrintSlice(&outputstring)
	set := Features.StringToSet(word.Word)
	var letter string
	fmt.Println("У вас ", attempts, " попыток")
	for (attempts) > 0 && (numberOfCorrectLetters != len(word.Word)) {
		if (badAttempts >= 2) && (numberOfUsedClues < allowedNumberOfClues) {
			fmt.Println("хотите подсказку? (Yes, No)")
			fmt.Scanln(&clueFlag)
			if strings.ToLower(clueFlag) == "yes" {
				numberOfUsedClues++
				id := Features.GetRandomindex(len(word.Clues))
				_, exists := idSet[id]
				//fmt.Println(id, exists)
				for exists {
					id := Features.GetRandomindex(len(word.Clues))
					_, exists = idSet[id]
				}
				fmt.Println(word.Clues[id])
				idSet[id] = true
				if numberOfUsedClues == allowedNumberOfClues {
					fmt.Println("Вы использовали все подсказки")
				}
			}
		}
		fmt.Println("Введите букву")
		fmt.Scanln(&letter)
		letter = strings.ToLower(letter)
		_, exists := answeredLettres[letter]
		for exists {
			fmt.Println("Вы уже вводили данную букву, введите другую букву")
			fmt.Scanln(&letter)
			letter = strings.ToLower(letter)
			_, exists = answeredLettres[letter]
		}
		answeredLettres[letter] = true
		//fmt.Println(letter)
		if _, exists := set[letter]; exists {
			fmt.Printf("Буква %s есть в слове! \n", letter)
			numberOfCorrectLetters += Features.EnterALetter(&outputstring, &word.Word, letter)
			if numberOfCorrectLetters == len(word.Word) {
				fmt.Println("Поздравляем, вы отгадали слово!!!")
			}
			Features.PrintSlice(&outputstring)
			badAttempts = 0
		} else {
			fmt.Printf("Буквы %s нет в слове \n", letter)
			Features.PrintHangman(len(Features.HangmanStages) - attempts)
			if attempts == 1 {
				fmt.Println("Вы проиграли(")
				break
			}
			badAttempts++
			attempts--
			Features.PrintSlice(&outputstring)
			fmt.Println("У вас осталось ", attempts, " попыток")

		}
		//attempts = 0

	}
}
