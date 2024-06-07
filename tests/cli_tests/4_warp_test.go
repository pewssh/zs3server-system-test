package cli_tests

import (
	"log"
	"os"
	"strings"
	"testing"
	"time"
	test "zs3server/tests/internal/cli/tests"
	cliutils "zs3server/tests/internal/cli/util"
)

func appendToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return err
	}
	return nil
}

func TestZs3serverWarpTests(testSetup *testing.T) {
	log.Println("Running Warp List Benchmark...")
	t := test.NewSystemTest(testSetup)

	output, err := cliutils.RunCommand(t, "./warp get --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword --objects 1", 1, time.Hour*2)
	if err != nil {
		testSetup.Fatalf("Error running warp list: %v\nOutput: %s", err, output)
	}
	log.Println("Warp List Output:\n", output)
	output_string := strings.Join(output, "\n")
	output_string = strings.Split(output_string, "----------------------------------------")[1]

	output_string = strings.Split(output_string, "warp: Starting cleanup")[0]

	output_string = "Condition 1: Get objects: 1 \n--------\n" + output_string
	err = appendToFile("warp-list_output.txt", output_string)

	if err != nil {
		testSetup.Fatalf("Error appending to file: %v\n", err)
	}
}
func TestZs3serverPutWarpTests(testSetup *testing.T) {
	log.Println("Running Warp Put Benchmark...")
	t := test.NewSystemTest(testSetup)

	output, err := cliutils.RunCommand(t, "./warp put --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword --objects 1", 1, time.Hour*2)
	if err != nil {
		testSetup.Fatalf("Error running warp put: %v\nOutput: %s", err, output)
	}
	log.Println("Warp Put Output:\n", output)
	output_string := strings.Join(output, "\n")
	output_string = strings.Split(output_string, "----------------------------------------")[1]

	output_string = strings.Split(output_string, "warp: Starting cleanup")[0]

	output_string = "Condition 1: Put  objects: 1 \n--------\n" + output_string
	err = appendToFile("warp-put_output.txt", output_string)

	if err != nil {
		testSetup.Fatalf("Error appending to file: %v\n", err)
	}
}

func TestZs3serverRetentionTests(testSetup *testing.T) {
	log.Println("Running Warp Retention Benchmark...")
	t := test.NewSystemTest(testSetup)

	output, err := cliutils.RunCommand(t, "./warp retention --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword --objects 1", 1, time.Hour*2)
	if err != nil {
		testSetup.Fatalf("Error running warp retention: %v\nOutput: %s", err, output)
	}
	log.Println("Warp Retention Output:\n", output)
	output_string := strings.Join(output, "\n")
	output_string = strings.Split(output_string, "----------------------------------------")[1]
	output_string = strings.Split(output_string, "warp: Starting cleanup")[0]

	output_string = "Condition 1: Retention : objects: 1 \n--------\n" + output_string
	err = appendToFile("warp-put_output.txt", output_string)

	if err != nil {
		testSetup.Fatalf("Error appending to file: %v\n", err)
	}
}

func TestZs3serverMultipartTests(testSetup *testing.T) {
	log.Println("Running Warp Multipart Benchmark...")
	t := test.NewSystemTest(testSetup)

	output, err := cliutils.RunCommand(t, "warp multipart --parts=500 --part.size=10MiB --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword", 1, time.Hour*2)

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

func TestMain(m *testing.M) {
	timeout := time.Duration(5 * time.Minute)
	os.Setenv("GO_TEST_TIMEOUT", timeout.String())

	os.Exit(m.Run())
}
