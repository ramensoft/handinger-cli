// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/ramensoft/handinger-cli/internal/apiquery"
	"github.com/ramensoft/handinger-cli/internal/requestflag"
	"github.com/ramensoft/handinger-go"
	"github.com/ramensoft/handinger-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var tasksCreate = cli.Command{
	Name:    "create",
	Usage:   "Run a new task against an existing worker. Send a `taskId` of a prior task to\nadd a follow-up turn instead of starting a fresh task. Send\n`multipart/form-data` to attach files; the bytes are bootstrapped into the\nworker's workspace before the task starts.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "input",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[string]{
			Name:     "budget",
			Usage:    `Allowed values: "low", "standard", "high", "unlimited".`,
			Default:  "standard",
			BodyPath: "budget",
		},
		&requestflag.Flag[bool]{
			Name:     "stream",
			Default:  false,
			BodyPath: "stream",
		},
		&requestflag.Flag[string]{
			Name:     "task-id",
			Usage:    "Optional client-provided task id. Reuse this id to add turns to an existing task.",
			BodyPath: "taskId",
		},
		&requestflag.Flag[string]{
			Name:     "worker-id",
			Usage:    "Worker id the task belongs to. If omitted, a new worker is created on-the-fly using the input as instructions.",
			BodyPath: "workerId",
		},
	},
	Action:          handleTasksCreate,
	HideHelpCommand: true,
}

var tasksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve a single task and its individual turns.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "task-id",
			Required:  true,
			PathParam: "taskId",
		},
	},
	Action:          handleTasksRetrieve,
	HideHelpCommand: true,
}

var tasksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Archive a task so it stops appearing in `GET /tasks` results. Turns and files\nare retained for audit purposes. Only the worker creator can archive a task.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "task-id",
			Required:  true,
			PathParam: "taskId",
		},
	},
	Action:          handleTasksDelete,
	HideHelpCommand: true,
}

func handleTasksCreate(ctx context.Context, cmd *cli.Command) error {
	client := handinger.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := handinger.TaskNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tasks.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tasks create",
		Transform:      transform,
	})
}

func handleTasksRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := handinger.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("task-id") && len(unusedArgs) > 0 {
		cmd.Set("task-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tasks.Get(ctx, cmd.Value("task-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tasks retrieve",
		Transform:      transform,
	})
}

func handleTasksDelete(ctx context.Context, cmd *cli.Command) error {
	client := handinger.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("task-id") && len(unusedArgs) > 0 {
		cmd.Set("task-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Tasks.Delete(ctx, cmd.Value("task-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "tasks delete",
		Transform:      transform,
	})
}
