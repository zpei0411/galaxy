package roman

import "testing"

func TestRomanAndNum(t *testing.T) {
	r := "MCMXLIV"
	num, err := RomanToNum(r)
	if err != nil {
		t.Fatalf("error:%s", err)
	}

	str := NumToRoman(num)
	if r != str {
		t.Fatalf("test wrong")
	}
}
