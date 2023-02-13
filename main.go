package main

import (
	"flag"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	file := flag.String("input", "", "detect file mime")
	flag.Parse()

	ctype := mime.TypeByExtension(filepath.Ext(*file))
	if ctype != "" {
		fmt.Println(ctype)
		return
	}
	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf [512]byte
	n, err := io.ReadFull(f, buf[:])
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		panic(err)
	}
	fmt.Println(http.DetectContentType(buf[:n]))
}
