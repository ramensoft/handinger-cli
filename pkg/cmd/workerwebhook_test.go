// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/ramensoft/handinger-cli/internal/mocktest"
)

func TestWorkersWebhooksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers:webhooks", "retrieve",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		)
	})
}

func TestWorkersWebhooksUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers:webhooks", "update",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
			"--url", "https://example.com/handinger-webhook",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("url: https://example.com/handinger-webhook")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workers:webhooks", "update",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		)
	})
}

func TestWorkersWebhooksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers:webhooks", "delete",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		)
	})
}

func TestWorkersWebhooksListExecutions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers:webhooks", "list-executions",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
			"--page", "1",
		)
	})
}

func TestWorkersWebhooksRegenerateToken(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers:webhooks", "regenerate-token",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		)
	})
}
