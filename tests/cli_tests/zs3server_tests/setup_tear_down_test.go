// create a tear down for all tests
package zs3servertests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"
	"time"
	zlogger "zs3server/tests/internal/cli/util/logger"
)

var allocationId string

func TestMain(m *testing.M) {
	globalSetup()
	timeout := time.Duration(60 * time.Minute)
	os.Setenv("GO_TEST_TIMEOUT", timeout.String())
	code := m.Run()
	globalTearDown()
	os.Exit(code)
}

func hasParentDir(path string) bool {
	return path != filepath.Dir(path)
}

func globalSetup() {
	currentDir, err := os.Getwd()
	if err != nil {
		zlogger.Logger.Fatal("Failed to get current working directory:", err)
	}

	// Check if a parent directory exists
	if !hasParentDir(currentDir) {
		zlogger.Logger.Fatal("Script must be run from a directory with a parent")
	}

	requiredCommands := map[string]string{
		"mc":    "../mc",
		"zbox":  "../zbox",
		"minio": "../minio",
		"warp":  "../warp",
	}

	for cmd, path := range requiredCommands {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			zlogger.Logger.Error(cmd + " is not installed")
			os.Exit(1)
		} else {
			zlogger.Logger.Info(cmd + " is  installed")
		}

		if requiredCommands[cmd] == requiredCommands["warp"] {
			zlogger.Logger.Info("All required commands are installed")
		} else {
			zlogger.Logger.Info("Checking for next command")

		}
	}

	// create allocation from allocation.yaml file
	data, parity, lock := read_file_allocation()
	cmd := exec.Command("../zbox", "newallocation", "--lock", lock, "--data", data, "--parity", parity)

	// get the allocation id for created

	output, err := cmd.CombinedOutput()

	if err != nil {
		zlogger.Logger.Error("Error creating allocation: ", err)
		os.Exit(1)
	} else {
		zlogger.Logger.Info("Allocation created successfully")
	}

	// use regex
	re := regexp.MustCompile(`Allocation created:\s*([a-f0-9]+)`)

	match := re.FindStringSubmatch(string(output))
	if len(match) > 1 {
		allocationId = match[1]
	}

	os.Setenv("MINIO_ROOT_USER", "someminiouser")
	os.Setenv("MINIO_ROOT_PASSWORD", "someminiopassword")

	cmd = exec.Command("./minio", "gateway", "zcn", "--console-address", ":8000")

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	zlogger.Logger.Info("Minio server started")

	println("Global setup code executed")

}

func globalTearDown() {
	println("Global teardown code Executing .......")
	err := exec.Command("../zbox", "delete", "--allocation", allocationId, "--remotepath", "/").Run()
	if err != nil {
		zlogger.Logger.Error("Error deleting allocation: ", err)
		os.Exit(1)
	} else {
		zlogger.Logger.Info("Allocation deleted successfully")
	}

}
