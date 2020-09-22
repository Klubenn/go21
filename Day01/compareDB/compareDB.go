package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	var old, new bool
	oldBase := flag.String("old", "", "old database file")
	newBase := flag.String("new", "", "new database file")
	flag.Parse()
	if *oldBase == "" || *newBase == "" {
		log.Fatalln("Both: old and new database files should be specified")
	}
	if old = strings.HasSuffix(*oldBase, ".json"); old {
		if new = strings.HasSuffix(*newBase, ".xml"); !new {
			log.Fatalln("Old and ")
		}
	}
	fmt.Println(*oldBase, *newBase)
}
