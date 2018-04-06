package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

}

func ReadLine(filename string) error {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		fmt.Println(line)
		if err != nil || io.EOF == err {
			break
		}
		HandleLine(line)
	}
}

func HandleLine(line string) {

}
