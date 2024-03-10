package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Place your code here.

	updEnvVarFromEnv(env)

	//nolint:gosec
	command := exec.Command(cmd[0], cmd[1:]...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return command.ProcessState.ExitCode()
	}

	return command.ProcessState.ExitCode()
}

func updEnvVarFromEnv(env Environment) {
	for key, value := range env {
		if value.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				log.Println(err)
			}
		}
		err := os.Setenv(key, value.Value)
		if err != nil {
			log.Println(err)
		}
	}
}
