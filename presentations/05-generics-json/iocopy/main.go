package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	src := strings.NewReader("Go пишет всё через io.Reader\n")
	io.Copy(os.Stdout, src) // без ручных буферов и циклов
	os.WriteFile("/tmp/go-course-demo.txt", []byte("привет"), 0o644)
	data, err := os.ReadFile("/tmp/go-course-demo.txt")
	fmt.Println(string(data), err)
}
