package main

import (
	"monkey-lang/token"
)

func main() {
	tok := token.Token{
		Type:   token.COMMA,
		String: ",",
	}
	tok.Display()
}
