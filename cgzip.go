// Fast compression using C gzip source code.
package cgzip

// #cgo CFLAGS: -O3
// #cgo LDFLAGS: /usr/local/lib/libz.a
// #include "src/gzip.h"
// #include "src/gzip.c"
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Get a char pointer to the first byte of a slice
func charp(input *[]byte) *C.char {
	result_header := (*reflect.SliceHeader)(unsafe.Pointer(input))
	return (*C.char)(unsafe.Pointer(result_header.Data))
}

func Uncompress(input []byte) ([]byte, error) {
	//压缩率不会超过100%
	outputlen := len(input) * 100
	if outputlen > 65536 {
		outputlen = 65536
	}
	output := make([]byte, 0, outputlen)
	ip, op := charp(&input), charp(&output)

	resultlen := int(C.GzipUncompress(ip, C.uLong(len(input)), op))

	if resultlen <= 0 || resultlen > cap(output) {
		return nil, fmt.Errorf("cgzip Uncompress failed!resultlen=%d inputlen:%d", resultlen, len(input))
	}

	output = (output)[0:resultlen]
	if len(output) != resultlen {
		return nil, fmt.Errorf("Failed to resize destination buffer")
	}
	return output, nil
}

func Compress(input []byte) ([]byte, error) {
	output := make([]byte, 0, len(input)*2)
	ip, op := charp(&input), charp(&output)
	resultlen := int(C.GzipCompress(ip, C.uLong(len(input)), op))
	if resultlen <= 0 || resultlen > cap(output) {
		return nil, fmt.Errorf("cgzip overran compression buffer. This shouldn't happen. "+
			"Expected: %d, got %d", cap(output), resultlen)
	}
	output = output[0:resultlen]
	return output, nil
}
