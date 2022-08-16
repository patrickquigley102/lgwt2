package main

import "fmt"

func main() {
	fmt.Println(Hello("World"))
}

const englishPrefix = "Hello "

// Hello returns greeting to world
func Hello(name string) string {
	return englishPrefix + name
}
