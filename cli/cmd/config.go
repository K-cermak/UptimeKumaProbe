package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"
	"UptimeKumaProbe/helpers"
)

type ErrorMessage string

const (
	INVALID_FORMAT 	 	 ErrorMessage = "Invalid config file format"
    DUPL_SCAN_NAME       ErrorMessage = "Duplicate scan name detected"
	INVALID_SCAN_NAME    ErrorMessage = "Invalid scan name"
	TIMEOUT_NOT_INT      ErrorMessage = "Timeout is not an integer"
	TIMEOUT_OUT_OF_RANGE ErrorMessage = "Timeout out of range"
	INVALID_CODE 	     ErrorMessage = "Invalid status code"
	INVALID_KEYWORD      ErrorMessage = "Invalid keyword"
)

func VerifyConfig(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	scanNames := make(map[string]bool)
	var namesList []string

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
			log.Fatal(INVALID_FORMAT)
		}

		scanName := fields[0]
		if err := validateScanName(scanName, scanNames); err != "" {
			log.Fatal("Invalid scan name:", err)
		}

		scanType := fields[1]
		if scanType != "http" && scanType != "ping" {
			log.Fatal("Invalid scan type:", scanType)
		}

		//scanAddress == fields[2]
		 
		scanTimeout := fields[3]
		scanTimeout = strings.TrimPrefix(scanTimeout, "timeout=")
		if err := validateScanTimeout(scanTimeout); err != "" {
			log.Fatal("Invalid scan interval:", err)
		}

		for _, field := range fields[4:] {
			if strings.HasPrefix(field, "status_code=") {
				if scanType == "ping" {
					log.Fatal("Status code is not supported for ping scans")
				}

				if err := validateStatusCode(field); err != "" {
					log.Fatal("Invalid status code:", err)
				}

			} else if strings.HasPrefix(field, "keyword=") {
				if scanType == "ping" {
					log.Fatal("Keyword is not supported for ping scans")
				}

				if err := validateKeyword(field); err != "" {
					log.Fatal("Invalid keyword:", err)
				}
			}
		}

		scanNames[scanName] = true
		namesList = append(namesList, scanName)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Failed to read file:", err)
	}

	return namesList
}

func validateScanName(scanName string, scanNames map[string]bool) ErrorMessage {
	if _, exists := scanNames[scanName]; exists {
		return DUPL_SCAN_NAME
	}
	if scanName == "" || len(scanName) > 32 {
		return INVALID_SCAN_NAME
	}
	for _, c := range scanName {
		if !('a' <= c && c <= 'z') && !('0' <= c && c <= '9') && c != '_' {
			return  INVALID_SCAN_NAME
		}
	}
	return ""
}

func validateScanTimeout(scanTimeout string) ErrorMessage {
	num, correct := helpers.StrToInt(scanTimeout)
	if !correct {
		return TIMEOUT_NOT_INT
	}
	
	if num < 0 || num > 30000 {
		return TIMEOUT_OUT_OF_RANGE
	}

	return ""
}

func validateStatusCode(statusCode string) ErrorMessage {
	statusCode = strings.TrimPrefix(statusCode, "status_code=")
	codes := strings.Split(statusCode, ",")
	
	if len(codes) == 0 {
		return INVALID_CODE
	}
	
	for _, code := range codes {
		code = strings.TrimSpace(code)
		if num, ok := helpers.StrToInt(code); !ok || num < 100 || num > 599 {
			return INVALID_CODE
		}
	}

	return ""
}

func validateKeyword(keyword string) ErrorMessage {
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
