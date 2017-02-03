package writeLogger

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

var toCopy string = "Hello World! The quick brown fox jumps over the lazy dog."
var tenBytes string = " lazy dog."

func TestCopyAndLog(t *testing.T) {
	outBuffer := bytes.NewBuffer(nil)
	writer := NewWriter(outBuffer)
	writer.Write([]byte(toCopy))
	outbytes, _ := ioutil.ReadAll(outBuffer)

	out := string(outbytes)
	if out != toCopy {
		t.Error(
			"For", toCopy,
			"got", out,
		)
	}

	end := string(writer.Read(10))
	if end != tenBytes {
		t.Error(
			"For", toCopy, 10,
			"expected", tenBytes,
			"got", end,
		)
	}

	if writer.ReadString() != toCopy {
		t.Error(
			"For", toCopy, "string",
			"expected", toCopy,
			"got", end,
		)
	}

	out = strings.Trim(string(writer.Read(999999999999999)), string(rune(0)))
	if out != toCopy {
		t.Error(
			"For", toCopy, "bigN",
			"got", out,
		)
	}
}

func TestRepeatedWrite(t *testing.T) {
	outBuffer := bytes.NewBuffer(nil)
	writer := NewWriter(outBuffer)

	for _, v := range toCopy {
		writer.Write([]byte{byte(v)})
	}

	outBuffer = writer.ReadBuffer()

	outbytes, _ := ioutil.ReadAll(outBuffer)
	out := strings.Trim(string(outbytes), string(rune(0)))
	if out != toCopy {
		t.Error(
			"For", toCopy,
			"got", out,
		)
	}
}

func TestWriteMoreThanBufferSize(t *testing.T) {
	outBuffer := bytes.NewBuffer(nil)
	writer := NewWriterWithSize(outBuffer, 10)

	writer.Write([]byte(toCopy))

	outBuffer = writer.ReadBuffer()

	outbytes, _ := ioutil.ReadAll(outBuffer)
	out := string(outbytes)
	if out != tenBytes {
		t.Error(
			"For", toCopy,
			"got", out,
		)
	}

	SetBufferSize(10)

	outBuffer = bytes.NewBuffer(nil)
	writer = NewWriter(outBuffer)

	writer.Write([]byte(toCopy))

	outBuffer = writer.ReadBuffer()

	outbytes, _ = ioutil.ReadAll(outBuffer)
	out = string(outbytes)
	if out != tenBytes {
		t.Error(
			"For", toCopy,
			"got", out,
		)
	}
}

func TestWriteMoreThanBufferSizeSlowly(t *testing.T) {
	outBuffer := bytes.NewBuffer(nil)
	writer := NewWriterWithSize(outBuffer, 10)

	for _, v := range toCopy {
		writer.Write([]byte{byte(v)})
	}

	outBuffer = writer.ReadBuffer()

	outbytes, _ := ioutil.ReadAll(outBuffer)
	out := string(outbytes)
	if out != tenBytes {
		t.Error(
			"For", toCopy,
			"got", out,
		)
	}
}
