// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/Ramensoft/handinger-cli/internal/apiquery"
	"github.com/Ramensoft/handinger-cli/internal/requestflag"
	"github.com/Ramensoft/handinger-go"
	"github.com/Ramensoft/handinger-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var workersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new worker. The worker is a reusable agent template; tasks are runs\nagainst this template. Use `POST /tasks` to actually run the agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
		&requestflag.Flag[string]{
			Name:     "instructions",
			Usage:    "Persistent system prompt the worker uses for every task it runs.",
			BodyPath: "instructions",
		},
		&requestflag.Flag[string]{
			Name:     "visibility",
			Usage:    "`public` (default) is visible to all org members. `private` is only visible to invited members.",
			BodyPath: "visibility",
		},
	},
	Action:          handleWorkersCreate,
	HideHelpCommand: true,
}

var workersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve the current worker state and messages from its most recent task.\nReturns a JSON worker object by default, or a server-sent event stream when\n`stream=true`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
		&requestflag.Flag[string]{
			Name:      "stream",
			Usage:     `Set to "true" to receive a server-sent event stream that replays all stored messages and then continues with live chunks from the active task (if any) before closing.`,
			QueryPath: "stream",
		},
	},
	Action:          handleWorkersRetrieve,
	HideHelpCommand: true,
}

var workersRetrieveEmail = cli.Command{
	Name:    "retrieve-email",
	Usage:   "Retrieve the inbound email address for a worker.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
	},
	Action:          handleWorkersRetrieveEmail,
	HideHelpCommand: true,
}

func handleWorkersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := handinger.WorkerNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workers.New(ctx, params, options...)
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
		Title:          "workers create",
		Transform:      transform,
	})
}

func handleWorkersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := handinger.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("worker-id") && len(unusedArgs) > 0 {
		cmd.Set("worker-id", unusedArgs[0])
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

	params := handinger.WorkerGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workers.Get(
		ctx,
		cmd.Value("worker-id").(string),
		params,
		options...,
	)
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
		Title:          "workers retrieve",
		Transform:      transform,
	})
}

func handleWorkersRetrieveEmail(ctx context.Context, cmd *cli.Command) error {
	client := handinger.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("worker-id") && len(unusedArgs) > 0 {
		cmd.Set("worker-id", unusedArgs[0])
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
	_, err = client.Workers.GetEmail(ctx, cmd.Value("worker-id").(string), options...)
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
		Title:          "workers retrieve-email",
		Transform:      transform,
	})
}
