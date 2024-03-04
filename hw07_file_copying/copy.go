package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile         = errors.New("unsupported file")
	ErrOffsetExceedsFileSize   = errors.New("offset exceeds file size")
	ErrOffsetNegative          = errors.New("offset has to be 0 or positive")
	ErrLimitNegative           = errors.New("limit has to be 0 or positive")
	ErrUnknownOriginalFileSize = errors.New("unknown original file size")
	ErrOpenFile                = errors.New("file not found in this destination")
	ErrCreateFile              = errors.New("can't create file")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.
	var (
		srcFile *os.File
		newFile *os.File
	)

	srcFile, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(ErrOpenFile) {
			return ErrOpenFile
		}
	}

	srcFileStat, err := srcFile.Stat()
	if err != nil {
		if os.IsNotExist(ErrOpenFile) {
			return ErrOpenFile
		}
	}

	srcFileSize := srcFileStat.Size()

	defer srcFile.Close()

	switch {
	case srcFileSize == 0:
		return ErrUnknownOriginalFileSize
	case offset > srcFileSize:
		return ErrOffsetExceedsFileSize
	case offset < 0:
		return ErrOffsetNegative
	case limit < 0:
		return ErrLimitNegative
	}

	newFile, ErrCreateFile = os.Create(toPath)

	if ErrCreateFile != nil {
		if os.IsNotExist(ErrCreateFile) {
			return ErrCreateFile
		}
	}

	defer newFile.Close()

	limit = realLimit(offset, limit, srcFileSize)

	bar := pb.Full.Start64(limit)

	if _, err = srcFile.Seek(offset, io.SeekStart); err != nil {
		return err
	}

	barReader := bar.NewProxyReader(srcFile)
	if _, err = io.CopyN(newFile, barReader, limit); err != nil {
		return err
	}

	bar.Finish()

	return nil
}

func realLimit(offset, limit, srcFileSize int64) int64 {
	if limit == 0 || offset+limit > srcFileSize {
		return srcFileSize - offset
	}
	return limit
}
