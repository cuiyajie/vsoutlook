package cmd

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func Execute(ctx context.Context) subcommands.ExitStatus {
	subcommands.Register(&dbCmd{}, "")
	subcommands.Register(&serveCmd{}, "")
	subcommands.Register(&taskCmd{}, "")
	flag.Parse()
	return subcommands.Execute(ctx)
}
