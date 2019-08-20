package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix

	switch lang {
	case "Spanish":
		prefix = spanishPrefix
	}

	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("Go", ""))
}
