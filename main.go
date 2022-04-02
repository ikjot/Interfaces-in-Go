// Interfaces helps to reuse logics.
package main

import "fmt"

// 1) Read it as: To all the different types in here: there's another custom type "bot"
type bot interface {
	// 2) If you are a type in this program with a func called "getGreeting", and
	// you return a string, then you are now an honorary member of type "bot".
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// func (eb englishBot) getGreeting() string { equivalent to:
func (englishBot) getGreeting() string {
	// Very custom logic for generating english greeting.
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// Same function declared twice: error
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// 3) Now that your're also an honorary member of type "bot", you can
// now call this function called "printGreeting".
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// $ go run main.go
// Hi There!
// Hola!

// Notes:
// 1) Interfaces are not generics types.
// 2) Interfaces are 'implicit'. e.g. We didn't write explicitly that ennglishbot and spanishbot are types of bot.
// 3) Interfaces are a contract to help us manage "types". They are not unit tests as they don't test the logic.
// 4) Interfaces are tough. Step #1 is understanding how to read them.
