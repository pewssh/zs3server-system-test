package cli_tests

import (
	"os"
	"testing"
	"time"
	test "zs3server/tests/internal/cli/tests"
	cli_utils "zs3server/tests/internal/cli/util"

	"github.com/stretchr/testify/assert"
)

func TestZs3Server(testSetup *testing.T) {
	t := test.NewSystemTest(testSetup)

	// check if mc command is available
	if _, err := os.Stat("mc"); os.IsNotExist(err) {
		t.Fatalf("mc is not installed")
	} else {
		t.Logf("mc is installed")
	}

	// listing the buckets in the command
	t.RunSequentially("Should list the buckets", func(t *test.SystemTest) {
		output, _ := cli_utils.RunCommand(t, "mc ls play", 1, time.Hour*2)
		// check if error exist in log
		assert.NotContains(t, output, "error")

	})

	t.RunSequentially("Test Bucket Creation", func(t *test.SystemTest) {
		output, _ := cli_utils.RunCommand(t, "mc mb zcn/custombucket", 1, time.Hour*2)
		assert.Contains(t, output, "Bucket created successfully `zcn/custombucket`.")
	})

	t.RunSequentially("Test Copying File Upload", func(t *test.SystemTest) {
		// create a file with content
		_, _ = cli_utils.RunCommand(t, "mc mb zcn/custombucket", 1, time.Hour*2)

		file, err := os.Create("a.txt")
		if err != nil {
			t.Fatalf("Error creating file: %v", err)
		}
		defer file.Close()

		_, err = file.WriteString("test")
		if err != nil {
			t.Fatalf("Error writing to file: %v", err)
		}

		output, _ := cli_utils.RunCommand(t, "mc cp a.txt zcn/custombucket", 1, time.Hour*2)

		assert.NotContains(t, output, "mc: <ERROR>")

		os.Remove("a.txt")
	})

	t.RunSequentially("Test for moving file", func(t *test.SystemTest) {
		_, _ = cli_utils.RunCommand(t, "mc mb zcn/custombucket", 1, time.Hour*2)

		file, err := os.Create("a.txt")
		if err != nil {
			t.Fatalf("Error creating file: %v", err)
		}
		defer file.Close()

		_, err = file.WriteString("test")
		if err != nil {
			t.Fatalf("Error writing to file: %v", err)
		}

		_, _ = cli_utils.RunCommand(t, "mc cp a.txt zcn/custombucket", 1, time.Hour*2)

		output, _ := cli_utils.RunCommand(t, "mc mv zcn/custombucket/a.txt zcn/custombucket/b", 1, time.Hour*2)
		assert.NotContains(t, output, "mc: <ERROR>")
	})

	t.RunSequentially("Test for copying file ", func(t *test.SystemTest) {
		// create a file with content
		output, _ := cli_utils.RunCommand(t, "mc cp a.txt zcn/custombucket", 1, time.Hour*2)

		assert.NotContains(t, output, "mc: <ERROR>")
	})

	t.RunSequentially("Test for removing file", func(t *test.SystemTest) {
		output, _ := cli_utils.RunCommand(t, "mc rm zcn/custombucket/b/a.txt", 1, time.Hour*2)
		assert.Contains(t, output, "Removed `zcn/custombucket/b/a.txt`.")
	})

	t.RunSequentially("Test for removing bucket", func(t *test.SystemTest) {
		output, _ := cli_utils.RunCommand(t, "mc rb zcn/custombucket --force", 1, time.Hour*2)
		assert.Contains(t, output, "Removed `zcn/custombucket` successfully.")
	})

	// clean up commands
	t.RunSequentially("Clean up", func(t *test.SystemTest) {
		_, _ = cli_utils.RunCommand(t, "rm -rf a.txt", 1, time.Hour*2)
		_, _ = cli_utils.RunCommand(t, "mc rb zcn/custombucket --force", 1, time.Hour*2)
	})

}
