// +build windows

package cmd

import "os"

const xdgStateEnvName = "XDG_STATE_HOME"

func crevDir() string {
	xdgStateDir := os.Getenv(xdgStateEnvName)

	if xdgStateDir == "" {
		xdgStateDir, err := os.UserCacheDir()

		if err != nil {
			panic(err)
		}
	}

	return xdgStateDir + "/crev"
}
