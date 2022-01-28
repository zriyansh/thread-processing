// this is an implementation of single and multi thread search algo to find files in a OS, given a root address and file name.

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	waitgroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
	// we meed locking system because multiple threads may want to update the 'matches slice of string at a time, creating duplicate data'
)

func fileSearchSingleThread(root, filename string) {
	fmt.Println("going to search in ", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			matches = append(matches, filepath.Join(root, file.Name())) // C:\tools\readme.md, adds filename to root name.
		}
		if file.IsDir() { // if file found is a directory, we make a recursive call
			fileSearchSingleThread(filepath.Join(root, file.Name()), filename)
		}
	}
}

func fileSearchMultiThread(root, filename string) {
	fmt.Println("going to search in ", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name())) // C:\tools\readme.md, adds filename to root name.
			lock.Unlock()
		}
		if file.IsDir() { // if file found is a directory, we make a recursive call
			waitgroup.Add(1)
			go fileSearchMultiThread(filepath.Join(root, file.Name()), filename)
		}
	}
	waitgroup.Done()
}

func main() {
	root := "/Users/priyanshkhodiyar"
	fileName := "pic.png"

	waitgroup.Add(1)
	go fileSearchMultiThread(root, fileName)
	waitgroup.Wait()
	// comment and above 3 lines and uncomment the fileSearchSingleThread() to run single thread file search and vice cersa.

	// change thr root and filename and search for whatever you want,
	// if filename = "pic.png", it will not exactly match for that, but "my_pic.png" file will also be a valid result

	// fileSearchSingleThread(root, fileName)

	for _, file := range matches {
		fmt.Println("Found file ", file)
	}
}
