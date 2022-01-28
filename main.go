package main

import "fmt"

var (
	matches []string
)

func fileSearch(root, filename string) {
	fmt.Printf("going to search in ", root)

}

func main() {
	fileSearch("C:/tools", "readme.md")
	for _, file := range matches {
		fmt.Println("Found file", file)
	}
}
