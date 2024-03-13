package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zenizh/go-capturer"
)

func TestRunCmd(t *testing.T) {
	// Place your code here

	t.Run("Check create new folder with file for Env var", func(t *testing.T) {
		dir := "temp/"
		err := os.Mkdir(dir, 0o777)
		require.NoError(t, err)
		defer os.RemoveAll(dir)

		err = os.Mkdir(filepath.Join(dir, "var/"), 0o777)
		require.NoError(t, err)

		err = os.WriteFile(filepath.Join(dir, "var/", "BAR"), []byte("bar"), 0o777)
		require.NoError(t, err)

		err = os.WriteFile(filepath.Join(dir, "t.sh"), []byte("#!/usr/bin/env bash\necho $1\necho $BAR\n"), 0o777)
		require.NoError(t, err)

		env, err := ReadDir("temp/var/")

		require.NoError(t, err)

		var returnCode int

		result := capturer.CaptureStdout(func() {
			returnCode = RunCmd([]string{filepath.Join(dir, "t.sh"), "something"}, env)
		})
		require.Equal(t, 0, returnCode)
		require.Equal(t, "something\nbar\n", result)
	})
}
