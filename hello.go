package main

import "fmt"

const romanian = "Romanian"
const german = "German"
const french = "French"
const englishRelayWorkPrefix = "Is the relay ready, "
const romanianRelayWorkPrefix = "E gata releul, "
const frenchRelayWorkPrefix = "Le relais est-il prêt, "
const germanRelayWorkPrefix = "Ist das Relais bereit, "

func Hello(name, language string) string {
	if name == "" {
		name = "freaks"
	}

	return relayWorkPrefix(language) + name + "?"
}

func relayWorkPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchRelayWorkPrefix
	case romanian:
		prefix = romanianRelayWorkPrefix
	case german:
		prefix = germanRelayWorkPrefix
	default:
		prefix = englishRelayWorkPrefix
	}
	return 
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
