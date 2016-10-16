package main

import (
	"archive/tar"
	//	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//Create a buffer to write our archive to
	//buf := new(bytes.Buffer)
	f, _ := os.Create("test.tar")

	tw := tar.NewWriter(f)

	//Create a new tar archive
	var files = []struct {
		name, Body string
	}{
		{"movies_list.txt", "this archive contains some text files"},
		{"getmovie.go", "test file"},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}

	}

	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	f.Seek(0, os.SEEK_SET)
	//r := bytes.NewReader()
	tr := tar.NewReader(f)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}
}
