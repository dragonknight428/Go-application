package recursionutil

import (
	"SampleGoLand/csd/recursion/recursionutil"
	"SampleGoLand/test/assert"
	"SampleGoLand/test/recursion/file"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// Author: Sinan Kesen
func TestReverseString(t *testing.T) {
	values := file.ReadTestFile("../testfiles/reverse_string_test_values")
	for _, val := range values {
		testVal := strings.Split(val, " ")
		str := testVal[0]
		expected := testVal[1]
		actual := recursionutil.ReverseString(str)
		message := fmt.Sprintf("ReverseStr(%s) = %s, want %s", str, actual, expected)
		assert.Equals(t, message, expected, actual)
	}
}

// Author: Sinan Kesen
func TestFactorial(t *testing.T) {
	values := file.ReadTestFile("../testfiles/factorial_test_values")
	for _, val := range values {
		testVal := strings.Split(val, " ")
		input, _ := strconv.Atoi(testVal[0])
		expected, _ := strconv.Atoi(testVal[1])
		actual := recursionutil.Factorial(input)
		message := fmt.Sprintf("Factorial(%d) = %d, want %d", input, actual, expected)
		assert.Equals(t, message, expected, actual)
	}
}

// Author: Sinan Kesen
func TestGcd(t *testing.T) {
	tests := []struct {
		a, b int
		gcd  int
	}{
		{12, 16, 4},
		{10, 20, 10},
		{6, 3, 3},
		{81, 6, 3},
		{8148, 3, 3},
		{49, 7, 7},
		{64, 48, 16},
	}

	for _, test := range tests {
		actual := recursionutil.Gcd(test.a, test.b)
		message := fmt.Sprintf("Gcd(%d, %d) = %d, want %d", test.a, test.b, actual, test.gcd)
		assert.Equals(t, message, test.gcd, actual)
	}
}
