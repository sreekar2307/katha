package main

import (
	"context"
	"fmt"
	"github.com/sreekar2307/katha/model/table"
)

func StartSeedCmd(ctx context.Context) error {
	d, err := GetDeps()
	if err != nil {
		return err
	}
	users := []table.User{
		{Email: "john.doe@email.com", Password: "password"},
		{Email: "john.doe@email1.com", Password: "password"},
		{Email: "john.doe@email2.com", Password: "password"},
		{Email: "john.doe@email3.com", Password: "password"},
		{Email: "john.doe@email4.com", Password: "password"},
	}
	for _, user := range users {
		_, err := d.UserService.AddUser(ctx, user)
		if err != nil {
			return fmt.Errorf("failed to seed user %s: %w", user.Email, err)
		}
	}
	return nil
}
