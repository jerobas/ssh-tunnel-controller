package repos

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"regexp"

	"github.com/jerobas/territo/config"
	"github.com/jerobas/territo/types"
	"github.com/jerobas/territo/utils"
)

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

	test := regexp.MustCompile(`(?m)^\s+(\d+)\sssh -f -N -R (\d+)\:[a-z]+:(\d+)`)

	matched := test.FindAllStringSubmatch(string(result), -1)

	var matches []types.ParsedPSResult

	for _, match := range matched {
		pid := utils.AtoiOrFatal(match[1])
		internalPort := utils.AtoiOrFatal(match[2])
		externalPort := utils.AtoiOrFatal(match[3])

		matches = append(matches, types.ParsedPSResult{
			PID:          pid,
			InternalPort: internalPort,
			ExternalPort: externalPort,
		})
	}

	return matches
}

func CreateTunnel(tunnel types.CreateTunnelDTO) bool {
	currentTunnels := GetTunnels()

	for _, currentTunnel := range currentTunnels {
		if currentTunnel.InternalPort == tunnel.InternalPort || currentTunnel.ExternalPort == tunnel.ExternalPort {
			return false
		}
	}

	command := fmt.Sprintf("ssh -f -N -R %d:localhost:%d -i %s %s", tunnel.InternalPort, tunnel.ExternalPort, config.GetConfig().KeyPath, config.GetConfig().Url)
	cmd := exec.Command("bash", "-c", command)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return true
}

func KillTunnel(tunnel types.KillTunnelDTO) bool {
	currentTunnels := GetTunnels()

	tunnelExist := false
	for _, currentTunnel := range currentTunnels {
		if currentTunnel.PID == tunnel.PID {
			tunnelExist = true
			break
		}
	}

	if !tunnelExist {
		return false
	}

	command := fmt.Sprintf("kill %d", tunnel.PID)
	cmd := exec.Command("bash", "-c", command)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return true
}
