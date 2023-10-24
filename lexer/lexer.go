package lexer

import (
	"monkey-lang/token"
	"unicode"
)

type Lexer struct {
	Input        []rune
	Position     int
	ReadPosition int
	Char         rune
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Input:        []rune(input),
		Position:     0,
		ReadPosition: 0,
		Char:         ' ',
	}
}

func (l *Lexer) Read_Char() {
	if l.ReadPosition >= len(l.Input) {
		l.Char = rune(0)
	} else {
		l.Char = l.Input[l.ReadPosition]
	}
	l.Position = l.ReadPosition
	l.ReadPosition += 1
}

func (l *Lexer) Skip_WhiteSpace() {
	for {
		if l.Char == ' ' || l.Char == '\n' {
			l.Read_Char()
		} else {
			break
		}
	}
}

func Is_Letter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func Is_Digit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (l *Lexer) Peak_Char() rune {
	if l.ReadPosition >= len(l.Input) {
		return rune(0)
	} else {
		return l.Input[l.ReadPosition]
	}
}

func (l *Lexer) Read_Identifier() string {
	var ident []rune
	for {
		if Is_Letter(l.Char) {
			ident = append(ident, l.Char)
			l.Read_Char()
		} else {
			break
		}
	}
	return string(ident)
}

func (l *Lexer) Read_Number() string {
	var num []rune
	for {
		if Is_Digit(l.Char) {
			num = append(num, l.Char)
			l.Read_Char()
		} else {
			break
		}
	}
	return string(num)
}

func (l *Lexer) Next_Token() token.Token {
	var tok token.Token
	switch l.Char {
	case '=':
		if l.Peak_Char() == '=' {
			l.Read_Char()
			tok = token.Token{
				Type:   token.EQUALS,
				String: "==",
			}
		} else {
			tok = token.NewToken(token.ASSIGN, l.Char)
		}
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.Char)
	case '+':
		tok = token.NewToken(token.PLUS, l.Char)
	case '(':
		tok = token.NewToken(token.LPAREN, l.Char)
	case ')':
		tok = token.NewToken(token.RPAREN, l.Char)
	case ',':
		tok = token.NewToken(token.COMMA, l.Char)
	case '{':
		tok = token.NewToken(token.LBRACE, l.Char)
	case '}':
		tok = token.NewToken(token.RBRACE, l.Char)
	case '!':
		if l.Peak_Char() == '=' {
			l.Read_Char()
			tok = token.Token{
				Type:   token.NOTEQUALS,
				String: "!=",
			}
		} else {
			tok = token.NewToken(token.BANG, l.Char)
		}
	case '-':
		tok = token.NewToken(token.MINUS, l.Char)
	case '/':
		tok = token.NewToken(token.SLASH, l.Char)
	case '<':
		tok = token.NewToken(token.LT, l.Char)
	case '>':
		tok = token.NewToken(token.GT, l.Char)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.Char)
	case rune(0):
		tok = token.NewToken(token.EOF, l.Char)
	default:
		if Is_Letter(l.Char) {
			str := l.Read_Identifier()
			Type := token.Lookup_Identifier(str)
			tok = token.Token{
				Type:   Type,
				String: str,
			}
		} else if Is_Digit(l.Char) {
			str := l.Read_Number()
			Type := token.TokenType(token.INTIGER)

			tok = token.Token{
				Type:   Type,
				String: str,
			}
		} else {
			tok = token.NewToken(token.ILLEGAL, l.Char)
		}

	}
	l.Read_Char()
	return tok
}
