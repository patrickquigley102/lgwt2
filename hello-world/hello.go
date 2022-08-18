package main

import "fmt"

func main() {
	fmt.Println(Hello("World", "English"))
}

const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bounjour "

// Hello returns greeting to world
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return buildGreeting(language) + name
}

func buildGreeting(language string) string {
	greeting := englishPrefix
	switch language {
	case "french":
		greeting = frenchPrefix
	case "spanish":
		greeting = spanishPrefix
	}
	return greeting
}
