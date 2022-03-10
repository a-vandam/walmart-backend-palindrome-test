package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductsDatabaseConfigs(t *testing.T) {
	testCases := []getDBConfigsTestCase{{
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
	}, {
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
	},
	}

	t.Log("----- Testing db env vars retrieval------")

	for i, testCase := range testCases {
		t.Logf("test number: %v --- testName: %v", i, testCase)
		t.Logf("setting up test environment variables")
		for envKey, envValue := range testCase.presentVariables {
			t.Setenv(envKey, envValue)
		}
		t.Logf("testing function")
		envVarsObtained, errObtained := GetProductsDatabaseConfigs()
		assert.Equal(t, testCase.expected, envVarsObtained, "difference in value expected (%v) and obtained (%v)", testCase.expected, envVarsObtained)
		if !testCase.expectedErr && errObtained != nil {
			t.Fatalf("test failed as the function returned unexpected error: %v", errObtained)
		}
		if testCase.expectedErr {
			assert.Equal(t, testCase.expectedErrMsg, errObtained.Error(), "difference in expected  error (%v) and obtained (%v", testCase.expectedErr, errObtained)
		}
		t.Logf("test number %v - ( %v ) - OK!!!!", i, testCase.testName)

	}

}

type getDBConfigsTestCase struct {
	testName         string
	presentVariables map[string]string
	expected         ProductsDBConfigurations
	expectedErr      bool
	expectedErrMsg   string
}
