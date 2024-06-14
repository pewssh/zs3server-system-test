package cli_tests

import (
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	test "zs3server/tests/internal/cli/tests"
	cliutils "zs3server/tests/internal/cli/util"

	"gopkg.in/yaml.v3"
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

	file, err := os.Open("hosts.yaml")
	if err != nil {
		testSetup.Fatalf("Error opening hosts.yaml file: %v\n", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var hosts map[string]interface{}
	err = decoder.Decode(&hosts)
	if err != nil {
		testSetup.Fatalf("Error decoding hosts.yaml file: %v\n", err)
	}

	accessKey := hosts["access_key"].(string)
	secretKey := hosts["secret_key"].(string)
	port := hosts["port"].(int)
	server := hosts["server"].(string)
	host := strconv.FormatInt(int64(port), 10)

	commandGenerated := "./warp get --host=" + server + ":" + host + " --access-key=" + accessKey + " --secret-key=" + secretKey + "--duration 30s --obj.size 1KiB"
	log.Println("Command Generated: ", commandGenerated)
	output, err := cliutils.RunCommand(t, commandGenerated, 1, time.Hour*2)
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

// func TestZs3serverWarpConcurrentTests(testSetup *testing.T) {
// 	log.Println("Running Warp List Benchmark with concurrent ...")
// 	t := test.NewSystemTest(testSetup)
// 	// ./zbox updateallocation --allocation $allocationId --size 999999999999

// 	output, err := cliutils.RunCommand(t, "./warp get --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword --objects 50 --concurrent 50", 1, time.Hour*2)
// 	if err != nil {
// 		testSetup.Fatalf("Error running warp list: %v\nOutput: %s", err, output)
// 	}
// 	log.Println("Warp List Output:\n", output)
// 	output_string := strings.Join(output, "\n")
// 	output_string = strings.Split(output_string, "----------------------------------------")[1]

// 	// output_string = strings.Split(output_string, "warp: Starting cleanup")[0]

// 	output_string = "Condition 1: Get objects: 100 concurrent 50::  \n--------\n" + output_string
// 	log.Println("APending to file with this stat ", output_string)

// 	err = appendToFile("warp-list_output.txt", output_string)

// 	if err != nil {
// 		testSetup.Fatalf("Error appending to file: %v\n", err)
// 	}
// }

// func TestZs3serverPutWarpTests(testSetup *testing.T) {
// 	log.Println("Running Warp Put Benchmark...")
// 	t := test.NewSystemTest(testSetup)

// 	output, err := cliutils.RunCommand(t, "./warp put --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword", 1, time.Hour*2)
// 	if err != nil {
// 		testSetup.Fatalf("Error running warp put: %v\nOutput: %s", err, output)
// 	}
// 	log.Println("Warp Put Output:\n", output)
// 	output_string := strings.Join(output, "\n")
// 	output_string = strings.Split(output_string, "----------------------------------------")[1]

// 	output_string = "Condition 2 : Put  \n--------\n" + output_string
// 	err = appendToFile("warp-put_output.txt", output_string)

// 	if err != nil {
// 		testSetup.Fatalf("Error appending to file: %v\n", err)
// 	}
// }

// func TestZs3serverRetentionTests(testSetup *testing.T) {
// 	log.Println("Running Warp Retention Benchmark...")
// 	t := test.NewSystemTest(testSetup)

// 	output, err := cliutils.RunCommand(t, "./warp retention --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword --objects 1", 1, time.Hour*2)
// 	if err != nil {
// 		testSetup.Fatalf("Error running warp retention: %v\nOutput: %s", err, output)
// 	}
// 	log.Println("Warp Retention Output:\n", output)
// 	output_string := strings.Join(output, "\n")
// 	output_string = strings.Split(output_string, "----------------------------------------")[1]
// 	output_string = strings.Split(output_string, "warp: Starting cleanup")[0]

// 	output_string = "Condition 1: Retention : objects: 1 \n--------\n" + output_string
// 	err = appendToFile("warp-put_output.txt", output_string)

// 	if err != nil {
// 		testSetup.Fatalf("Error appending to file: %v\n", err)
// 	}
// }

// func TestZs3serverMultipartTests(testSetup *testing.T) {
// 	log.Println("Running Warp Multipart Benchmark...")
// 	t := test.NewSystemTest(testSetup)

// 	output, err := cliutils.RunCommand(t, "warp multipart --parts=500 --part.size=10MiB --host=localhost:9000 --access-key=someminiouser --secret-key=someminiopassword", 1, time.Hour*2)

// 	if err != nil {
// 		testSetup.Fatalf("Error running warp multipart: %v\nOutput: %s", err, output)
// 	}
// 	log.Println("Warp Multipart Output:\n", output)
// 	output_string := strings.Join(output, "\n")
// 	output_string = strings.Split(output_string, "----------------------------------------")[1]
// 	output_string = strings.Split(output_string, "warp: Starting cleanup")[0]

// 	output_string = "Condition 1: Retention : objects: 1 \n--------\n" + output_string
// 	err = appendToFile("warp-put_output.txt", output_string)

// 	if err != nil {
// 		testSetup.Fatalf("Error appending to file: %v\n", err)
// 	}
// }

func TestMain(m *testing.M) {
	timeout := time.Duration(15 * time.Minute)
	os.Setenv("GO_TEST_TIMEOUT", timeout.String())

	os.Exit(m.Run())
}
