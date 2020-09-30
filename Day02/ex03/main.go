package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type flagInfo struct {
	path	string
	args	[]string
}

func newName(file string, path string) (string, error) {
	var newname string
	fileStat, err := os.Stat(file)
	if err == nil {
		timeStamp := fileStat.ModTime().Unix()
		newname = filepath.Base(file)
		newname = strings.TrimSuffix(newname, filepath.Ext(newname)) + "_" + fmt.Sprint(timeStamp)
		if path != "" {
			newname = filepath.Join(path, newname) + ".tar.gz"
		} else {
			newname = filepath.Join(filepath.Dir(file), newname) + ".tar.gz"
		}
	}
	return newname, err
}

func archiveLog(file string, path string, c chan error) {
	newname, err := newName(file, path)
	if err != nil {
		c <- err
		return
	}
	cmd := exec.Command("tar", "-czvf", newname, file)
	q, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("%v, %v", err, string(q))
	}
	c <- err
}

func parseFlags()  flagInfo {
	aFlag := flag.String("a", "", "specifies output directory for archived log")
	flag.Parse()

	info := flagInfo {path: *aFlag, args: flag.Args()}
	return info
}

func main() {
	info := parseFlags()
	length := len(info.args)
	c := make(chan error, length)
	for _, file := range info.args {
		go archiveLog(file, info.path, c)
	}
	for i := 0; i < length; i++ {
		switch a := <-c; a {
		case nil:
			continue
		default:
			fmt.Println("An error occured:", a)
		}
	}
}