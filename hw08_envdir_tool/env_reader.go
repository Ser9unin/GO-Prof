package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	// Place your code here
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	environment := make(Environment, len(files))

	for _, file := range files {
		// If filename contain "=" skip file
		if strings.Contains(file.Name(), "=") {
			continue
		}

		fileInfo, err := file.Info()
		if err != nil {
			continue
		}

		fileSize := fileInfo.Size()

		// if file is empty set Needremove = true
		if fileSize == 0 {
			environment[file.Name()] = EnvValue{NeedRemove: true}
			continue
		}

		// Get value from file contents
		value, err := getValue(dir, file.Name())
		if err != nil {
			return nil, err
		}

		environment[file.Name()] = EnvValue{Value: value}
	}

	return environment, nil
}

func getValue(dir, fileName string) (string, error) {
	f, err := os.Open(filepath.Join(dir, fileName))
	if err != nil {
		return "", err
	}

	defer f.Close()

	readFile := bufio.NewScanner(f)
	if !readFile.Scan() {
		return "", nil
	}

	readLine := readFile.Text()
	readLine = strings.ReplaceAll(readLine, "\x00", "\n")
	readLine = strings.TrimRightFunc(readLine, unicode.IsSpace)
	return readLine, nil
}
