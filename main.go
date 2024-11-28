package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

var ctx = context.Background()

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("input needed , append folder/file path to your command above")
		return
	}

	listFilesInDir(args[1])
}

func listFilesInDir(path string) {

	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {

		if file.Name() == ".git" {
			continue
		}

		if file.IsDir() {
			listFilesInDir(filepath.Join(path, file.Name()))
		} else {
			fmt.Println(filepath.Join(path, file.Name()), "  ", file.Name())
			if err != nil {
				panic(err)
			}
		}
	}

}
