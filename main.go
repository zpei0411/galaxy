package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"galaxy/parse_roman"
)

func main() {
	wordMap := make(map[string]string)
	corrency := make(map[string]int)
	filename := "galaxy_roman.txt"
	strs, err := ReadLine(filename, wordMap, corrency)
	if err != nil {
		fmt.Printf("some errors occured:%s", err)
	}

	fmt.Println("------OUTPUT-----")
	for _, s := range strs {
		fmt.Println(s)
	}
}

func ReadLine(filename string, wordMap map[string]string, corrency map[string]int) ([]string, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	resultStr := []string{}
	buf := bufio.NewReader(f)
	for {
		// line, err := buf.ReadString('\n')
		// if err != nil || io.EOF == err {
		// 	break
		// }
		line, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		str, err := HandleLine(string(line), wordMap, corrency)
		if err != nil {
			return nil, err
		}
		if str != "" {
			resultStr = append(resultStr, str)
		}
	}
	return resultStr, nil
}

func HandleLine(line string, wordMap map[string]string, corrency map[string]int) (string, error) {
	str, err := parse_roman.HandleLine(line, wordMap, corrency)
	if err != nil {
		return "", err
	}
	return str, nil
}
