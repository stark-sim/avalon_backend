package db

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/pkg/ent"
)

// WithTx wrap transaction according to doc of Ent
// https://entgo.io/docs/transactions/
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		logrus.Errorf("begin tx: %v", err)
		return err
	}
	defer func() {
		// check if anything went wrong
		if v := recover(); v != nil {
			if err = tx.Rollback(); err != nil {
				logrus.Errorf("tx rollback: %v", err)
				return
			}
			panic(v)
		}
	}()
	// run the process
	if err = fn(tx); err != nil {
		// process go wrong
		if rErr := tx.Rollback(); rErr != nil {
			// rollback go wrong
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rErr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		// commit go wrong
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
