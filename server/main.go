package main

import (
	"context"
	"os"

	"vsoutlook.com/vsoutlook/cmd"
	"vsoutlook.com/vsoutlook/infra/config"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "serve")
	}

	config.SetupEnv()
	ctx := context.Background()
	os.Exit(int(cmd.Execute(ctx)))
}
