package data

import (
	"fmt"
	"strconv"
)

type RegistrationNumber int64

func (r RegistrationNumber) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
