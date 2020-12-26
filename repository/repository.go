package repository

import (
	"bytes"
	"compress/zlib"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func GetGitObjectFilePath(sha1 string) string {
	return filepath.Join(".git", "objects", sha1[0:2], sha1[2:])
}

func GetUncompressedContent(targetPath string) string {
	buf, err := ioutil.ReadFile(targetPath)
	if err != nil {
		log.Fatal(err)
	}

	r, err := zlib.NewReader(bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// https://stackoverflow.com/questions/9644139/from-io-reader-to-string-in-go
	b := new(strings.Builder)
	n, err := io.Copy(b, r)
	if err != nil {
		log.Fatal(err, n)
	}
	return b.String()
}
