package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path"

	"go.formulabun.club/metadatadb"
	"go.formulabun.club/storage"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s [file]\n", os.Args[0])
		return
	}

	pathToFile := os.Args[1]
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo := metadatadb.File{
		path.Base(pathToFile),
		string(getSum(file)[:]),
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	err = storage.Store(fileinfo, file)
	if err != nil {
		panic(err)
	}

}

func getSum(file io.Reader) string {
	hasher := md5.New()
	io.Copy(hasher, file)
	hash := hasher.Sum([]byte{})
	if len(hash) != md5.Size {
		panic("non expected hash result")
	}
	return fmt.Sprintf("%x", hash)
}
