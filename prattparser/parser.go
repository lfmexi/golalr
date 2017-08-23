// Copyright 2017 Luis Fernando Morales

/***
    This library is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This library is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this library.  If not, see <http://www.gnu.org/licenses/>.
***/

// Package prattparser gives the core of a Top Down Operator precedence
// parser. You can extend it by implementing the Parselets interfaces
// defined here.
package prattparser

import (
	"fmt"

	"github.com/lfmexi/golalr/symbols"
	"github.com/lfmexi/golalr/types"
)

// Parser is the struct that represents a Top Down Operator Precedence Parser
type Parser struct {
	tokenIterator   symbols.TokenIterator
	tokensBuffer    []symbols.Token
	prefixParselets map[types.TokenType]PrefixParselet
	infixParselets  map[types.TokenType]InfixParselet
}

// NewParser is a constructor of Parser values. It takes an implementation of the TokenIterator interface
func NewParser(iterator symbols.TokenIterator) Parser {
	return Parser{
		iterator,
		make([]symbols.Token, 0),
		make(map[types.TokenType]PrefixParselet),
		make(map[types.TokenType]InfixParselet),
	}
}

// RegisterPrefixParselet takes a token type and a PrefixParselet implementation and
// stores them into the inner prefix parselets
func (p *Parser) RegisterPrefixParselet(tType types.TokenType, parselet PrefixParselet) {
	p.prefixParselets[tType] = parselet
}

// RegisterInfixParselet takes a token type and a InfixParselet implementation and
// stores them into the inner infix parselets
func (p *Parser) RegisterInfixParselet(tType types.TokenType, parselet InfixParselet) {
	p.infixParselets[tType] = parselet
}

// LookNext gets the next token from the tokens buffer without removing it
func (p *Parser) LookNext() symbols.Token {
	for len(p.tokensBuffer) == 0 {
		p.tokensBuffer = append(p.tokensBuffer, p.tokenIterator.Next())
	}
	return p.tokensBuffer[0]
}

// Consume gets the next token from the tokens buffer and removes it from there
func (p *Parser) Consume(tType types.TokenType) (symbols.Token, error) {
	token := p.LookNext()

	if tType != nil && token.TokenType() != tType {
		return nil, NewParseError(token.Text(), nil)
	}

	p.tokensBuffer = append(p.tokensBuffer[:0], p.tokensBuffer[1:]...)
	return token, nil
}

// GetInfixParselet gets the infix parselet registered with the given TokenType
func (p *Parser) GetInfixParselet(tType types.TokenType) InfixParselet {
	return p.infixParselets[tType]
}

// GetPrefixParselet gets the infix parselet registered with the given TokenType
func (p *Parser) GetPrefixParselet(tType types.TokenType) PrefixParselet {
	return p.prefixParselets[tType]
}

// Parse generates a new Expression by consuming the tokens buffer with the given precedence
// It could return an Expression or an Error on parsing an input
func (p *Parser) Parse(precedence int) (Expression, error) {
	token, _ := p.Consume(nil)
	prefix := p.prefixParselets[token.TokenType()]
	if prefix == nil {
		return nil, NewParseError(token.Text(), nil)
	}

	left, err := prefix.Parse(p, token)
	if err != nil {
		return nil, err
	}

	nextToken := p.LookNext()
	for precedence < nextToken.TokenType().GetPrecedence() {
		token, _ = p.Consume(nil)
		infix := p.infixParselets[token.TokenType()]
		if infix == nil {
			infixErr := fmt.Errorf("Parse error: no infix operator for token %v", token.Text())
			return nil, NewParseError(token.Text(), infixErr)
		}
		left, err = infix.Parse(p, left, token)
		if err != nil {
			return nil, err
		}
		nextToken = p.LookNext()
	}
	return left, nil
}
