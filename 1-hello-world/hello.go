package main

import "fmt"

func main() {
	fmt.Println(Hello("World"))
}

const englishPrefix = "Hello "

// Hello returns greeting to world
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}
