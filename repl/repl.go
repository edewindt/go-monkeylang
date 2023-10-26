package repl

import (
	"bufio"
	"fmt"
	"monkey-lang/lexer"
	"monkey-lang/token"
	"os"
)

func Start() {
	for {
		fmt.Print(">> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		lexer := lexer.NewLexer(input)
		for {
			tok := lexer.NextToken()
			if tok.Type == token.EOF {
				break
			}
			tok.Display()
		}
	}
}
