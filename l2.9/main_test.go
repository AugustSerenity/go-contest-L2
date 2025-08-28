package main

import (
	"testing"
)

func Test_Unpacing(t *testing.T) {
	expectedResult := "aaaabccddddde"
	var expectedError error = nil

	actualResult, actualError := Unpacking("a4bc2d5e")
	if actualError != expectedError {
		t.Errorf("Ожидали ошибку nil, но получили: %v", actualError)
	}
	if actualResult != expectedResult {
		t.Errorf("Ожидали результат %q, но получили %q", expectedResult, actualResult)
	}

	expectedResult2 := "abcd"
	var expectedError2 error = nil

	actualResult2, actualError2 := Unpacking("abcd")
	if actualError2 != expectedError2 {
		t.Errorf("Ожидали ошибку nil, но получили: %v", actualError2)
	}
	if actualResult2 != expectedResult2 {
		t.Errorf("Ожидали результат %q, но получили %q", expectedResult2, actualResult2)
	}

	expectedResult3 := ""
	var expectedError3 error = nil

	actualResult3, actualError3 := Unpacking("")
	if actualError3 != expectedError3 {
		t.Errorf("Ожидали ошибку nil, но получили: %v", actualError3)
	}
	if actualResult3 != expectedResult3 {
		t.Errorf("Ожидали результат %q, но получили %q", expectedResult3, actualResult3)
	}

	expectedResult4 := ""
	expectedError4 := ErrInvalidString

	actualResult4, actualError4 := Unpacking("45")

	if actualError4 != expectedError4 {
		t.Errorf("Ожидали ошибку %v, но получили %v", expectedError4, actualError4)
	}

	if actualResult4 != expectedResult4 {
		t.Errorf("Ожидали результат %q, но получили %q", expectedResult4, actualResult4)
	}
}
