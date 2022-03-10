package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfEmptyFail(t *testing.T) {
	testCases := []ifEmptyFailTestCase{{
		testName:      "test set and get env var",
		envKeyToTest:  "env_var_key",
		valueSet:      "databaseName",
		valueExpected: "databaseName",
		expectedErr:   false,
	}, {
		testName:      "env var not set",
		envKeyToTest:  "missing_env",
		valueSet:      "",
		valueExpected: "",
		expectedErr:   true,
		expectedMsg:   `missing: "missing_env" compulsory variable`,
	},
	}

	t.Log("----- Testing env vars retrieval------")
	for i, testCase := range testCases {
		t.Logf("test number: %v --- testName: %v", i, testCase)
		t.Logf("setting up test environment")
		t.Setenv(testCase.envKeyToTest, testCase.valueSet)
		t.Logf("testing function")
		envVarObtained, errObtained := getCompulsoryEnvVar(testCase.envKeyToTest)
		assert.Equal(t, testCase.valueExpected, envVarObtained, "difference in value expected (%v) and obtained (%v)", testCase.valueExpected, envVarObtained)
		if !testCase.expectedErr && errObtained != nil {
			t.Fatalf("test failed as the function returned unexpected error: %v", errObtained)
		}
		if testCase.expectedErr {
			assert.Equal(t, testCase.expectedMsg, errObtained.Error(), "difference in expected  error (%v) and obtained (%v", testCase.expectedErr, errObtained)
		}
		t.Logf("test number %v: %v - OK!!!!", i, testCase)
	}
}

type ifEmptyFailTestCase struct {
	testName      string
	envKeyToTest  string
	valueSet      string
	valueExpected string
	expectedErr   bool
	expectedMsg   string
}
