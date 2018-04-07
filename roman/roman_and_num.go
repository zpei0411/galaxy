package roman

import (
	"errors"
	"sort"
)

const (
	M  = 1000
	CM = 900
	D  = 500
	CD = 400
	C  = 100
	XC = 90
	L  = 50
	XL = 40
	X  = 10
	IX = 9
	V  = 5
	IV = 4
	I  = 1
)

var RomanNumberMap = map[byte]int{
	'I': I,
	'V': V,
	'X': X,
	'L': L,
	'C': C,
	'D': D,
	'M': M,
}

var NumberRomanMap = map[int]string{
	M:  "M",
	CM: "CM",
	D:  "D",
	CD: "CD",
	C:  "C",
	XC: "XC",
	L:  "L",
	XL: "XL",
	X:  "X",
	IX: "IX",
	V:  "V",
	IV: "IV",
	I:  "I",
}

func RomanToNum(roman string) (int, error) {
	err := verfy(roman)
	if err != nil {
		return 0, err
	}
	var sum int

	for i := 0; i < len(roman)-1; i++ {
		if RomanNumberMap[roman[i]] < RomanNumberMap[roman[i+1]] {
			sum -= RomanNumberMap[roman[i]]
		} else {
			sum += RomanNumberMap[roman[i]]
		}
	}
	return sum + RomanNumberMap[roman[len(roman)-1]], nil

}

func NumToRoman(num int) string {
	var numbers []int

	for k, _ := range NumberRomanMap {
		numbers = append(numbers, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	result := ""

	for _, i := range numbers {
		for num >= i {
			result += NumberRomanMap[i]
			num -= i
		}
	}
	return result

}

func verfy(roman string) error {
	for i := 0; i < len(roman); i++ {
		_, ok := RomanNumberMap[roman[i]]
		if !ok {
			return errors.New("The format of roman is wrong!")
		}
	}
	return nil
}
