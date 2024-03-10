package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Place your code here.
	// unset Environment variables and set new based on env containts
	updEnvVarFromEnv(env)

	command := exec.Command(cmd[0], cmd[1:]...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		log.Fatal(err)
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
