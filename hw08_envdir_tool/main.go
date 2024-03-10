package main

import (
	"log"
	"os"
)

func main() {
	// Place your code here.
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("not enough args")
	}
	env, err := ReadDir(args[0])
	if err != nil {
		log.Fatal("go-envdir: fatal: %w", err)
	}

	code := RunCmd(args[1:], env)

	os.Exit(code)
}
