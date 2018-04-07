package parse_roman

import "testing"

func TestHandleLine(t *testing.T) {
	wordMap := make(map[string]string)
	corrency := make(map[string]int)

	str1 := "glob is I"
	s, err := HandleLine(str1, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "" {
		t.Fatal("test wrong str1")
	}

	str2 := "prok is V"
	s, err = HandleLine(str2, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "" {
		t.Fatal("test wrong str2")
	}

	str3 := "pish is X"
	s, err = HandleLine(str3, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "" {
		t.Fatal("test wrong str3")
	}

	str4 := "tegj is L"
	s, err = HandleLine(str4, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "" {
		t.Fatal("test wrong str4")
	}

	str5 := "glob glob Silver is 34 Credits"
	s, err = HandleLine(str5, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "" {
		t.Fatal("test wrong str5")
	}

	str6 := "glob prok Gold is 57800 Credits"
	s, err = HandleLine(str6, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "" {
		t.Fatal("test wrong str6")
	}

	str7 := "pish pish Iron is 3910 Credits"
	s, err = HandleLine(str7, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "" {
		t.Fatal("test wrong str7")
	}

	str8 := "how much is pish tegj glob glob ?"
	s, err = HandleLine(str8, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "pish tegj glob glob is 42" {
		t.Fatal("test wrong str8")
	}

	str9 := "how many Credits is glob prok Silver ?"
	s, err = HandleLine(str9, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "glob prok Silver is 68 Credits" {
		t.Fatal("test wrong str9")
	}

	str10 := "how many Credits is glob prok Gold ?"
	s, err = HandleLine(str10, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "glob prok Gold is 57800 Credits" {
		t.Fatal("test wrong str10")
	}

	str11 := "how many Credits is glob prok Iron ?"
	s, err = HandleLine(str11, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "glob prok Iron is 780 Credits" {
		t.Fatal("test wrong str11")
	}

	str12 := "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?"
	s, err = HandleLine(str12, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "I have no idea what you are talking about" {
		t.Fatal("test wrong str12")
	}

	str13 := "how drew rtt"
	s, err = HandleLine(str13, wordMap, corrency)
	if err != nil {
		t.Fatalf("error:%s", err)
	}
	if s != "I have no idea what you are talking about" {
		t.Fatal("test wrong str13")
	}
}
