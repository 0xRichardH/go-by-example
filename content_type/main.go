package main

import (
	"fmt"
	"net/http"
	"os"
)

const SampleFileName = "./C0113.MP4"
const TheCountOfBytesToBeDected = 512

// To use http.DetectContentType to detect the content type of a file
func main() {
	f, err := os.Open(SampleFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Println("Failed to close fiale. %w", err)
		}
	}(f)

	buffer := make([]byte, TheCountOfBytesToBeDected)
	number, err := f.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	contentType := http.DetectContentType(buffer[:number])
	fmt.Println("ContentType: ", contentType)
}
