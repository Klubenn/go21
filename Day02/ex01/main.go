package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type myFlags struct {
	w	bool
	l	bool
	m	bool
}

type data struct {
	num		int
	file	string
}

func parseFlags() (myFlags, []string) {
	w := flag.Bool("w", false, "count words")
	l := flag.Bool("l", false, "count lines")
	m := flag.Bool("m", false, "count characters")
	flag.Parse()

	if !*w && !*l && !*m {
		*w = true
	}
	if *w && *l && *m || *w && *l || *w && *m || *l && *m {
		fmt.Println("only one flag can be specified:\n-w\tcount words\n-l\tcount lines\n-m\tcount characters")
		os.Exit(1)
	}
	flags := myFlags{w: *w, l: *l, m: *m}
	return  flags, flag.Args()
}

func parseFile(file string, flags myFlags, c chan *data) {
	var lines int
	var words int
	var chars int
	var myData data
	dat, err := os.Open(file)
	if err != nil {
		c <- nil
		return
	}
	defer dat.Close()
	read := bufio.NewScanner(dat)
	for read.Scan() {
		lines++
		str := read.Text()
		words += strings.Count(str, " ") + 1
		chars += utf8.RuneCountInString(str) + 1
	}
	if flags.w {
		myData.num = words
	} else if flags.l {
		myData.num = lines
	} else {
		myData.num = chars
	}
	myData.file = file
	c <- &myData
}

func main() {
	flags, args := parseFlags()
	length := len(args)
	c := make(chan *data, length)
	for _, file := range args {
		go parseFile(file, flags, c)
	}
	for i := 0; i < len(args); i++ {
		data := <- c
		if data != nil {
			fmt.Printf("%v\t%v\n", data.num, data.file)
		}
	}
}