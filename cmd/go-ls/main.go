package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	lFlag := flag.Bool("l", false, "multi line print")
	aFlag := flag.Bool("a", false, "show hidden files")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		os.Exit(-1)
	}
	if flag.Arg(0) != "" {
		dir = flag.Arg(0)
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("No such file or directory")
	}

	for _, fileInfo := range fileInfos {
		// show hidden files
		if *aFlag == false {
			if strings.HasPrefix(fileInfo.Name(), ".") {
				continue
			}
		}
		fmt.Print(fileInfo.Name())
		if *lFlag == false {
			fmt.Print("\t")
		} else {
			fmt.Println()
		}
	}
	if *lFlag == false {
		fmt.Println()
	}
}
