package main

import "fmt"

const englishRelayWorkPrefix = "Is the relay ready, "

func Hello(name string) string {
	if name == "" {
		name = "freaks"
	}
	return englishRelayWorkPrefix + name + "?"
}

func main() {
	fmt.Println(Hello("Chris"))
}
