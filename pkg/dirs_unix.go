// +build !windows

package cmd

import "os"

const xdgStateEnvName = "XDG_STATE_HOME"

func crevDir() string {
	xdgStateDir := os.Getenv(xdgStateEnvName)

	if xdgStateDir == "" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			panic(err)
		}
		xdgStateDir = homeDir + "/.local/state"
	}

	return xdgStateDir + "/crev"
}
