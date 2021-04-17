package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		os.Exit(-1)
	}
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		os.Exit(-1)
	}

	for _, fileInfo := range fileInfos {
		if !strings.HasPrefix(fileInfo.Name(), ".") {
			fmt.Print(fileInfo.Name())
			fmt.Print("\t")
		}
	}
	fmt.Println()
}
