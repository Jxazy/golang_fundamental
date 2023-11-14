package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		CheckError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		CheckError(errorCommit)
	}
}
