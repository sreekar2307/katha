package main

import (
	"context"
	"fmt"
	"os"
)

func StartGenerateAuthToken(ctx context.Context) error {
	d, err := GetDeps()
	if err != nil {
		return err
	}
	if len(os.Args) < 4 {
		return fmt.Errorf("usage: %s <command> <email> <password>", os.Args[0])
	}
	emailID := os.Args[2]
	password := os.Args[3]
	_, token, err := d.UserService.Authenticate(ctx, emailID, password)
	if err != nil {
		return fmt.Errorf("failed to authenticate user %s: %w", emailID, err)
	}
	_, _ = fmt.Fprintf(os.Stdout, "Generated token for user %s: %s\n", emailID, token)
	return nil
}
