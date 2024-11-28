package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	args := os.Args

	if len(args) < 2 {
		fmt.Println("input needed , append folder/file path to your command above")
		return
	}

	rdb.FlushAll(ctx)

	listFilesInDir(args[1], rdb)
}

func listFilesInDir(path string, rdb *redis.Client) {

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
			listFilesInDir(filepath.Join(path, file.Name()), rdb)
		} else {
			fmt.Println(filepath.Join(path, file.Name()), "  ", file.Name())

			// err := rdb.Set(ctx, file.Name(), filepath.Join(path, file.Name()), 0).Err()
			err = rdb.SAdd(ctx, file.Name(), filepath.Join(path, file.Name())).Err()
			if err != nil {
				panic(err)
			}
		}
	}
}


