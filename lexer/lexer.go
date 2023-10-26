package lexer

import (
	"monkey-lang/helpers"
	"monkey-lang/token"
	"unicode"
)

type Lexer struct {
	Input        []rune
	Position     int
	ReadPosition int
	Char         rune
	Pos          helpers.Pos
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Input:        []rune(input),
		Position:     0,
		ReadPosition: 0,
		Char:         ' ',
		Pos: helpers.Pos{
			Column: 0,
			Line:   1,
		},
	}
}

func (l *Lexer) Read_Char() {
	if l.Char == '\n' {
		l.Pos.Line++
		l.Pos.Column = 0
	} else {
		l.Pos.Column++
	}
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
		if l.Char == ' ' || l.Char == '\n' || l.Char == '\r' {
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
		if Is_Letter(l.Char) || Is_Digit(l.Char) {
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
func (l *Lexer) Read_String() string {
	var str []rune
	if l.Char == '"' {
		l.Read_Char()
	} else {
		return "Invalid String"
	}

	for {
		if l.Char == '"' {
			l.Read_Char()
			break
		} else if l.Char == 0 {
			return "Unclosed String"
		}
		str = append(str, l.Char)
		l.Read_Char()
	}

	return string(str)
}

func (l *Lexer) Next_Token() token.Token {
	var tok token.Token
	l.Skip_WhiteSpace()
	switch l.Char {
	case '=':
		if l.Peak_Char() == '=' {
			l.Read_Char()
			tok = token.Token{
				Type:   token.EQUALS,
				String: "==",
			}
		} else {
			tok = token.NewToken(token.ASSIGN, l.Char, l.Pos)
		}
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.Char, l.Pos)
	case '+':
		tok = token.NewToken(token.PLUS, l.Char, l.Pos)
	case '(':
		tok = token.NewToken(token.LPAREN, l.Char, l.Pos)
	case ')':
		tok = token.NewToken(token.RPAREN, l.Char, l.Pos)
	case ',':
		tok = token.NewToken(token.COMMA, l.Char, l.Pos)
	case '{':
		tok = token.NewToken(token.LBRACE, l.Char, l.Pos)
	case '}':
		tok = token.NewToken(token.RBRACE, l.Char, l.Pos)
	case '!':
		if l.Peak_Char() == '=' {
			l.Read_Char()
			tok = token.Token{
				Type:   token.NOTEQUALS,
				String: "!=",
				Pos:    l.Pos,
			}
		} else {
			tok = token.NewToken(token.BANG, l.Char, l.Pos)
		}
	case '-':
		tok = token.NewToken(token.MINUS, l.Char, l.Pos)
	case '/':
		tok = token.NewToken(token.SLASH, l.Char, l.Pos)
	case '<':
		tok = token.NewToken(token.LT, l.Char, l.Pos)
	case '>':
		tok = token.NewToken(token.GT, l.Char, l.Pos)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.Char, l.Pos)
	case '"':
		str := l.Read_String()
		l.ReadPosition -= 1
		Type := token.TokenType(token.STRING)
		tok = token.Token{
			Type:   Type,
			String: str,
			Pos:    l.Pos,
		}
	case rune(0):
		tok = token.NewToken(token.EOF, l.Char, l.Pos)
	default:
		if Is_Letter(l.Char) {
			str := l.Read_Identifier()
			l.ReadPosition -= 1
			l.Pos.Column -= 1
			Type := token.Lookup_Identifier(str)
			tok = token.Token{
				Type:   Type,
				String: str,
				Pos:    l.Pos,
			}
		} else if Is_Digit(l.Char) {
			str := l.Read_Number()
			l.ReadPosition -= 1
			l.Pos.Column -= 1
			Type := token.TokenType(token.INTIGER)

			tok = token.Token{
				Type:   Type,
				String: str,
				Pos:    l.Pos,
			}
		} else {
			tok = token.NewToken(token.ILLEGAL, l.Char, l.Pos)
		}

	}
	l.Read_Char()
	return tok
}
