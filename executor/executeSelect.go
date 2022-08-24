package executor

import (
	"com.fadinglight/db/cursor"
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
)

func ExecuteSelect(statement *types.Statement, t *table.Table) (ExecuteResult, error) {
	c := cursor.CreateCursor(t, false)
	var r *table.Row
	for !c.IsEnd() {
		r = c.Next()
		r.Print()
	}
	return EXECUTE_SUCESS, nil
}
