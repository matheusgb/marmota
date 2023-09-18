package lexer

import "github.com/matheusgb/marmota/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tokenVar token.Token

	switch lexer.character {
	case '=':
		tokenVar = newToken(token.ASSIGN, lexer.character)
	case ';':
		tokenVar = newToken(token.SEMICOLON, lexer.character)
	case '(':
		tokenVar = newToken(token.LPAREN, lexer.character)
	case ')':
		tokenVar = newToken(token.RPAREN, lexer.character)
	case ',':
		tokenVar = newToken(token.COMMA, lexer.character)
	case '+':
		tokenVar = newToken(token.PLUS, lexer.character)
	case '{':
		tokenVar = newToken(token.LBRACE, lexer.character)
	case '}':
		tokenVar = newToken(token.RBRACE, lexer.character)
	case 0:
		tokenVar.Literal = ""
		tokenVar.Type = token.EOF
	default:
		if isLetter(lexer.character) {
			tokenVar.Literal = lexer.readIdentifier()
			return tokenVar
		} else {
			tokenVar = newToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readChar()
	return tokenVar
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.character) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}
