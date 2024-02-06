package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	stdsync "sync"

	sync "github.com/gm0stache/go-conc/internal/sync"
)

type asyncStore[T any] struct {
	values []T
	mtx    *stdsync.Mutex
}

func NewStore[T any]() *asyncStore[T] {
	return &asyncStore[T]{
		values: []T{},
		mtx:    &stdsync.Mutex{},
	}
}

func (as *asyncStore[T]) append(el T) {
	as.mtx.Lock()
	defer as.mtx.Unlock()
	as.values = append(as.values, el)
}

func (as *asyncStore[T]) getVals() []T {
	as.mtx.Lock()
	defer as.mtx.Unlock()
	return as.values
}

func main() {
	if len(os.Args) != 3 { // index 0 is always the program name, so here 'filesrch'
		printErrAndExit(errors.New("invalid invocation. usage: some_directory filesrch some_filename"))
	}

	dir := os.Args[1]
	filename := os.Args[2]
	wg := sync.NewWaitGroup()
	wg.Add(1)

	store := NewStore[string]()
	searchFile(dir, filename, store, wg) // does the 'go' here make a difference
	wg.Wait()

	paths := store.getVals()
	sort.Strings(paths)
	print(strings.Join(paths, "\n"))
}

func searchFile(dir string, filename string, store *asyncStore[string], wg *sync.WaitGroup) {
	files, err := os.ReadDir(dir)
	if err != nil {
		printErrAndExit(err)
	}

	for _, f := range files {
		fullpath := path.Join(dir, f.Name())
		if f.IsDir() {
			wg.Add(1)
			go searchFile(fullpath, filename, store, wg)
		}
		if strings.Contains(f.Name(), filename) {
			store.append(fullpath)
		}
	}

	wg.Done()
}

func printErrAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(-1)
}
