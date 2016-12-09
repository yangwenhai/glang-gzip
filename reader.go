package cgzip

import (
	"bytes"
	"io"
	"io/ioutil"
)

type Reader struct {
	bytes.Buffer
	underlying io.Reader
}

func (r *Reader) Read(data []byte) (int, error) {
	if len(data) < r.Buffer.Len() {
		input, err := ioutil.ReadAll(r.underlying)
		if err != nil {
			return 0, err
		}
		output, err := Uncompress(input)
		if err != nil {
			return 0, err
		}
		r.Buffer.Write(output)
	}
	return r.Buffer.Read(data)
}

func NewReader(r io.Reader) *Reader {
	return &Reader{underlying: r}
}

var _ io.Reader = NewReader(&Reader{})
