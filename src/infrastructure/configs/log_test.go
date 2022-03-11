package configs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogConfigs(t *testing.T) {

	testCase := getLogConfigTestCase{
		testName: "test log level present",
		presentVariables: map[string]string{
			logLevelKey: "inexistent log level",
		},
		expected: LogConfigs{
			LogLevel: "inexistent log level",
		},
		expectedErr: false,
	}
	testAndAssetLogConfigRetrieval(t, testCase)

}

func TestMissingLogLevel(t *testing.T) {

	testCase := getLogConfigTestCase{

		testName: "missing log level config",
		presentVariables: map[string]string{
			"notLogLevel": "not log level",
		},
		expected:       LogConfigs{},
		expectedErr:    true,
		expectedErrMsg: fmt.Sprintf(MissingEnvVarErrorMsg, logLevelKey),
	}
	testAndAssetLogConfigRetrieval(t, testCase)

}

func testAndAssetLogConfigRetrieval(t *testing.T, testCase getLogConfigTestCase) {
	t.Cleanup(func() {
		t.Logf("---- testName: %v ----", testCase.testName)
		t.Logf("setting up test environment variables")
		for envKey, envValue := range testCase.presentVariables {
			t.Setenv(envKey, envValue)
		}
		t.Logf("testing function")
		/**---------------------- FUNCTION UNDER TEST -----------------------**/
		envVarsObtained, errObtained := GetLogConfigs()
		/**---------------------- END FUNCTION UNDER TEST -----------------------**/

		if !assert.Equal(t, &testCase.expected, envVarsObtained, "difference in value expected (%v) and obtained (%v)", testCase.expected, envVarsObtained) {
			return
		}
		if !testCase.expectedErr && errObtained != nil {
			t.Fatalf("test failed as the function returned unexpected error: %v", errObtained)
			return
		}
		if testCase.expectedErr {
			if errObtained == nil {
				t.Fatalf("test failed as the function did not return error")
				return
			}
			if !assert.Equal(t, testCase.expectedErrMsg, errObtained.Error(), "difference in expected  error (%v) and obtained (%v", testCase.expectedErrMsg, errObtained) {
				return
			}

		}
		t.Logf("OK!!!! - test case:  %v  - OK!!!!", testCase.testName)
	})

}

type getLogConfigTestCase struct {
	testName         string
	presentVariables map[string]string
	expected         LogConfigs
	expectedErr      bool
	expectedErrMsg   string
}
