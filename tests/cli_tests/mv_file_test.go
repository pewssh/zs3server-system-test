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

	// check if mc command is aviailable
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

}
