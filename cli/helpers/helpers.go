package helpers

import (
	"strconv"
)

func StrToInt(str string) (int, bool) {
	num, err := strconv.Atoi(str)
    if err != nil {
        return 0, false
    }

	return num, true
}

func ArgsMatch(args []string, expectedArgs []string) bool {
	if len(args) != len(expectedArgs) {
		return false
	}

	for i, arg := range args {
		if arg != expectedArgs[i] && expectedArgs[i] != "*" {
			return false
		}
	}

	return true
}