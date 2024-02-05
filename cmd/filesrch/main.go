package main

import (
	"errors"
	"fmt"
	sync "github.com/gm0stache/go-conc/internal/sync"
	"os"
	"path"
	"strings"
)

func main() {
	if len(os.Args) != 3 { // index 0 is always the program name, so here 'filesrch'
		printErrAndExit(errors.New("invalid invocation. usage: some_directory filesrch some_filename"))
	}

	dir := os.Args[1]
	filename := os.Args[2]
	wg := sync.NewWaitGroup()
	wg.Add(1)

	searchFile(dir, filename, wg) // does the 'go' here make a difference
	wg.Wait()
}

func searchFile(dir string, filename string, wg *sync.WaitGroup) {
	files, err := os.ReadDir(dir)
	if err != nil {
		printErrAndExit(err)
	}

	for _, f := range files {
		fullpath := path.Join(dir, f.Name())
		if f.IsDir() {
			wg.Add(1)
			go searchFile(fullpath, filename, wg)
		}
		if strings.Contains(f.Name(), filename) {
			fmt.Println(fullpath)
		}
	}

	wg.Done()
}

func printErrAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(-1)
}
