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
	size   int
	lastId int
	flat   []int
}

func (fs *FS) NewEmpty(len int) {
	fs.size += len
	fs.s = append(fs.s, File{id: EMPTY, size: len})
}

func (fs *FS) NewFile(len int) {
	fs.size += len
	fs.s = append(fs.s, File{id: fs.lastId, size: len})
	fs.lastId++
}

func (f File) IdToChar() string {
	if f.id == EMPTY {
		return "."
	}
	if f.id > 9 {
		return "*"
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

func (fs *FS) Flatten() {
	fs.flat = make([]int, fs.size)
	i := 0
	for _, f := range fs.s {
		for range f.size {
			fs.flat[i] = f.id
			i++
		}
	}
}

func (fs *FS) Defrag() {
	i := 0
	j := fs.size - 1
	for i <= j {
		if fs.flat[j] == EMPTY {
			j--
			continue
		}
		if fs.flat[i] != EMPTY {
			i++
			continue
		}
		fs.flat[i], fs.flat[j] = fs.flat[j], fs.flat[i]
		i++
		j--
	}
}

func (fs *FS) FlatString() string {
	var buf strings.Builder
	for _, id := range fs.flat {
		if id == EMPTY {
			buf.WriteByte('.')
		} else {
			if id > 9 {
				buf.WriteRune('*')
			} else {
				buf.WriteByte(byte(id + '0'))
			}
		}
	}
	return buf.String()
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
	fmt.Printf("FS is %d blocks (%s)\n", fs.size, fs)
	fs.Flatten()
	fmt.Printf("Flattened is %s (%v)\n", fs.FlatString(), fs.flat)
	fs.Defrag()
	fmt.Printf("Defragged is %s (%v)\n", fs.FlatString(), fs.flat)
	// fmt.Println("Part1:", fs.Checksum())
}
