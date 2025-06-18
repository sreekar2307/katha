package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	if len(os.Args) < 2 {
		panic("base command is required")
	}
	switch os.Args[1] {
	case "token":
		err := StartGenerateAuthToken(ctx)
		if err != nil {
			panic(err.Error())
		}
	case "seed":
		err := StartSeedCmd(ctx)
		if err != nil {
			panic(err.Error())
		}
	case "migrate":
		err := StartMigrateCmd(ctx)
		if err != nil {
			panic(err.Error())
		}
	case "http":
		err := StartHTTPCmd(ctx)
		if err != nil {
			panic(err.Error())
		}
	}
}
