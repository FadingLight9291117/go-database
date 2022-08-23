package executor

import (
	"com.fadinglight/db/table"
	"com.fadinglight/db/types"
)

func ExecuteSelect(statement *types.Statement, t *table.Table) (ExecuteResult, error) {
	for _, v := range t.Select() {
		v.Print()
	}
	return EXECUTE_SUCESS, nil
}
