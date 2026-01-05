package lexer

import (
	"slices"
	"unicode"

	"github.com/chaitanya-Uike/lemon/token"
)

type Lexer struct {
	input     string
	pos       int
	readPos   int
	ch        byte
	prevToken *token.Token
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	l.pos = l.readPos
	if l.readPos >= len(l.input) {
		l.ch = 0
		return
	}
	l.ch = l.input[l.readPos]
	l.readPos++
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '\n':
		if l.shouldInsertSemicolon() {
			tok = newToken(token.SEMICOLON, ';')
		} else {
			l.readChar()
			return l.NextToken()
		}
	case 0:
		if l.shouldInsertSemicolon() {
			tok = newToken(token.SEMICOLON, ';')
		} else {
			tok.Literal = ""
			tok.Type = token.EOF
		}

	default:
		if isLetter(l.ch) {
			tok.Literal, tok.Type = l.readIdentifier()
			l.prevToken = &tok
			return tok
		} else if unicode.IsDigit(rune(l.ch)) {
			tok.Literal, tok.Type = l.readNumber()
			l.prevToken = &tok
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	l.prevToken = &tok
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func (l *Lexer) readIdentifier() (string, token.TokenType) {
	pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	ident := l.input[pos:l.pos]
	return ident, token.LookupIdent(ident)
}

func (l *Lexer) readNumber() (string, token.TokenType) {
	pos := l.pos
	isFloat := false

	for unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}

	if l.ch == '.' {
		isFloat = true
		l.readChar()

		for unicode.IsDigit(rune(l.ch)) {
			l.readChar()
		}
	}

	if isFloat {
		return l.input[pos:l.pos], token.FLOAT
	}
	return l.input[pos:l.pos], token.INT
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) shouldInsertSemicolon() bool {
	if l.prevToken == nil {
		return false
	}

	closingTypes := []token.TokenType{
		token.IDENT,

		token.INT,
		token.FLOAT,
		token.TRUE,
		token.FALSE,

		token.RETURN,

		token.RPAREN,
	}

	return slices.Contains(closingTypes, l.prevToken.Type)
}
