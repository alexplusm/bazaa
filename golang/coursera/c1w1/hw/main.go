package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

//Writable kek
type Writable interface {
	Write([]byte) (int, error)
}

const (
	lastVerticalSymbol = "└"
	verticalSymbol     = "├"
	dashes             = "───"
	pipe               = "│"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(stdout Writable, path string, printFiles bool) error {
	var err error

	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	filesCount := len(filesInfo)

	sort.Slice(filesInfo, func(i, j int) bool {
		return filesInfo[i].Name() < filesInfo[j].Name()
	})

	for index, fileInfo := range filesInfo {
		fileName := fileInfo.Name()
		if fileName == ".DS_Store" || (!fileInfo.IsDir() && printFiles) {
			continue
		}

		namePrefix(stdout, pathNestingLevel(path))
		if index == filesCount-1 {
			stdout.Write([]byte(lastVerticalSymbol))
		} else {
			stdout.Write([]byte(verticalSymbol))
		}
		stdout.Write([]byte(dashes))

		// fmt.Println(fileName, "(", fileInfo.Size(), "b)")
		fileSize := fileInfo.Size()
		if fileSize != 0 {
			fmt.Printf("%s (%db)\n", fileName, fileInfo.Size())
		} else {
			fmt.Printf("%s (empty)\n", fileName)
		}

		if fileInfo.IsDir() {
			dirTree(stdout, path+"/"+fileName, printFiles)
		}
	}

	return nil
}

func namePrefix(stdout Writable, lvl int) {
	bytesToWrite := []byte(pipe + "\t")
	for i := 0; i < lvl-1; i++ {
		stdout.Write(bytesToWrite)
	}
}

func pathNestingLevel(path string) (lvl int) {
	pathParts := strings.Split(path, "/")
	lvl = len(pathParts)
	return
}
