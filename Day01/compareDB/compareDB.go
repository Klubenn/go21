package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func checkInput() (string, string) {
	var old, new1 bool
	oldBase := flag.String("old", "", "old database file")
	newBase := flag.String("new", "", "new database file")
	flag.Parse()
	if *oldBase == "" || *newBase == "" {
		log.Fatalln("Both: old and new database files should be specified")
	}
	if old = strings.HasSuffix(*oldBase, ".json"); !old {
		old = strings.HasSuffix(*oldBase, ".xml")
	}
	if new1 = strings.HasSuffix(*newBase, ".json"); !new1 {
		new1 = strings.HasSuffix(*newBase, ".xml")
	}
	if !old || !new1 {
		log.Fatalln("Each database should have either '.json' or '.xml' extension")
	}
	return *oldBase, *newBase
}

func getInfoFromDB(file string) Recipes {
	var db DBReader
	var js *JSON
	var xm *XML

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("Error: can't read from file:", err)
	}

	if strings.HasSuffix(file, ".json") {
		db = js
	} else {
		db = xm
	}
	result, err := db.recipy(dat)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

func compareRecipy(oldCake, newCake Cake) {
	if oldCake.Time != newCake.Time {
		fmt.Println("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"", oldCake.Name, newCake.Time, oldCake.Time)
	}
	for _, oldIngr := range oldCake.Ingredients {
		var ingrPresent bool
		for _, newIngr := range newCake.Ingredients {
			if oldIngr.IngredientName == newIngr.IngredientName {
				ingrPresent = true
				if oldIngr.IngredientCount != newIngr.IngredientCount {
					fmt.Println("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"", oldIngr.IngredientCount, oldCake.Name, newIngr.IngredientCount, oldIngr.IngredientCount)
				}
				if oldIngr.IngredientUnit != newIngr.IngredientUnit{
					if oldIngr.IngredientUnit != "" && newIngr.IngredientUnit != "" {
						fmt.Println("CHANGED unit for ingredient \"Flour\" for cake  \"Red Velvet Strawberry Cake\" - \"mugs\" instead of \"cups\"")
					} else if 
				}
			}
		}
	}
}

func compareDB(oldDB, newDB Recipes) {
	for _, oldCake := range oldDB.Cake {
		var cakePresent bool
		for _, newCake := range newDB.Cake {
			if oldCake.Name == newCake.Name {
				cakePresent = true
				compareRecipy(oldCake, newCake)
			}
		}
		if !cakePresent {
			fmt.Println("REMOVED cake \"%s\"", oldCake.Name)
		}
	}
	for _, newCake := range newDB.Cake {
		var cakePresent bool
		for _, oldCake := range oldDB.Cake {
			if newCake.Name == oldCake.Name {
				cakePresent = true
			}
		}
		if !cakePresent {
			fmt.Println("ADDED cake \"%s\"", newCake.Name)
		}
	}
}

func main() {
	oldFile, newFile := checkInput()
	oldDB := getInfoFromDB(oldFile)
	newDB := getInfoFromDB(newFile)
	compareDB(oldDB, newDB)
}
