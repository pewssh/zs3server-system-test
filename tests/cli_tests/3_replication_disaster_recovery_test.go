package cli_tests

import (
	"testing"
	"time"
	test "zs3server/tests/internal/cli/tests"
	cli_utils "zs3server/tests/internal/cli/util"

	"github.com/stretchr/testify/assert"
)

func TestZs3ServerReplication(testSetup *testing.T) {
	t := test.NewSystemTest(testSetup)

	t.RunSequentially("Test for replication", func(t *test.SystemTest) {
		// creating two server
		_, _ = cli_utils.RunCommand(t, "mc alias set primary http://localhost:9000 someminiouser someminiopassword --api S3v2", 1, time.Hour*2)
		_, _ = cli_utils.RunCommand(t, "mc alias set secondary http://localhost:9001 someminiouser someminiopassword --api S3v2", 1, time.Hour*2)

		// create bucket in primary
		_, _ = cli_utils.RunCommand(t, "mc mb primary/mybucket", 1, time.Hour*2)

		// enable mirror in primary
		output, _ := cli_utils.RunCommand(t, "mc mirror --watch --force primary/mybucket secondary/mybucket", 1, time.Hour*2)

		assert.NotContains(t, output, "error")
	})

	t.RunSequentially("Test for Disaster Recovery", func(t *test.SystemTest) {
		// creating two server
		_, _ = cli_utils.RunCommand(t, "mc alias set primary http://localhost:9000 someminiouser someminiopassword --api S3v2", 1, time.Hour*2)
		_, _ = cli_utils.RunCommand(t, "mc alias set secondary http://localhost:9001 someminiouser someminiopassword --api S3v2", 1, time.Hour*2)

		// create bucket in primary
		_, _ = cli_utils.RunCommand(t, "mc mb primary/mybucket", 1, time.Hour*2)

		// enable mirror in primary
		_, _ = cli_utils.RunCommand(t, "mc mirror --watch --force primary/mybucket secondary/mybucket", 1, time.Hour*2)

		// lets remove bucket from primary server and recover from secondary bucket

		// remove bucket from primary
		_, _ = cli_utils.RunCommand(t, "mc rb primary/mybucket", 1, time.Hour*2)

		// mirro from secondary bucket to primary bucket
		output, _ := cli_utils.RunCommand(t, "mc mirror --watch --force secondary/mybucket primary/mybucket", 1, time.Hour*2)

		assert.NotContains(t, output, "error")
	})

}
