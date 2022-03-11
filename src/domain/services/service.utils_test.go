package services

import "testing"

func TestIsPalindromeUint(t *testing.T) {

	tests := []struct {
		testName string
		id       uint
		want     bool
	}{
		{
			testName: "a palindrome 3 digits id",
			id:       181,
			want:     true,
		},
		{
			testName: "a non palindrome id",
			id:       1811,
			want:     false,
		},
		{
			testName: "a palindrome 2 digits id",
			id:       11,
			want:     true,
		},
		{
			testName: "a non palindrome 2 digits id",
			id:       13,
			want:     false,
		},
		{
			testName: "a palindrome 4 digits id",
			id:       1331,
			want:     true,
		}, {
			testName: "a non palindrome 4 digits id",
			id:       1334,
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			t.Logf("id under test: %v", tt.id)
			if got := isPalindromeUint(tt.id); got != tt.want {
				t.Errorf("palindrome check for uint test failed. got: %v, want: %v", got, tt.want)
				t.Fail()
			}
		})
	}
}

func TestIsPalindromeString(t *testing.T) {

	tests := []struct {
		testName      string
		stringToCheck string
		want          bool
	}{
		{
			testName:      "a palindrome 3 char string",
			stringToCheck: "ada",
			want:          true,
		},
		{
			testName:      "a non palindrome 3 char string",
			stringToCheck: "abc",
			want:          false,
		},
		{
			testName:      "a palindrome 3 char string with diff cases",
			stringToCheck: "Ada",
			want:          true,
		},
		{
			testName:      "a palindrome 4 char string",
			stringToCheck: "adda",
			want:          true,
		},
		{
			testName:      "a non palindrome 4 char string",
			stringToCheck: "adca",
			want:          false,
		},
		{
			testName:      "a palindrome 4 char string with diff cases",
			stringToCheck: "adDa",
			want:          true,
		},
		{
			testName:      "a palindrome 7  char string with spaces ",
			stringToCheck: "abc cba",
			want:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			t.Logf("string under test: %v", tt.stringToCheck)
			if got := isPalindromeString(tt.stringToCheck); got != tt.want {
				t.Errorf("palindrome check for a string type test failed. got: %v, want: %v", got, tt.want)
				t.Fail()
			}
		})
	}
}
