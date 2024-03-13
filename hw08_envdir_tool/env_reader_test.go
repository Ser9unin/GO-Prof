package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	// Place your code here
	t.Run("check testdata dir", func(t *testing.T) {
		expectEnv := Environment{
			"BAR":   EnvValue{Value: "bar"},
			"EMPTY": EnvValue{Value: ""},
			"FOO":   EnvValue{Value: "   foo\nwith new line"},
			"HELLO": EnvValue{Value: `"hello"`},
			"UNSET": EnvValue{NeedRemove: true},
		}
		env, err := ReadDir("testdata/env/")

		require.NoError(t, err)
		require.Equal(t, env, expectEnv)
	})

	t.Run("Check filename", func(t *testing.T) {
		err := os.Mkdir("temp", 0o777)
		require.NoError(t, err)
		defer os.RemoveAll("temp/")

		// File with "=" must be skipped
		err = os.WriteFile(filepath.Join("temp/", "t=t"), []byte("bar"), 0o777)
		require.NoError(t, err)
		// File name in lower case
		err = os.WriteFile(filepath.Join("temp/", "test"), []byte("test"), 0o777)
		require.NoError(t, err)
		expectEnv := Environment{
			"test": EnvValue{Value: "test", NeedRemove: false},
		}

		env, err := ReadDir("temp/")
		require.NoError(t, err)
		require.Equal(t, env, expectEnv)
	})

	t.Run("Check empty dir", func(t *testing.T) {
		err := os.Mkdir("temp", 0o777)
		require.NoError(t, err)
		defer os.RemoveAll("temp/")

		env, err := ReadDir("temp/")
		require.NoError(t, err)
		require.Len(t, env, 0)
	})

	t.Run("Check dir not exists", func(t *testing.T) {
		_, err := ReadDir("some name")
		require.Error(t, err)
	})
}
