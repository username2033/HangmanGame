package Parsers

import (
	"encoding/json"
	"io/ioutil"
)

type Riddle struct {
	Word  string   `json:"word"`  // Теги JSON внутри определения
	Clues []string `json:"clues"` // Теги JSON внутри определения
}

type Level struct {
	Difficalty string   `json:"difficulty"` // Теги JSON внутри определения
	Words      []Riddle `json:"words"`      // Теги JSON внутри определения
}

type Categories struct {
	Arr []string `json:"Categories"`
}

//type parseLevels struct{}

func ParseLevelsFile(filename string) ([]Level, error) {
	pathToFile := "./Config/" + filename
	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}

	var levels []Level
	err = json.Unmarshal(data, &levels)
	if err != nil {
		return nil, err
	}

	return levels, nil
}

//type ParseCategories struct{}

func ParseCategoriesFile(filename string) (Categories, error) {
	//categories, err := os.Open("./Config/Categories.json")
	nilCategories := Categories{
		Arr: []string{}, // Явное указание пустого среза
	}
	pathToFile := "./Config/" + filename
	data, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return nilCategories, err
	}
	var categories Categories
	err = json.Unmarshal(data, &categories)
	if err != nil {
		return nilCategories, err
	}

	return categories, nil
}
