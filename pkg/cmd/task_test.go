// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/Ramensoft/handinger-cli/internal/mocktest"
)

func TestTasksCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tasks", "create",
			"--worker-id", "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
			"--instructions", "instructions",
			"--prompt", "prompt",
			"--title", "Brand voice analyzer",
			"--visibility", "public",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"workerId: t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM\n" +
			"instructions: instructions\n" +
			"prompt: prompt\n" +
			"title: Brand voice analyzer\n" +
			"visibility: public\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tasks", "create",
		)
	})
}

func TestTasksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tasks", "retrieve",
			"--task-id", "tsk_01HZY31W2SZJ8MJ2FQTR3M1K9D",
		)
	})
}
