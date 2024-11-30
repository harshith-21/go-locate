package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

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

	var wg sync.WaitGroup

	if len(args) < 2 {
		fmt.Println("input needed , append folder/file path to your command above")
		return
	}

	wg.Add(1)
	if args[1] == "--update" {
		refresh(rdb, "/", &wg)
		wg.Wait()
	} else if args[1] == "--dev" {
		refresh(rdb, args[2], &wg)
		wg.Wait()
	} else {
		locateFile(rdb, args[1])
	} 
}


func listFilesInDir(path string, rdb *redis.Client, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the WaitGroup counter when this goroutine finishes

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			wg.Add(1) // Increase the counter for the new goroutine
			go listFilesInDir(filepath.Join(path, file.Name()), rdb, wg) // Launch a new goroutine for the subdirectory
		} else {
			// Handle Redis insertion in the same goroutine
			err = rdb.SAdd(ctx, file.Name(), filepath.Join(path, file.Name())).Err()
			if err != nil {
				fmt.Println("Error adding to Redis:", err)
				return
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

func refresh(rdb *redis.Client, path string, wg *sync.WaitGroup) {
	rdb.FlushAll(ctx)
	listFilesInDir(path, rdb, wg)
}