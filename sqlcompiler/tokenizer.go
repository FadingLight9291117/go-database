package sqlcompiler

import (
	"strings"
)

/**
 * 1. 词法分析；
 * 2. 语法分析，构建AST - abstract syntax tree;
 */

type SQL struct {
	sql    string
	offset int
}

func NewSQL(sql string) *SQL { return new(SQL).init(sql) }

func (sql *SQL) init(s string) *SQL {
	sql.offset = 0
	sql.sql = s
	return sql
}

func (sql *SQL) NextToken() (*Token, error) {
	sql.skipWhiteSpace()
	if sql.isIdentifierBegin() {
		return sql.scanIdentifier(), nil
	} else if sql.isSymbol() {
		return sql.scanSymbol(), nil
	} else if sql.IsEnd() {
		return sql.end(), nil
	}
	return nil, nil
}

func (sql *SQL) skipWhiteSpace() {
	length := 0
	for length = sql.offset; length < len(sql.sql); length++ {
		if sql.sql[length] != ' ' {
			break
		}
	}
	sql.offset = length
}

func (sql *SQL) IsEnd() bool {
	return sql.offset == len(sql.sql)
}

func (sql *SQL) isIdentifierBegin() bool {
	return isAlphabet(sql.sql[sql.offset])
}

func (sql *SQL) isSymbol() bool {
	for _, symbol := range tokenTypes[SYMBOL] {
		if sql.sql[sql.offset] == symbol[0] {
			return true
		}
	}
	return false
}

func (sql *SQL) end() *Token {
	return &Token{
		Type:        ASSIST,
		Literals:    "",
		EndPosition: sql.offset,
	}
}

func (sql *SQL) scanIdentifier() *Token {
	length := 0
	for sql.offset+length < len(sql.sql) && isAlphabet(sql.sql[sql.offset+length]) {
		length++
	}
	literals := sql.sql[sql.offset : sql.offset+length]
	tokenType := findTokenType(literals, IDENTIFIER)
	token := &Token{
		Type:        tokenType,
		Literals:    literals,
		EndPosition: length,
	}
	sql.offset += length
	return token
}

func (sql *SQL) scanSymbol() *Token {
	literals := sql.sql[sql.offset : sql.offset+1]
	sql.offset++
	return &Token{
		Type:        SYMBOL,
		Literals:    literals,
		EndPosition: sql.offset,
	}
}

type Token struct {
	Type        TokenType
	Literals    string
	EndPosition int // not contain
}

func findTokenType(literals string, literalsType Literals) TokenType {
	for _, keyword := range tokenTypes[KEYWORD] {
		if strings.EqualFold(literals, keyword) {
			return KEYWORD
		}
	}
	return LITERALS
}

// Tokenizer 词法分析
func Tokenizer(sql string) ([]Token, error) {
	return nil, nil
}

func isAlphabet(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}
