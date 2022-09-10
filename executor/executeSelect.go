package executor

import (
	"com.fadinglight/db/BTree"
	"com.fadinglight/db/cursor"
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"fmt"
)

func ExecuteSelect(statement *types.Statement, t *table.Table) (ExecuteResult, error) {
	c := cursor.CreateCursor(t, false)
	var r *BTree.Row
	i := 0
	for !c.IsEnd() {
		r = c.Next()
		r.Print()
		i++
	}
	fmt.Printf("%d rows selected.\n", i)
	return EXECUTE_SUCESS, nil
}
