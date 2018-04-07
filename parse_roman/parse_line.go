package parse_roman

import (
	"fmt"
	"galaxy/roman"
	"strconv"
	"strings"
)

func HandleLine(line string, wordMap map[string]string, corrency map[string]int) (string, error) {
	romanArray := strings.Split(line, " ")
	//1
	_, ok := roman.RomanNumberMap[line[len(line)-1]]
	if ok {
		wordMap[romanArray[0]] = romanArray[len(romanArray)-1]
		return "", nil
	}

	//2
	if romanArray[len(romanArray)-1] == "Credits" {

		tempStr := wordMap[romanArray[0]] + wordMap[romanArray[1]]
		tempNum, err := roman.RomanToNum(tempStr)
		if err != nil {
			return "", err
		}
		money, err := strconv.Atoi(romanArray[4])
		if err != nil {
			return "", err
		}
		corrency[romanArray[2]] = money / tempNum
		return "", nil
	}

	// 3

	if romanArray[len(romanArray)-1] == "?" {
		tempstr1 := ""
		tempstr2 := ""
		if romanArray[1] == "much" && romanArray[2] == "is" {
			for _, s := range romanArray[3 : len(romanArray)-1] {
				tempstr1 += s + " "
				tempstr2 += wordMap[s]
			}

			resultNum, err := roman.RomanToNum(tempstr2)
			if err != nil {
				return "", err
			}
			result := fmt.Sprintf("%sis %d", tempstr1, resultNum)
			return result, nil
		} else if romanArray[1] == "many" && romanArray[3] == "is" {
			tempstr1 := ""
			tempstr2 := ""
			for _, s := range romanArray[4 : len(romanArray)-2] {
				tempstr1 += s + " "
				tempstr2 += wordMap[s]
			}

			resultNum, err := roman.RomanToNum(tempstr2)
			if err != nil {
				return "", err
			}
			result := fmt.Sprintf("%s%s is %d Credits", tempstr1, romanArray[len(romanArray)-2], corrency[romanArray[len(romanArray)-2]]*resultNum)
			return result, nil
		} else {
			return fmt.Sprintf("I have no idea what you are talking about"), nil
		}
	}

	return fmt.Sprintf("I have no idea what you are talking about"), nil
}
