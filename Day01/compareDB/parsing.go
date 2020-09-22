package main

import (
	"encoding/json"
	"encoding/xml"
)

type Cake struct {
	Name        string `json:"name" xml:"name"`
	Time        string `json:"time" xml:"stovetime"`
	Ingredients []struct {
		IngredientName  string `json:"ingredient_name" xml:"itemname"`
		IngredientCount string `json:"ingredient_count" xml:"itemcount"`
		IngredientUnit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
	} `json:"ingredients" xml:"ingredients>item"`
}

type Recipes struct {
	Cake [] Cake `json:"cake" xml:"cake"`
}

type JSON struct {
}

type XML struct {
}

func (r *JSON) recipy(dat []byte) (Recipes, error) {
	var data Recipes
	err := json.Unmarshal(dat, &data)

	return data, err
}

func (r *XML) recipy(dat []byte) (Recipes, error) {
	var data Recipes
	err := xml.Unmarshal(dat, &data)

	return data, err
}

type DBReader interface {
	recipy(dat []byte) (Recipes, error)
}

