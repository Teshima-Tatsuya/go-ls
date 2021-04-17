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
		if !strings.HasPrefix(fileInfo.Name(), ".") {
			fmt.Print(fileInfo.Name())
			fmt.Print("\t")
		}
	}
	fmt.Println()
}
