package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"vsoutlook.com/vsoutlook/models"
)

type dbCmd struct{}

func (cmd *dbCmd) Name() string {
	return "db"
}

func (cmd *dbCmd) Synopsis() string {
	return "db commands"
}

func (cmd *dbCmd) Usage() string {
	return fmt.Sprintf("run with: %s migrate\n", cmd.Name())
}

func (cmd *dbCmd) SetFlags(fs *flag.FlagSet) {
}

func (cmd *dbCmd) Execute(_ context.Context, fs *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	args := fs.Args()
	if len(args) > 0 && args[0] == "migrate" {
		models.ReadyDB()
		models.Migrate("")
	}
	if len(args) > 0 && args[0] == "migrate-testdb" {
		models.ReadyMaintainTestDB()
		models.Migrate("test")
	}
	return subcommands.ExitSuccess
}
