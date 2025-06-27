package repos

import (
	"io"
	"log"
	"os/exec"
	"regexp"

	"jerobas.com/yepee/types"
)

// var (
// 	mu sync.RWMutex
// )

func GetTunnels() []types.ParsedPSResult {
	cmd := exec.Command("bash", "-c", "ps -eo pid,args")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	result, err := io.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	// err = os.WriteFile("result.txt", result, 0064)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	test := regexp.MustCompile(`(?m)^\s+(\S+)\sssh -f -N -R (\d+)\:[a-z]+:(\d+)`)

	matched := test.FindAllStringSubmatch(string(result), -1)

	var matches []types.ParsedPSResult

	for _, match := range matched {
		matches = append(matches, types.ParsedPSResult{
			PID:     match[1],
			Command: match[2],
		})
	}

	return matches

	// mu.RLock()
	// mu.RUnlock()
}
