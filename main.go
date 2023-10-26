package main

import (
	"fmt"
	"monkey-lang/lexer"
	"monkey-lang/repl"
	"monkey-lang/token"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "repl" {
			fmt.Println("Type Anything to start")
			repl.Start()
		}
	}
	if len(os.Args) < 3 {
		fmt.Println("usage: repl TO START REPL")
		fmt.Println("usage: <sourcefile> <output>")
		os.Exit(1)
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	fmt.Println(arg1, " ", arg2)
	contents, err := os.ReadFile(arg1)
	if err != nil {
		fmt.Println(err)
	}
	input := string(contents)

	l := lexer.NewLexer(input)
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		tok.Display()
	}
}
