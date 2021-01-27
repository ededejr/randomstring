package randomstring

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func isNumeric(str string) bool {
	_, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false
	}

	return true
}

func testNumeric(t *testing.T, out string) {
	if !isNumeric(out) {
		t.Fatalf("%s is not numeric", out)
	}
}

func testLengthIsTen(t *testing.T, out string) {
	if len(out) != 10 {
		t.Fatalf("%s is not numeric", out)
	}
}

func containsNumbersOrSpecialCharacters(str string) bool {
	for _, charVariable := range str {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}

func testUppercaseOnly(t *testing.T, out string) {
	upper := strings.ToUpper(out)
	if upper != out || !containsNumbersOrSpecialCharacters(out) {
		t.Fatalf("%s contains non-uppercase characters", out)
	}
}

func testAlphabetsOnly(t *testing.T, out string) {
	if !containsNumbersOrSpecialCharacters(out) {
		t.Fatalf("%s contains non-alphabetic characters", out)
	}
}

func testSpecialCharactersOnly(t *testing.T, out string) {
	if containsNumbersOrSpecialCharacters(out) {
		t.Fatalf("%s contains non-special characters", out)
	}

	if isNumeric(out) {
		t.Fatalf("%s contains numeric characters", out)
	}
}

type testCase struct {
	in          string
	length      int
	description string
	expected    func(*testing.T, string)
}

func TestRandomString(t *testing.T) {
	cases := []testCase{
		{"0", 5, "All numeric characters with a length of 5", testNumeric},
		{"0", 10, "All numeric with a length of 10", testLengthIsTen},
		{"A", 10, "All uppercase characters with a length of 10", testUppercaseOnly},
		{"Aa", 10, "All alphabetic characters with a length of 10", testAlphabetsOnly},
		{"!", 1, "All special characters with a length of 1", testSpecialCharactersOnly},
	}

	for _, test := range cases {
		fmt.Println(fmt.Sprintf("Case: %s", test.description))
		observed := RandomString(test.in, test.length)
		test.expected(t, observed)
	}
}
