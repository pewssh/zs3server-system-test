// create a tear down for all tests
package cli_tests

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
	zlogger "zs3server/tests/internal/cli/util/logger"
)

func TestMain(m *testing.M) {
	globalSetup()
	timeout := time.Duration(15 * time.Minute)
	os.Setenv("GO_TEST_TIMEOUT", timeout.String())
	code := m.Run()
	globalTearDown()
	os.Exit(code)
}

// func testforErrors(cmd *exec.Cmd) {
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	zlogger.Logger.Info(string(out))
// }

func globalSetup() {
	required_commands := []string{"mc", "zbox", "minio", "warp"}
	for i, j := range required_commands {
		if _, err := os.Stat(j); os.IsNotExist(err) {
			zlogger.Logger.Error(j + " is not installed")
			os.Exit(1)
		} else {
			zlogger.Logger.Info(j + " is  installed")
		}
		if i == len(required_commands)-1 {
			zlogger.Logger.Info("All required commands are installed")
		} else {
			zlogger.Logger.Info("Checking for next command")

		}
	}

	// _ = exec.Command("./zbox", "newallocation", "--lock", "10")
	os.Setenv("MINIO_ROOT_USER", "someminiouser")
	os.Setenv("MINIO_ROOT_PASSWORD", "someminiopassword")

	cmd := exec.Command("./minio", "gateway", "zcn", "--console-address", ":8000")

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	zlogger.Logger.Info("Minio server started")

	println("Global setup code executed")
}

func globalTearDown() {
	println("Global teardown code executed")
}
