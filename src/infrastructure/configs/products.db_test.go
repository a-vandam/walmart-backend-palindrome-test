package configs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductDBAllConfigsPresent(t *testing.T) {

	testCase := getDBConfigsTestCase{
		testName: "test all vars present",
		presentVariables: map[string]string{
			mongoDatabaseUserKey:       "db user",
			mongoDatabaseHostKey:       "db host",
			mongoDatabaseNameKey:       "db name",
			mongoDatabasePasswordKey:   "db pwd",
			mongoDatabasePortKey:       "db port",
			mongoDatabaseAuthSourceKey: "authsrc",
		},
		expected: ProductsDBConfigurations{
			MongoDatabaseUsername: "db user",
			MongoDatabaseHost:     "db host",
			MongoDatabaseName:     "db name",
			MongoDatabasePassword: "db pwd",
			MongoDatabasePort:     "db port",
			MongoAuthSource:       "authsrc",
		},
		expectedErr: false,
	}
	testAndAssertDBConfigRetrieval(t, testCase)

}

func TestMissingUsername(t *testing.T) {

	testCase := getDBConfigsTestCase{

		testName: "missing username config",
		presentVariables: map[string]string{
			mongoDatabaseHostKey:       "db host",
			mongoDatabaseNameKey:       "db name",
			mongoDatabasePasswordKey:   "db pwd",
			mongoDatabasePortKey:       "db port",
			mongoDatabaseAuthSourceKey: "authsrc",
		},
		expected:       ProductsDBConfigurations{},
		expectedErr:    true,
		expectedErrMsg: fmt.Sprintf(MissingEnvVarErrorMsg, mongoDatabaseUserKey),
	}
	testAndAssertDBConfigRetrieval(t, testCase)

}

func testAndAssertDBConfigRetrieval(t *testing.T, testCase getDBConfigsTestCase) {
	t.Cleanup(func() {
		t.Logf("---- testName: %v ----", testCase.testName)
		t.Logf("setting up test environment variables")
		for envKey, envValue := range testCase.presentVariables {
			t.Setenv(envKey, envValue)
		}
		t.Logf("testing function")
		/**---------------------- FUNCTION UNDER TEST -----------------------**/
		envVarsObtained, errObtained := GetProductsDatabaseConfigs()
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

type getDBConfigsTestCase struct {
	testName         string
	presentVariables map[string]string
	expected         ProductsDBConfigurations
	expectedErr      bool
	expectedErrMsg   string
}
