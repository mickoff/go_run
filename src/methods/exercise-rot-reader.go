
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i, v := range b {
		switch {
		case 'a' <= v && v < 'n', 'A' <= v && v < 'N':
			b[i] += 13
		case 'n' <= v && v <= 'z', 'N' <= v && v <= 'Z':
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
