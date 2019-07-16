package cat

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Input - Cat()の引数を表します
type Input struct {
	InStream  io.Reader
	OutStream io.Writer
}

type parameterError struct {
	Message string
}

func (p *parameterError) Error() string {
	return fmt.Sprintf("Parameter Error [%s]\n", p.Message)
}

// Cat - Catコマンドを模倣している関数
func Cat(parameter *Input) error {
	if parameter == nil {
		return &parameterError{Message: "Input"}
	}

	var out io.Writer = os.Stdout
	if parameter.OutStream != nil {
		out = parameter.OutStream
	}

	scanner := bufio.NewScanner(parameter.InStream)
	for scanner.Scan() {
		txt := scanner.Text()
		out.Write([]byte(txt))
	}

	return nil
}
