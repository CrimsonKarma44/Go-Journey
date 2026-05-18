package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	if len(os.Args) < 3 {

		fmt.Println("Usage: program <dir1> <dir2>")
		os.Exit(1)
	}

	var mutex sync.Mutex
	var wg sync.WaitGroup

	// First task: collect files from dir1
	dir1 := os.Args[1]
	fileStore := make(map[string]struct{}) // Using map for faster lookups

	wg.Add(1)
	go recul(dir1, &fileStore, &wg, &mutex)
	wg.Wait()

	// Second task: compare with dir2
	dir2 := os.Args[2]
	wg.Add(1)
	go findMatchingFiles(dir2, &fileStore, &wg, &mutex)
	wg.Wait()
}

func entries(dir string) ([]fs.DirEntry, error) {
	return os.ReadDir(dir)
}

func recul(dir string, fileStore *map[string]struct{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	ent, err := entries(dir)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", dir, err)
		return
	}

	for _, item := range ent {
		fullPath := filepath.Join(dir, item.Name())

		if !item.IsDir() {
			mutex.Lock()
			(*fileStore)[item.Name()] = struct{}{} // Store just the filename
			mutex.Unlock()
		} else if item.Name()[0] != '.' {
			wg.Add(1)
			go recul(fullPath, fileStore, wg, mutex)
		}
	}
}

func findMatchingFiles(dir string, fileStore *map[string]struct{}, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	ent, err := entries(dir)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", dir, err)
		return
	}

	for _, item := range ent {
		fullPath := filepath.Join(dir, item.Name())

		if item.IsDir() && item.Name()[0] != '.' {
			wg.Add(1)
			go findMatchingFiles(fullPath, fileStore, wg, mutex)
		} else if !item.IsDir() {
			mutex.Lock()
			_, exists := (*fileStore)[item.Name()]
			mutex.Unlock()

			if exists {
				fmt.Println(fullPath) // Print full path of matching file
				if err = os.Remove(fullPath); err != nil {
					fmt.Printf("Error removing %s: %v\n", fullPath, err)
				}
			}
		}
	}
}
