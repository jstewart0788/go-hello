package main

import (
	"github.com/jstewart0788/greeting"
)

func main() {
	var s = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("?"), true)
}
