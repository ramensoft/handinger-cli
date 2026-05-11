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

var workersSchedulesCreate = cli.Command{
	Name:    "create",
	Usage:   "Schedule a worker instruction for a future or recurring run.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
		&requestflag.Flag[string]{
			Name:     "input",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "when",
			Required: true,
			BodyPath: "when",
		},
		&requestflag.Flag[string]{
			Name:     "budget",
			Usage:    `Allowed values: "low", "standard", "high", "unlimited".`,
			Default:  "standard",
			BodyPath: "budget",
		},
	},
	Action:          handleWorkersSchedulesCreate,
	HideHelpCommand: true,
}

var workersSchedulesList = cli.Command{
	Name:    "list",
	Usage:   "List scheduled tasks for a worker.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
	},
	Action:          handleWorkersSchedulesList,
	HideHelpCommand: true,
}

var workersSchedulesCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancel a scheduled task for a worker.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
		&requestflag.Flag[string]{
			Name:      "schedule-id",
			Required:  true,
			PathParam: "scheduleId",
		},
	},
	Action:          handleWorkersSchedulesCancel,
	HideHelpCommand: true,
}

func handleWorkersSchedulesCreate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := handinger.WorkerScheduleNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workers.Schedules.New(
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
		Title:          "workers:schedules create",
		Transform:      transform,
	})
}

func handleWorkersSchedulesList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Workers.Schedules.List(ctx, cmd.Value("worker-id").(string), options...)
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
		Title:          "workers:schedules list",
		Transform:      transform,
	})
}

func handleWorkersSchedulesCancel(ctx context.Context, cmd *cli.Command) error {
	client := handinger.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("schedule-id") && len(unusedArgs) > 0 {
		cmd.Set("schedule-id", unusedArgs[0])
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

	params := handinger.WorkerScheduleCancelParams{
		WorkerID: cmd.Value("worker-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workers.Schedules.Cancel(
		ctx,
		cmd.Value("schedule-id").(string),
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
		Title:          "workers:schedules cancel",
		Transform:      transform,
	})
}
