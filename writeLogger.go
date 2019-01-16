package writeLogger

import (
	"bytes"
	"io"
	"strings"

	"github.com/armon/circbuf"
)

const defaultSize int = 10240

var newSize int = defaultSize

type WriteLogger struct {
	buffer *circbuf.Buffer
	out    io.Writer
	size   int
}

func NewWriter(target io.Writer) *WriteLogger {
	return NewWriterWithSize(target, newSize)
}

func NewWriterWithSize(target io.Writer, size int) *WriteLogger {
	if size < 0 {
		panic("Size must be >0")
	}
	b, _ := circbuf.NewBuffer(int64(size))
	wl := WriteLogger{
		out:    target,
		buffer: b,
		size:   int(size),
	}
	return &wl
}

func SetBufferSize(size int) {
	newSize = size
}

func (wl *WriteLogger) Write(p []byte) (n int, err error) {
	outn, outerr := wl.out.Write(p)
	n, err = wl.buffer.Write(p)
	if err != nil {
		return n, err
	}
	return outn, outerr
}

func (wl WriteLogger) Read(n int) []byte {
	b := wl.buffer.Bytes()
	//Copy slice so we don't modify circbuf internals
	buf := append(b[:0:0], b...)
	if n > len(buf) {
		return buf
	}
	return buf[len(buf)-n:]
}

func (wl WriteLogger) ReadBuffer() *bytes.Buffer {
	return bytes.NewBuffer(wl.Read(wl.size))
}

func (wl WriteLogger) ReadString() string {
	return strings.Trim(string(wl.Read(wl.size)), string(rune(0)))
}
