package rputils

import "github.com/jmoiron/sqlx"

// Transaction implements commit or rollback in sql transactions
func Transaction(db *sqlx.DB, fn func(*sqlx.Tx) error) (err error) {
	var tx *sqlx.Tx

	tx, err = db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return fn(tx)
}

func CloseDBTransaction(resp *sqlx.Rows, err error) error {
	if err != nil {
		return err
	}

	if err := resp.Close(); err != nil {
		return err
	}

	return err
}

// CloseTransactionReturningID
func CloseDBTransactionReturningID(resp *sqlx.Rows, err error) (int64, error) {
	if err != nil {
		return 0, err
	}

	var id int64
	resp.Next()

	if err := resp.Scan(&id); err != nil {
		return 0, err
	}

	if err := resp.Close(); err != nil {
		return 0, err
	}

	return id, err
}
