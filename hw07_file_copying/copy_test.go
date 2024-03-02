package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	// Place your code here.
	var err error

	t.Run("copy complete file", func(t *testing.T) {
		err = Copy("testdata/input.txt", "testdata/output.txt", 0, 0)
		require.NoError(t, err)

		actual, err := os.Open("testdata/output.txt")
		require.NoError(t, err)

		expected, err := os.Open("testdata/input.txt")

		require.NoError(t, err)

		actualByte, _ := io.ReadAll(actual)
		expectedByte, _ := io.ReadAll(expected)

		require.Equal(t, expectedByte, actualByte)

		expected.Close()
		actual.Close()
		_ = os.Remove("testdata/output.txt")
	})

	t.Run("copy limit 100", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", 0, 100)
		require.NoError(t, err)

		actual, err := os.Open("testdata/output.txt")
		require.NoError(t, err)

		actualByte, _ := io.ReadAll(actual)

		require.Equal(t, 100, len(actualByte))

		actual.Close()
		_ = os.Remove("testdata/output.txt")
	})

	t.Run("copy negative offset", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", -10, 0)

		require.Error(t, err)
	})

	t.Run("copy limit bigger than file", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", 0, 9999999)
		require.NoError(t, err)

		srcFile, _ := os.Open("testdata/input.txt")
		srcFileStat, _ := srcFile.Stat()
		srcFileSize := srcFileStat.Size()

		actual, err := os.Open("testdata/output.txt")
		require.NoError(t, err)

		actualByte, _ := io.ReadAll(actual)

		require.Equal(t, int(srcFileSize), len(actualByte))

		srcFile.Close()
		actual.Close()
		_ = os.Remove("testdata/output.txt")
	})

	t.Run("copy offset 100 limit 100", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", 100, 100)
		require.NoError(t, err)

		actual, err := os.Open("testdata/output.txt")
		require.NoError(t, err)

		actualByte, _ := io.ReadAll(actual)

		require.Equal(t, 100, len(actualByte))

		actual.Close()
		_ = os.Remove("testdata/output.txt")
	})
}
