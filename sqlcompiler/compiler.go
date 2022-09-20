package sqlcompiler

import (
	"errors"
	"strings"
	"unicode"
)

type SelectStatement struct {
	Table  string
	Fields []string
}
type InsertStatement struct {
}

// Tokenizer_ 词法分析
func Tokenizer_(sql string) (interface{}, error) {
	sql = strings.TrimSpace(sql)
	switch strings.ToUpper(begin(sql)) {
	case "SELECT":
		return SelectTokenizer(sql)
	case "INSERT":
		return InsertTokenizer(sql)
	default:
		return nil, errors.New("not recognized SQL statement")
	}
}

func InsertTokenizer(sql string) (*InsertStatement, error) {
	return &InsertStatement{}, nil
	//return nil, errors.New("not recognized insert statement")
}

func SelectTokenizer(sql string) (*SelectStatement, error) {
	// "select * from table"
	words := strings.Fields(sql)[1:]

	var fieldStrs []string
	var tableStr string
	for i, word := range words {
		if strings.EqualFold(word, "FROM") {
			fieldStrs = words[:i]
			tableStr = words[i+1]
			if i+1 < len(words)-1 {
				return nil, errors.New("select statement error")
			}
			break
		}
	}

	fieldStr := strings.Join(fieldStrs, "")
	fieldStrs = strings.FieldsFunc(fieldStr, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '*'
	})

	return &SelectStatement{Table: tableStr, Fields: fieldStrs}, nil
	//return nil, errors.New("not recognized select statement")
}

func begin(sql string) string {
	firstSpaceIndex := strings.Index(sql, " ")
	if firstSpaceIndex > -1 {
		return sql[:firstSpaceIndex]
	}
	return sql
}
