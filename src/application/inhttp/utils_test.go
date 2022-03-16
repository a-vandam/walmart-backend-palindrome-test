package inhttp

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInterfToInt32(t *testing.T) {
	testCases := []struct {
		testName       string
		interfaceValue interface{}
		expected       int
	}{
		{
			testName:       "a value that can be parsed to uint",
			interfaceValue: "230",
			expected:       230,
		},
		{
			testName:       "a non parseable value",
			interfaceValue: "example",
			expected:       0,
		},
	}

	for i, testCase := range testCases {
		t.Logf("testing parsing in method ParseInterfaceToInt32: test nmber: %v - %v", i, testCase.testName)
		valueObtained := parseInterfaceToInt(testCase.interfaceValue)
		if !assert.Equal(t, testCase.expected, valueObtained, "diff between expected ( %v ) and obtained ( %v )", testCase.expected, valueObtained) {
			t.FailNow()
		}
		t.Logf("--- OK --- test number %v - testName: %v ---", i, testCase.testName)
	}
}

func TestWrapErrAsJson(t *testing.T) {
	testCases := []struct {
		testName string
		err      error
		expected string
	}{
		{
			testName: "a common error1 embedded as json",
			err:      errors.New("failed to process sth"),
			expected: `{"error":"failed to process sth", "resources":{}}`,
		},
	}

	for i, testCase := range testCases {
		t.Logf("testing parsing in method ParseInterfaceToInt32: test nmber: %v - %v", i, testCase.testName)
		valueObtained := wrapErrAsJson(testCase.err)
		if !assert.Equal(t, testCase.expected, valueObtained, "diff between expected ( %v ) and obtained ( %v )", testCase.expected, valueObtained) {
			t.FailNow()
		}
		t.Logf("--- OK --- test number %v - testName: %v ---", i, testCase.testName)
	}
}

func TestParseInterfaceIntArrayToInt(t *testing.T) {
	testCases := []struct {
		testName       string
		interfaceValue interface{}
		expected       int
	}{
		{
			testName:       "a value that can be parsed to int",
			interfaceValue: []int{230},
			expected:       230,
		},
		{
			testName:       "a non parseable value",
			interfaceValue: "[example]",
			expected:       0,
		},
	}

	for i, testCase := range testCases {
		t.Logf("testing parsing in method ParseInterfaceToInt32: test nmber: %v - %v", i, testCase.testName)
		valueObtained := parseIdPathParamToInt(testCase.interfaceValue)
		if !assert.Equal(t, testCase.expected, valueObtained, "diff between expected ( %v ) and obtained ( %v )", testCase.expected, valueObtained) {
			t.FailNow()
		}
		t.Logf("--- OK --- test number %v - testName: %v ---", i, testCase.testName)
	}
}
