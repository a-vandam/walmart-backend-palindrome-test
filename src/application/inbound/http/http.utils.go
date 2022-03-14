package http

import (
	"fmt"
	"strconv"
)

type ProductIdCtxKey struct {
}

const DefaultErrorInt int = 0

func parseInterfToInt(element interface{}) int {
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

const ErrorResponseBody string = `{"error":"%v", "resources":{}}`

func wrapErrAsJson(err error) string {
	return fmt.Sprintf(ErrorResponseBody, err.Error())
}
