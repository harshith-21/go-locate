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

	if args[1] == "--update" {
		refresh(rdb, "/")
	} else if args[1] == "--dev" {
		refresh(rdb, args[2])
	} else {
		locateFile(rdb, args[1])
	} 
}

func listFilesInDir(path string, rdb *redis.Client) {

	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {

		if file.IsDir() {
			listFilesInDir(filepath.Join(path, file.Name()), rdb)
		} else {
			err = rdb.SAdd(ctx, file.Name(), filepath.Join(path, file.Name())).Err()
			if err != nil {
				panic(err)
			}
		}
	}
}

func locateFile(rdb *redis.Client, filename string) {
	paths, err := rdb.SMembers(ctx, filename).Result()

	if err == redis.Nil {
		fmt.Println("Key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		for _, path := range paths {
			fmt.Println(path)
		}
	}
}

func refresh(rdb *redis.Client, path string) {
	rdb.FlushAll(ctx)
	listFilesInDir(path, rdb)
}