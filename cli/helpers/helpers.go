package helpers

import (
	"os"
	"strconv"
)

func PrintInfo(info string) {
	println("\033[0;34m[*]\033[0m " + info)
}

func PrintSuccess(success string) {
	println("\033[0;32m[OK]\033[0m " + success)
}

func PrintQuestion(question string) {
	println("\033[0;36m[?]\033[0m " + question)
}

func PrintWarning(warning string) {
	println("\033[0;33m[WARN]\033[0m " + warning)
}

func PrintError(fatal bool, err string) {
	println("\033[0;31m[ERROR]\033[0m " + err)
	if fatal {
		os.Exit(1)
	}
}

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