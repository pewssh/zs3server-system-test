package cli_tests

import (
	"testing"
	"time"
	test "zs3server/tests/internal/cli/tests"
	cli_utils "zs3server/tests/internal/cli/util"

	"github.com/stretchr/testify/assert"
)

func TestZs3ServerBucket(testSetup *testing.T) {
	t := test.NewSystemTest(testSetup)
	// create two buckets for testing purpose

	// test for moving the file from custombucket1 to custombucket 2
	t.RunSequentially("Test for moving file from custombucket to testbucket", func(t *test.SystemTest) {
		cli_utils.RunCommand(t, "mc mb zcn/testbucket", 1, time.Hour*2)

		cli_utils.RunCommand(t, "mc mb zcn/testbucket2", 1, time.Hour*2)

		cli_utils.RunCommand(t, "echo 'test' > a.txt", 0, time.Hour*2)

		_, _ = cli_utils.RunCommand(t, "mc mv a.txt zcn/testbucket", 1, time.Hour*2)

		output, _ := cli_utils.RunCommand(t, "mc mv  zcn/custombucket  zcn/custombucket2 --recursive", 1, time.Hour*2)
		assert.NotContains(t, output, "mc: <ERROR>")
	})

}
