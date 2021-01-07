package array

import (
	"fmt"
	"strings"
)

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
