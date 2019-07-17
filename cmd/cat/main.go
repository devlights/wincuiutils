package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/devlights/wincuiutils/cat"
)

func main() {
	// 引数処理
	var (
		fpath     string
		buf       bytes.Buffer
		inStream  io.Reader
		outStream io.Writer = &buf
	)

	flag.Parse()
	fpath = flag.Arg(0)

	// 入力ストリームを決定
	inStream = os.Stdin
	if fpath != "" {
		_, err := os.Stat(fpath)
		if os.IsNotExist(err) {
			message := fmt.Sprintf("ファイルが見つかりません [%s]\n", fpath)
			log.Fatal(message)
		}

		f, err := os.Open(fpath)
		if err != nil {
			message := fmt.Sprintf("ファイルのオープンに失敗. [%s]\n", fpath)
			log.Fatal(message)
		}

		defer f.Close()

		inStream = f
	}

	// 呼び出し
	p := cat.Input{
		InStream:  inStream,
		OutStream: outStream,
	}

	_, err := cat.Cat(&p)
	if err != nil {
		message := fmt.Sprintf("エラー [%s]\n", err)
		log.Fatal(message)
	}

	// 結果出力
	fmt.Print(buf.String())
}
