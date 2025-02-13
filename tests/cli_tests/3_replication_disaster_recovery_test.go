package cli_tests

import (
	"os"
	"strconv"
	"testing"
	"time"
	test "zs3server/tests/internal/cli/tests"
	cli_utils "zs3server/tests/internal/cli/util"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func read_file_mc(testSetup *testing.T) (string, string, string, string, string, string, string) {
	file, err := os.Open("mc_hosts.yaml")
	if err != nil {
		testSetup.Fatalf("Error opening hosts.yaml file: %v\n", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var hosts map[string]interface{}
	err = decoder.Decode(&hosts)
	if err != nil {
		testSetup.Fatalf("Error decoding mc_hosts.yaml file: %v\n", err)
	}

	accessKey := hosts["access_key"].(string)
	secretKey := hosts["secret_key"].(string)
	port := hosts["port"].(int)
	concurrent := hosts["concurrent"].(int)
	server := hosts["server"].(string)
	secondary_server := hosts["secondary_server"].(string)
	s_port := hosts["secondary_port"].(int)

	host := strconv.FormatInt(int64(port), 10)
	secondary_port := strconv.FormatInt(int64(s_port), 10)
	concurrent_no := strconv.FormatInt(int64(concurrent), 10)
	return server, host, accessKey, secretKey, concurrent_no, secondary_port, secondary_server

}

func TestZs3ServerReplication(testSetup *testing.T) {
	t := test.NewSystemTest(testSetup)

	server, port, accessKey, secretKey, _, s_port, s_server := read_file_mc(testSetup)

	t.RunSequentially("Test for replication", func(t *test.SystemTest) {
		// creating two server
		_, _ = cli_utils.RunCommand(t, "mc alias set primary "+server+":"+port+" "+accessKey+" "+secretKey+" --api S3v2", 1, time.Hour*2)
		_, _ = cli_utils.RunCommand(t, "mc alias set secondary"+s_server+":"+s_port+" "+accessKey+" "+secretKey+" --api S3v2", 1, time.Hour*2)

		// create bucket in primary
		_, _ = cli_utils.RunCommand(t, "mc mb primary/mybucket", 1, time.Hour*2)

		// enable mirror in primary
		output, _ := cli_utils.RunCommand(t, "mc mirror --watch --force primary/mybucket secondary/mybucket", 1, time.Hour*2)

		assert.NotContains(t, output, "error")
	})

	t.RunSequentially("Test for Disaster Recovery", func(t *test.SystemTest) {
		// creating two server
		_, _ = cli_utils.RunCommand(t, "mc alias set primary "+server+":"+port+" "+accessKey+" "+secretKey+" --api S3v2", 1, time.Hour*2)
		_, _ = cli_utils.RunCommand(t, "mc alias set secondary"+s_server+":"+s_port+" "+accessKey+" "+secretKey+" --api S3v2", 1, time.Hour*2)

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
