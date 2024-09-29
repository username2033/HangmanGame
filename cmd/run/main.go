package main

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Categories"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/Game"
)

func main() {
	// TODO: write your code here
	category, err := Categories.GetCategory()
	if err == nil {
		Game.Play(category)
	} else {
		fmt.Println("Невозможно получить категорию")
	}
}
