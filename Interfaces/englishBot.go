package main 

type englishBot struct{}

func (englishBot) getGreeting() string {
	// Very custom logic for english
	return "Hello there!"
}