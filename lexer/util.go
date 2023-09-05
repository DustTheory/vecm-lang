package lexer

import "github.com/0xM-D/interpreter/token"

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition++
	l.coln++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
		if l.ch == '\n' {
			l.linen++
			l.coln = 0
		}
	}
}

func (l *Lexer) newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal, Linen: l.linen, Coln: l.coln}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() (token.TokenType, string) {
	position := l.position
	tokenType := token.INT

	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' && isDigit(l.peekChar()) {
		tokenType = token.FLOAT64
		l.readChar()

		for isDigit(l.ch) {
			l.readChar()
		}
	}

	if l.ch == 'f' {
		tokenType = token.FLOAT32
		l.readChar()
	}
	return tokenType, l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

type TokenMapping struct {
	byte
	token.TokenType
}

func (l *Lexer) getTokenWithPeek(defaultToken token.TokenType, tokenMappings ...TokenMapping) token.Token {
	for _, tokenMapping := range tokenMappings {
		if l.peekChar() == tokenMapping.byte {
			ch := l.ch
			l.readChar()
			return token.Token{Type: tokenMapping.TokenType, Literal: string(ch) + string(l.ch)}
		}
	}

	return l.newToken(defaultToken, string(l.ch))
}
