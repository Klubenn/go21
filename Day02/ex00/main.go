package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Flags struct {
	sl		bool
	d		bool
	f		bool
	ext		string
	path	string
}

func flagsParse() Flags {
	var flags Flags

	flagSl := flag.Bool("sl", false, "print symlinks")
	flagD := flag.Bool("d", false, "print directories")
	flagF := flag.Bool("f", false, "print files")
	flagExt := flag.Bool("ext", false, "print files with certain extension - can be specified only with '-f' flag")
	flag.Parse()

	if *flagExt && !*flagF {
		fmt.Println("Flag '-ext' can be specified only with '-f' flag")
		os.Exit(1)
	}
	if !*flagSl && !*flagD && !*flagF {
		*flagSl = true
		*flagF = true
		*flagD = true
	}
	path := flag.Args()
	if len(path) != 1 && !*flagExt {
		fmt.Println("Specify one path for the search")
		os.Exit(1)
	} else if len(path) != 2 && *flagExt {
		fmt.Println("Specify extension and a path for the search")
		os.Exit(1)
	}
	flags.sl = *flagSl
	flags.d = *flagD
	flags.f = *flagF
	if *flagExt {
		flags.ext = path[0]
	}
	flags.path = path[len(path)-1]
	return flags
}

func printFile(name string, details Flags) {
	if details.ext != "" && strings.HasSuffix(name, details.ext) || details.f && details.ext == "" {
		fmt.Println(name)
	}
}

func processDir(path string, details Flags) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		log.Fatalln(err)
	}
	if details.sl && fileInfo.Mode() & os.ModeSymlink != 0{
		s, err := os.Readlink(path)
		if err == nil {
			fmt.Println(path, "->", s)
		} else {
			fmt.Println(path, "-> [broken]")
		}
	} else if fileInfo.Mode() & os.ModeSymlink == 0 {
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		if details.d {
			fmt.Println(path)
		}
		for _, name := range dir {
			fullName := filepath.Join(path, name.Name())
			file, err := os.Stat(fullName)
			if err != nil {
				_, err := os.Lstat(path)
				if err == nil && details.sl {
					fmt.Println(fullName, "-> [broken]")
				} else if err != nil {
					fmt.Println(err)
				}
				continue
			}
			switch fileType := file.Mode(); {
			case fileType.IsRegular():
				printFile(filepath.Join(path, name.Name()), details)
			case fileType.IsDir():
				processDir(filepath.Join(path, name.Name()), details)
			}
		}
	}
}

func main() {
	details := flagsParse()

	file, err := os.Stat(details.path)
	if err != nil {
		log.Fatalln(err)
	}
	switch fileType := file.Mode(); {
	case fileType.IsRegular():
		printFile(details.path, details)
		os.Exit(0)
	case fileType.IsDir():
		processDir(details.path, details)
		os.Exit(0)
	}
}
