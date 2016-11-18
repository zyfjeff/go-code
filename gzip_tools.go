package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	for _, file := range os.Args[1:] {
		compress(file)
	}
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return nil
	}

	defer in.Close()
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}

	defer out.Close()
	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
