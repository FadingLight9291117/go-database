package main

import (
	"fmt"
	"log"

	"com.fadinglight/db/sqlcompiler"
)

func main() {
	sql := "SELECT id FROM table"
	SQL := sqlcompiler.NewSQL(sql)

	for !SQL.IsEnd() {
		token, err := SQL.NextToken()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%v\n", token)
	}
}
