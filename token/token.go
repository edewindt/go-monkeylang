package token

import "fmt"

type Token struct {
	Type   TokenType
	String string
}

func (t Token) Display() {
	fmt.Println("Type:", t.Type.to_string(), " ", "String:", t.String)
}

func Lookup_Identifier(i string) TokenType {
	switch i {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "if":
		return IF
	case "else":
		return ELSE
	case "return":
		return RETURN
	default:
		return IDENTIFIER
	}
}

type TokenType int

func (t TokenType) to_string() string {
	switch t {
	case 0:
		return "RETURN"
	case 1:
		return "ILLEGAL"
	case 2:
		return "EOF"
	case 3:
		return "IDENTIFIER"
	case 4:
		return "INTIGER"
	case 5:
		return "ASSIGN"
	case 6:
		return "PLUS"
	case 7:
		return "MINUS"
	case 8:
		return "BANG"
	case 9:
		return "ASTERISK"
	case 10:
		return "SLASH"
	case 11:
		return "LT"
	case 12:
		return "GT"
	case 13:
		return "COMMA"
	case 14:
		return "SEMICOLON"
	case 15:
		return "EQUALS"
	case 16:
		return "NOTEQUALS"
	case 17:
		return "LPAREN"
	case 18:
		return "RPAREN"
	case 19:
		return "LBRACE"
	case 20:
		return "RBRACE"
	case 21:
		return "FUNCTION"
	case 22:
		return "LET"
	case 23:
		return "TRUE"
	case 24:
		return "FALSE"
	case 25:
		return "IF"
	case 26:
		return "ELSE"
	default:
		return "ILLEGAL"
	}
}

const (
	RETURN = iota
	ILLEGAL
	EOF
	IDENTIFIER
	INTIGER
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	LT
	GT
	COMMA
	SEMICOLON
	EQUALS
	NOTEQUALS
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
)
