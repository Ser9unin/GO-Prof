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
	case offset > limit:
		return ErrOffsetExceedsFileSize
	case offset < 0:
		return ErrOffsetNegative
	}

	newFile, ErrCreateFile = os.Create(toPath)

	if err != nil {
		if os.IsNotExist(ErrCreateFile) {
			return ErrCreateFile
		}
	}

	defer newFile.Close()

	var buf []byte

	switch {
	case limit == 0 && offset == 0:
		bar := pb.Full.Start64(srcFileSize)
		// proxy reader
		barReader := bar.NewProxyReader(srcFile)

		_, err = io.Copy(newFile, barReader)
		if err != nil {
			return err
		}
		bar.Finish()
	default:
		var bufSize int64
		bar := pb.Full.Start64(limit)

		if offset+limit > srcFileSize {
			bufSize = srcFileSize - offset
		} else {
			bufSize = limit
		}

		buf = make([]byte, bufSize)

		srcFile.ReadAt(buf, offset)

		barWriter := bar.NewProxyWriter(newFile)
		_, err = barWriter.Write(buf)

		if err != nil {
			return err
		}
		bar.Finish()
	}

	return nil
}
