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
			"--instructions", "instructions",
			"--output-schema", "{foo: bar}",
			"--prompt", "prompt",
			"--summary", "summary",
			"--task-id", "tsk_2Z-YWz3hFq6VlW",
			"--title", "Brand voice analyzer",
			"--visibility", "public",
			"--worker-id", "wrk_vk81XUHKHG-qr4",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instructions: instructions\n" +
			"outputSchema:\n" +
			"  foo: bar\n" +
			"prompt: prompt\n" +
			"summary: summary\n" +
			"taskId: tsk_2Z-YWz3hFq6VlW\n" +
			"title: Brand voice analyzer\n" +
			"visibility: public\n" +
			"workerId: wrk_vk81XUHKHG-qr4\n")
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

func TestTasksDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tasks", "delete",
			"--task-id", "tsk_01HZY31W2SZJ8MJ2FQTR3M1K9D",
		)
	})
}
