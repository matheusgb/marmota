package lexer

import (
	"github.com/matheusgb/marmota/token"
)

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

	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		if lexer.peekChar() == '=' {
			character := lexer.character
			lexer.readChar()
			tokenVar = token.Token{Type: token.EQ, Literal: string(character) + string(lexer.character)}
		} else if lexer.peekChar() == '>' {
			character := lexer.character
			lexer.readChar()
			tokenVar = token.Token{Type: token.ARROW_FUNCTION, Literal: string(character) + string(lexer.character)}
		} else {
			tokenVar = newToken(token.ASSIGN, lexer.character)
		}
	case '-':
		tokenVar = newToken(token.MINUS, lexer.character)
	case '!':
		if lexer.peekChar() == '=' {
			character := lexer.character
			lexer.readChar()
			tokenVar = token.Token{Type: token.NOT_EQ, Literal: string(character) + string(lexer.character)}
		} else {
			tokenVar = newToken(token.BANG, lexer.character)
		}
	case '/':
		tokenVar = newToken(token.SLASH, lexer.character)
	case '*':
		tokenVar = newToken(token.ASTERISK, lexer.character)
	case '<':
		tokenVar = newToken(token.LT, lexer.character)
	case '>':
		tokenVar = newToken(token.GT, lexer.character)
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
			tokenVar.Type = LookupIdentifier(tokenVar.Literal)
			return tokenVar
		} else if isDigit(lexer.character) {
			tokenVar.Type = token.INT
			tokenVar.Literal = lexer.readNumber()
			return tokenVar
		} else {
			tokenVar = newToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readChar()
	return tokenVar
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.character) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.readChar()
	}
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

func LookupIdentifier(identifier string) token.TokenType {
	if tokenType, ok := token.Keywords[identifier]; ok {
		return tokenType
	}
	return token.IDENT
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}
