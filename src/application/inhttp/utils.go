package inhttp

import (
	"fmt"
	"strconv"
)

type ProductIdPathParamCtxKey struct {
}

const DefaultErrorInt int = 0

func parseInterfaceToInt(element interface{}) int {

	valueAsInt, parseableToInt := element.(int)
	if parseableToInt {
		return valueAsInt
	}
	valueAsString, parseable := element.(string)

	if !parseable {
		return DefaultErrorInt
	}

	valueAsInt, err := strconv.Atoi(valueAsString)
	if err != nil {
		return 0
	}
	defer func() {
		parseConvErr := recover()
		if parseConvErr != nil {
			valueAsInt = DefaultErrorInt
		}
	}()
	return valueAsInt
}

func parseIdPathParamToInt(element interface{}) int {
	elementAsInts, parseableToArray := element.([]int)
	if !parseableToArray {
		return DefaultErrorInt
	}
	elementAsInt := elementAsInts[0]
	defer func() {
		parseConvErr := recover()
		if parseConvErr != nil {
			elementAsInt = DefaultErrorInt
		}
	}()
	return elementAsInt
}

const ErrorResponseBody string = `{"error":"%v", "resources":{}}`

func wrapErrAsJson(err error) string {
	return fmt.Sprintf(ErrorResponseBody, err.Error())
}
