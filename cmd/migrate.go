package main

import (
	"context"
	"fmt"
	"github.com/sreekar2307/katha/model/table"
)

func StartMigrateCmd(ctx context.Context) error {
	d, err := GetDeps()
	if err != nil {
		return fmt.Errorf("failed to get dependencies: %w", err)
	}
	return d.PrimaryDB.WithContext(ctx).AutoMigrate(table.All...)
}
