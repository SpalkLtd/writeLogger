package writeLogger

import (
	"bytes"
	"io"
	"strings"
)

const defaultSize uint = 10240

var newSize uint = defaultSize

type WriteLogger struct {
	buffer []byte
	out    io.Writer
}

func NewWriter(target io.Writer) WriteLogger {
	wl := WriteLogger{
		out:    target,
		buffer: make([]byte, newSize, newSize),
	}
	return wl
}

func NewWriterWithSize(target io.Writer, size uint) WriteLogger {
	wl := WriteLogger{
		out:    target,
		buffer: make([]byte, size, size),
	}
	return wl
}

func SetBufferSize(size uint) {
	newSize = size
}

func (wl *WriteLogger) Write(p []byte) (n int, err error) {
	outn, outerr := wl.out.Write(p)
	bsize := len(wl.buffer)
	isize := len(p)

	// fmt.Printf("buffer length and cap: %v %v\n", len(wl.buffer), cap(wl.buffer))

	// fmt.Println(bsize)
	// fmt.Println(isize)
	// fmt.Println(bsize - isize)

	if isize > bsize {
		wl.buffer = p[isize-bsize:]
		// fmt.Printf("%v\n", string(wl.buffer))

	} else {
		//bsize >= isize
		wl.buffer = append(wl.buffer[isize:], p...)

	}
	return outn, outerr
}

func (wl WriteLogger) Read(n int) []byte {
	if n > len(wl.buffer) {
		return wl.buffer
	}
	return wl.buffer[len(wl.buffer)-n:]
}

func (wl WriteLogger) ReadBuffer() *bytes.Buffer {
	return bytes.NewBuffer(wl.buffer)
}

func (wl WriteLogger) ReadString() string {
	return strings.Trim(string(wl.buffer), string(rune(0)))
}
