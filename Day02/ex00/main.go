package main

import (
	"flag"
	"fmt"
	"os"
)

func flagsParse() (bool, bool, bool, bool, []string) {
	flagSl := flag.Bool("sl", false, "print symlinks")
	flagD := flag.Bool("d", false, "print directories")
	flagF := flag.Bool("f", false, "print files")
	flagExt := flag.Bool("ext", false, "print files with certain extension - can be specified only with '-f' flag")
	flag.Parse()

	if *flagExt && !*flagF {
		fmt.Println("Flag '-ext' can be specified only with '-f' flag")
		os.Exit(1)
	}
	if noFlags := !*flagSl && !*flagD && !*flagF; noFlags {
		*flagSl = true
		*flagF = true
		*flagD = true
	}
	path := flag.Args()
	if len(path) != 1 && !*flagExt{
		fmt.Println("Specify one path for the search")
		os.Exit(1)
	} else if len(path) != 2 && *flagExt {
		fmt.Println("Specify extension and a path for the search")
	}

	return *flagSl, *flagD, *flagF, *flagExt, path
}

func main() {
	sl, d, f, ext, path := flagsParse()



}
