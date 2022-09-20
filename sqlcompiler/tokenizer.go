package sqlcompiler

import (
	"errors"
	"log"
)

/**
 * 1. 词法分析；
 * 2. 语法分析，构建AST - abstract syntax tree;
 */

type SQL struct {
	sql string
	c   int
}

func NewSQL(sql string) *SQL { return new(SQL).init() }

func (sql *SQL) init() *SQL {
	sql.c = 0
	return sql
}

func (sql *SQL) nextToken() (*Token, error) {
	for sql.sql[sql.c] == ' ' {
		sql.c++
	}

	return nil, errors.New("no Token found")
}

type Token struct {
	Type        TokenType
	Literals    string
	EndPosition string
}

// Tokenizer 词法分析
func Tokenizer(sql string) ([]Token, error) {
	return nil, nil
}

func main() {
	sql := NewSQL("select * from users")
	_, err := sql.nextToken()
	if err != nil {
		log.Fatalln(err)
	}
}
