package cat

import (
	"fmt"
	"io"
	"os"
)

// Input - Cat()の引数を表します
type Input struct {
	// 入力
	InStream io.Reader
	// 出力
	OutStream io.Writer
}

type parameterError struct {
	Message string
}

func (p *parameterError) Error() string {
	return fmt.Sprintf("Parameter Error [%s]\n", p.Message)
}

// Cat - Catコマンドを模倣している関数
func Cat(p *Input) (int64, error) {
	if p == nil {
		return 0, &parameterError{Message: "Input"}
	}

	var out io.Writer = os.Stdout
	if p.OutStream != nil {
		out = p.OutStream
	}

	written, err := io.Copy(out, p.InStream)
	if err != nil {
		return 0, err
	}

	return written, nil
}
