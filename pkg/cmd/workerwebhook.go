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

var workersWebhooksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve the webhook URL and shared token configured for a worker. Both fields\nare `null` when no webhook is configured. Only the worker creator can read the\nwebhook configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
	},
	Action:          handleWorkersWebhooksRetrieve,
	HideHelpCommand: true,
}

var workersWebhooksUpdate = cli.Command{
	Name:    "update",
	Usage:   "Set or replace the webhook URL for a worker. A fresh token is generated the\nfirst time a URL is set; subsequent updates keep the existing token. Pass\n`url: null` to clear the webhook (use the dedicated DELETE for the same effect).\nOnly the worker creator can update the webhook.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
		&requestflag.Flag[*string]{
			Name:     "url",
			Usage:    "HTTPS endpoint Handinger should POST to when a task finishes. Pass `null` to remove the webhook and clear its token.",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handleWorkersWebhooksUpdate,
	HideHelpCommand: true,
}

var workersWebhooksDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove the webhook from a worker. Both `url` and `token` are cleared and no\nfurther deliveries are attempted. Only the worker creator can delete the\nwebhook.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
	},
	Action:          handleWorkersWebhooksDelete,
	HideHelpCommand: true,
}

var workersWebhooksListExecutions = cli.Command{
	Name:    "list-executions",
	Usage:   "List recent webhook delivery attempts for a worker, newest first, paginated 50\nper page. Only the worker creator can read execution history.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number (1-indexed). Defaults to 1.",
			QueryPath: "page",
		},
	},
	Action:          handleWorkersWebhooksListExecutions,
	HideHelpCommand: true,
}

var workersWebhooksRegenerateToken = cli.Command{
	Name:    "regenerate-token",
	Usage:   "Issue a new shared token for the webhook, invalidating the previous one. The\nwebhook URL is preserved. Only the worker creator can regenerate the token.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "worker-id",
			Required:  true,
			PathParam: "workerId",
		},
	},
	Action:          handleWorkersWebhooksRegenerateToken,
	HideHelpCommand: true,
}

func handleWorkersWebhooksRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Workers.Webhooks.Get(ctx, cmd.Value("worker-id").(string), options...)
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
		Title:          "workers:webhooks retrieve",
		Transform:      transform,
	})
}

func handleWorkersWebhooksUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := handinger.WorkerWebhookUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workers.Webhooks.Update(
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
		Title:          "workers:webhooks update",
		Transform:      transform,
	})
}

func handleWorkersWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Workers.Webhooks.Delete(ctx, cmd.Value("worker-id").(string), options...)
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
		Title:          "workers:webhooks delete",
		Transform:      transform,
	})
}

func handleWorkersWebhooksListExecutions(ctx context.Context, cmd *cli.Command) error {
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

	params := handinger.WorkerWebhookListExecutionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workers.Webhooks.ListExecutions(
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
		Title:          "workers:webhooks list-executions",
		Transform:      transform,
	})
}

func handleWorkersWebhooksRegenerateToken(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Workers.Webhooks.RegenerateToken(ctx, cmd.Value("worker-id").(string), options...)
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
		Title:          "workers:webhooks regenerate-token",
		Transform:      transform,
	})
}
