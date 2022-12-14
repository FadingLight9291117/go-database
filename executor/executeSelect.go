package executor

import (
	"com.fadinglight/db/btree"
	"com.fadinglight/db/cursor"
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
	"fmt"
)

func ExecuteSelect(statement *types.Statement, t *table.Table) (ExecuteResult, error) {
	c := cursor.CreateStartCursor(t)
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
