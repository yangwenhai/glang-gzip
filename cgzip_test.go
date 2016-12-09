package cgzip

import (
	"strings"
	"testing"
)

func TestCompression(t *testing.T) {
	input := []byte(strings.Repeat("Hello world, this is quite something", 10))
	output, err := Compress(input)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}
	if len(output) == 0 {
		t.Fatal("Output buffer is empty..")
	}
	t.Logf("Sizes: input=%d, output=%d, ratio=%.2f", len(input), len(output),
		float64(len(output))/float64(len(input)))
	decompressed, err := Uncompress(output)
	if err != nil {
		t.Fatalf("Decompression failed: %v", err)
	}
	if string(decompressed) != string(input) {
		t.Fatalf("Decompressed output != input: %q != %q", decompressed, input)
	}
}
