package cli_tests

import (
	"log"
	"strings"
	"testing"
	"time"
	test "zs3server/tests/internal/cli/tests"
	cliutils "zs3server/tests/internal/cli/util"
)

func TestZs3serverFanoutTests(testSetup *testing.T) {
	log.Println("Running Warp Fanout Benchmark...")
	t := test.NewSystemTest(testSetup)

	output, err := cliutils.RunCommand(t, "./warp fanout --copies=50 --obj.size=512KiB --concurrent=8 --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword", 1, time.Hour*2)

	if err != nil {
		testSetup.Fatalf("Error running warp multipart: %v\nOutput: %s", err, output)
	}
	log.Println("Warp Multipart Output:\n", output)
	output_string := strings.Join(output, "\n")
	output_string = strings.Split(output_string, "----------------------------------------")[1]
	output_string = strings.Split(output_string, "warp: Starting cleanup")[0]

	output_string = "Condition 1: Retention : objects: 1 \n--------\n" + output_string
	err = appendToFile("warp-put_output.txt", output_string)

	if err != nil {
		testSetup.Fatalf("Error appending to file: %v\n", err)
	}
}
