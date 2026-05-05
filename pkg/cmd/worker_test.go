// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Ramensoft/handinger-cli/internal/mocktest"
)

func TestWorkersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workers", "create",
			"--title", "Brand voice analyzer",
			"--instructions", "instructions",
			"--prompt", "prompt",
			"--visibility", "public",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"title: Brand voice analyzer\n" +
			"instructions: instructions\n" +
			"prompt: prompt\n" +
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
