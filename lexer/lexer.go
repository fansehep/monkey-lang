package lexer

import (
	"github.com/fansehep/monkey-lang/token"
)

type Lexer struct {
	input           string
	position        int
	readPosition    int
	currentReadChar byte
}

func New(intput string) *Lexer {
	l := &Lexer{
		input: intput,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentReadChar = 0
	} else {
		l.currentReadChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(ty token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    ty,
		Literal: string(ch),
	}
}

func (l *Lexer) skipWhiteSpace() {
	if l.currentReadChar == ' ' ||
		l.currentReadChar == '\t' ||
		l.currentReadChar == '\n' ||
		l.currentReadChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	switch l.currentReadChar {
	case '=':
		tok = newToken(token.ASSIGN, l.currentReadChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentReadChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentReadChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentReadChar)
	case ',':
		tok = newToken(token.COMMA, l.currentReadChar)
	case '+':
		tok = newToken(token.PLUS, l.currentReadChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentReadChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentReadChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currentReadChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.currentReadChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.LLLEGAL, l.currentReadChar)
		}
	}
	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

// 读取标识符
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentReadChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 读取用户的数字
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentReadChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
