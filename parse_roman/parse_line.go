package parse_roman

import (
	"strconv"
	"strings"
	"test/thoughtworks/galaxy/roman"
)

func HandleLine(line string, wordMap map[string]string, corrency map[string]int) error {
	lastIndex := len(line) - 1
	romanArray := strings.Split(line, "")
	_, ok := roman.RomanNumberMap[line[lastIndex]]
	if ok {
		wordMap[romanArray[0]] = romanArray[lastIndex]
		return nil
	}

	if romanArray[len(romanArray)-1] == "Credits" {
		tempStr := romanArray[0] + romanArray[1]
		tempNum, err := roman.RomanToNum(tempStr)
		if err != nil {
			return err
		}
		money, err := strconv.Atoi(romanArray[4])
		if err != nil {
			return err
		}
		corrency[romanArray[2]] = money / tempNum
		return nil
	}

}
