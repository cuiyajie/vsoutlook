package cmd

import (
	"context"
	"flag"
	"log"

	"github.com/google/subcommands"
	"vsoutlook.com/vsoutlook/infra/redis"
	"vsoutlook.com/vsoutlook/models"
)

type taskCmd struct {
	action string
	target string
}

func (cmd *taskCmd) Name() string {
	return "task"
}

func (cmd *taskCmd) Synopsis() string {
	return "common tasks"
}

func (cmd *taskCmd) Usage() string {
	return `
	run with:
		task -action convert -target <FILE_ID> # re-start processing file
		task -action check -target <FILE_ID>   # check file processing status
	`
}

func (cmd *taskCmd) SetFlags(fs *flag.FlagSet) {
	// fs.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
	fs.StringVar(&cmd.action, "action", "convert", "Action to be performed")
	fs.StringVar(&cmd.target, "target", "", "file ID")
}

func (cmd *taskCmd) Execute(_ context.Context, fs *flag.FlagSet, _ ...any) subcommands.ExitStatus {

	if cmd.action == "check" {
		if len(cmd.target) == 0 {
			log.Println("no target specified")
			return subcommands.ExitStatus(1)
		}
		redis.ConnectRedis()
		models.ReadyDB()

		return subcommands.ExitSuccess
	}

	return subcommands.ExitSuccess
}
