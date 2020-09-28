package internals

import (
	"fmt"
)

func WhereQueryIntUpdate(str, key string, value int64, condition string) string {

	if str != "" && condition != "" {
		str = fmt.Sprintf("%s %s ", str, condition)
	}

	str = fmt.Sprintf("%s %s=%d", str, key, value)

	return str
}

func WhereQueryStrUpdate(str, key, value, condition string) string {

	if str != "" && condition != "" {
		str = fmt.Sprintf("%s %s ", str, condition)
	}

	if value == "" {
		return str
	}

	str = fmt.Sprintf("%s %s='%s'", str, key, value)

	return str
}
