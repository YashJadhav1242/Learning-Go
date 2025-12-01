package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sig, err := sha("sha1.go")
	if err != nil {
		fmt.Println("error:", err)
		return

	}
	fmt.Println("SHA1:", sig)

	fig, err := sha("http.log.gz")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("SHA1 (gz):", fig)

}
func sha(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(filename, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
