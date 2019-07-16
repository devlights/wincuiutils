package cat

import (
	"bytes"
	"testing"
)

func TestCat(t *testing.T) {
	// Arrange
	var (
		str    = "hello"
		inBuf  bytes.Buffer
		outBuf bytes.Buffer
	)

	inBuf.Write([]byte(str))

	parameter := Input{
		InStream:  &inBuf,
		OutStream: &outBuf,
	}

	// Act
	err := Cat(&parameter)

	// Assert
	if err != nil {
		t.Errorf("エラーが発生 %v", err)
	}

	outStr := outBuf.String()
	if str != outStr {
		t.Errorf("正しく出力できていない in[%s] out[%s]", str, outStr)
	}
}
