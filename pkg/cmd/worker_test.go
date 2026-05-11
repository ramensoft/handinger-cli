// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/ramensoft/handinger-cli/internal/mocktest"
)

func TestWorkersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers", "create",
			"--instructions", "You are a brand voice analyzer. Read the input text and report whether it matches Acme's playful, plain-spoken house style. Quote specific phrases.",
			"--output-schema", "{type: bar, required: bar, properties: bar}",
			"--prompt", "A worker that fact-checks short claims and returns a verdict with citations.",
			"--summary", "Audits copy against the Acme brand voice guide.",
			"--title", "Brand voice analyzer",
			"--visibility", "public",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: >-\n" +
			"  You are a brand voice analyzer. Read the input text and report whether it\n" +
			"  matches Acme's playful, plain-spoken house style. Quote specific phrases.\n" +
			"outputSchema:\n" +
			"  type: bar\n" +
			"  required: bar\n" +
			"  properties: bar\n" +
			"prompt: A worker that fact-checks short claims and returns a verdict with citations.\n" +
			"summary: Audits copy against the Acme brand voice guide.\n" +
			"title: Brand voice analyzer\n" +
			"visibility: public\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workers", "create",
		)
	})
}

func TestWorkersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers", "retrieve",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
			"--stream", "true",
		)
	})
}

func TestWorkersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers", "update",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
			"--instructions", "You are a brand voice analyzer. Read the input text and report whether it matches Acme's playful, plain-spoken house style. Quote specific phrases.",
			"--output-schema", "{type: bar, required: bar, properties: bar}",
			"--summary", "Audits copy against the Acme brand voice guide.",
			"--title", "Claim verdict v2",
			"--visibility", "private",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: >-\n" +
			"  You are a brand voice analyzer. Read the input text and report whether it\n" +
			"  matches Acme's playful, plain-spoken house style. Quote specific phrases.\n" +
			"outputSchema:\n" +
			"  type: bar\n" +
			"  required: bar\n" +
			"  properties: bar\n" +
			"summary: Audits copy against the Acme brand voice guide.\n" +
			"title: Claim verdict v2\n" +
			"visibility: private\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workers", "update",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		)
	})
}

func TestWorkersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers", "delete",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		)
	})
}

func TestWorkersRetrieveEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers", "retrieve-email",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		)
	})
}
