package sqlcompiler

import (
	"flag"
	"fmt"
	"log"
)

var sql = flag.String("sql", "", "")

func main() {
	flag.Parse()
	log.Println(*sql)
	statement, err := Tokenizer_(*sql)

	if err != nil {
		log.Fatalln(err)
	}
	switch statement.(type) {
	case *SelectStatement:
		selectStatement := statement.(*SelectStatement)
		fmt.Printf("SelectStatement(\n\ttable=%s,\n\tfield=%v,\n)\n", selectStatement.Table, selectStatement.Fields)
	case *InsertStatement:
		//insertStatement := statement.(*InsertStatement)
		//fmt.Printf("InsertStatement(\n\ttable=%s,\n\tfield=%v,\n)\n", insertStatement.Table, insertStatement.Fields)
	}
}
