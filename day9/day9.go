package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type File struct {
	id   int
	size int
}

func (fs *FS) New(p []byte) (n int, err error) {
	return 0, nil
}

const EMPTY = -1

type FS struct {
	s      []File // fileids
	lastId int
}

func (fs *FS) NewEmpty(len int) {
	fs.s = append(fs.s, File{id: EMPTY, size: len})
}

func (fs *FS) NewFile(len int) {
	fs.s = append(fs.s, File{id: fs.lastId, size: len})
	fs.lastId++
}

func (f File) IdToChar() string {
	if f.id == EMPTY {
		return "."
	}
	return string(byte(f.id + '0'))
}

func (f File) String() string {
	return strings.Repeat(f.IdToChar(), f.size)
}

func (fs FS) String() string {
	var buf strings.Builder
	for _, f := range fs.s {
		buf.WriteString(f.String())
	}
	return buf.String()
}

func CharToLen(c byte) int {
	return int(c - '0')
}

func main() {
	fs := FS{}
	m, _ := io.ReadAll(os.Stdin)
	for i, c := range m {
		if c == '\n' {
			break
		}
		if i%2 == 0 {
			fs.NewFile(CharToLen(c))
		} else {
			fs.NewEmpty(CharToLen(c))
		}
	}
	fmt.Println(fs)

}
