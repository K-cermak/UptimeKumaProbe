package cmd

import (
	"UptimeKumaProbe/db"
	"UptimeKumaProbe/helpers"
	"bufio"
	"os"
	"strings"
)

const (
	INVALID_FORMAT       string = "Invalid config file format"
	DUPL_SCAN_NAME       string = "Duplicate scan name detected"
	INVALID_SCAN_NAME    string = "Invalid scan name"
	TIMEOUT_NOT_INT      string = "Timeout is not an integer"
	TIMEOUT_OUT_OF_RANGE string = "Timeout out of range"
	INVALID_CODE         string = "Invalid status code"
	CODE_TOO_LONG        string = "Status code too long"
	INVALID_KEYWORD      string = "Invalid keyword"
)

func VerifyConfig(path string) {
	helpers.PrintInfo("Verifying config file")

	file, err := os.Open(path)
	if err != nil {
		helpers.PrintError(true, "Failed to open file ("+err.Error()+")")
	}
	defer file.Close()

	scanNames := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if len(fields) < 4 {
			helpers.PrintError(true, INVALID_FORMAT)
		}

		scanName := fields[0]
		if err := validateScanName(scanName, scanNames); err != "" {
			helpers.PrintError(true, "Invalid scan name: "+err)
		}

		scanType := fields[1]
		if scanType != "http" && scanType != "ping" {
			helpers.PrintError(true, "Invalid scan type: "+scanType)
		}

		scanAddress := fields[2]
		if len(scanAddress) > 256 {
			helpers.PrintError(true, "Invalid scan address")
		}

		scanTimeout := fields[3]
		scanTimeout = strings.TrimPrefix(scanTimeout, "timeout=")
		if err := validateTimeout(scanTimeout); err != "" {
			helpers.PrintError(true, "Invalid scan interval: "+err)
		}

		for _, field := range fields[4:] {
			if strings.HasPrefix(field, "status_code=") {
				if scanType == "ping" {
					helpers.PrintError(true, "Status code is not supported for ping scans")
				}

				if err := validateStatusCode(field); err != "" {
					helpers.PrintError(true, "Invalid status code: "+err)
				}

			} else if strings.HasPrefix(field, "keyword=") {
				if scanType == "ping" {
					helpers.PrintError(true, "Keyword is not supported for ping scans")
				}

				if err := validateKeyword(field); err != "" {
					helpers.PrintError(true, "Invalid keyword: "+err)
				}
			}
		}

		scanNames[scanName] = true
	}

	if err := scanner.Err(); err != nil {
		helpers.PrintError(true, "Failed to read file ("+err.Error()+")")
	}

	helpers.PrintSuccess("Config file verified successfully")
}

func SetConfig(path string) {
	helpers.PrintInfo("Replacing config file")

	if !db.DatabaseExist() {
		helpers.PrintError(true, "Database does not exist, run <kprobe db init> first")
	}

	//CONTROL CODE
	//DELETE ALL SCANS
	//ADD NEW

	helpers.PrintSuccess("Config file replaced successfully")
}

func validateScanName(scanName string, scanNames map[string]bool) string {
	if _, exists := scanNames[scanName]; exists {
		return DUPL_SCAN_NAME
	}
	if scanName == "" || len(scanName) > 32 {
		return INVALID_SCAN_NAME
	}
	for _, c := range scanName {
		if !('a' <= c && c <= 'z') && !('0' <= c && c <= '9') && c != '_' {
			return INVALID_SCAN_NAME
		}
	}
	return ""
}

func validateTimeout(scanTimeout string) string {
	num, correct := helpers.StrToInt(scanTimeout)
	if !correct {
		return TIMEOUT_NOT_INT
	}

	if num < 0 || num > 30000 {
		return TIMEOUT_OUT_OF_RANGE
	}

	return ""
}

func validateStatusCode(statusCode string) string {
	if len(statusCode) > 256 {
		return CODE_TOO_LONG
	}

	statusCode = strings.TrimPrefix(statusCode, "status_code=\"")
	statusCode = statusCode[:len(statusCode)-1]

	codes := strings.Split(statusCode, ",")
	if len(codes) == 0 {
		return INVALID_CODE
	}

	for _, code := range codes {
		if num, ok := helpers.StrToInt(code); !ok || num < 100 || num > 599 {
			return INVALID_CODE
		}
	}

	return ""
}

func validateKeyword(keyword string) string {
	if !strings.HasPrefix(keyword, "keyword=") {
		return INVALID_KEYWORD
	}

	startIdx := strings.Index(keyword, "\"")
	if startIdx == -1 {
		return INVALID_KEYWORD
	}

	endIdx := strings.LastIndex(keyword, "\"")
	if endIdx == startIdx {
		return INVALID_KEYWORD
	}

	return ""
}
