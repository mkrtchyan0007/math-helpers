package mathhelpers

import "testing"

func TestFactorial(t *testing.T) {
	if Factorial(5) != 120 {
		t.Error("Factorial failed for 5")
	}
	if Factorial(0) != 1 {
		t.Error("Factorial failed for 0")
	}
	if Factorial(-1) != -1 {
		t.Error("Factorial failed for -1")
	}
}

func TestIsPrime(t *testing.T) {
	if !IsPrime(7) {
		t.Error("IsPrime failed for 7")
	}
	if IsPrime(4) {
		t.Error("IsPrime failed for 4")
	}
}

func TestSum(t *testing.T) {
	if Sum([]int{1, 2, 3, 4}) != 10 {
		t.Error("Sum failed for [1,2,3,4]")
	}
}
