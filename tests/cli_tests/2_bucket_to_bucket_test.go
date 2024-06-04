package cli_tests

import (
	"os"
	"testing"
	"time"
	test "zs3server/tests/internal/cli/tests"
	cli_utils "zs3server/tests/internal/cli/util"

	"github.com/stretchr/testify/assert"
)

func TestZs3ServerBucket(testSetup *testing.T) {
	t := test.NewSystemTest(testSetup)
	// create two buckets for testing purpose

	// test for moving the file from testbucket to testbucket2
	t.RunSequentially("Test for moving file from testbucket to testbucket2", func(t *test.SystemTest) {
		cli_utils.RunCommand(t, "mc mb zcn/testbucket", 1, time.Hour*2)

		cli_utils.RunCommand(t, "mc mb zcn/testbucket2", 1, time.Hour*2)

		file, err := os.Create("a.txt")
		if err != nil {
			t.Fatalf("Error creating file: %v", err)
		}
		defer file.Close()

		_, err = file.WriteString("test")
		if err != nil {
			t.Fatalf("Error writing to file: %v", err)
		}

		output_ls, _ := cli_utils.RunCommand(t, "mc mv a.txt zcn/testbucket", 1, time.Hour*2)

		assert.NotContains(t, output_ls, "mc: <ERROR>")
		output, _ := cli_utils.RunCommand(t, "mc mv zcn/testbucket/a.txt  zcn/testbucket2 ", 1, time.Hour*2)

		cli_utils.RunCommand(t, "mc rb zcn/testbucket2 --force", 1, time.Hour*2)

		cli_utils.RunCommand(t, "mc rb zcn/testbucket --force", 1, time.Hour*2)

		assert.NotContains(t, output, "mc: <ERROR>")

	})

}
