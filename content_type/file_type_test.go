package main_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/h2non/filetype"
)

const TestFileName = "./test.png"

func BenchmarkFiltypeMatch(b *testing.B) {
	f, err := os.Open(TestFileName)
	if err != nil {
		b.Fatal(err)
		return
	}
	defer closeFile(f)

	head := make([]byte, 261)
	if _, err := f.Read(head); err != nil {
		b.Fatal(err)
		return
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		kind, err := filetype.Match(head)
		if err != nil {
			b.Fatal(err)
			return
		}
		if kind == filetype.Unknown {
			b.Fatal("Unknown filetype")
			return
		}
		if kind.MIME.Value != "image/jpeg" {
			b.Fatal("Get the filetype failed")
			return
		}
	}
}

func BenchmarkHttpDetectContentType(b *testing.B) {
	f, err := os.Open(TestFileName)
	if err != nil {
		b.Fatal(err)
		return
	}
	defer closeFile(f)

	head := make([]byte, 512)
	if _, err := f.Read(head); err != nil {
		b.Fatal(err)
		return
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		contentType := http.DetectContentType(head)
		if contentType != "image/jpeg" {
			b.Fatal("Get the filetype failed")
			return
		}
	}
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println("close file filed")
	}
}
