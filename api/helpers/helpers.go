package helpers

import (
	"log"
	"time"
)

type ScanRes struct {
	Generated string
	Passed    bool
}

func PrintInfo(info string) {
	log.Println("\033[0;34m[*]\033[0m " + info)
}

func PrintSuccess(success string) {
	log.Println("\033[0;32m[OK]\033[0m " + success)
}

func PrintQuestion(question string) {
	log.Println("\033[0;36m[?]\033[0m " + question)
}

func PrintWarning(warning string) {
	log.Println("\033[0;33m[WARN]\033[0m " + warning)
}

func PrintError(err string) {
	log.Println("\033[0;31m[ERROR]\033[0m " + err)
}

func GetCurrTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
