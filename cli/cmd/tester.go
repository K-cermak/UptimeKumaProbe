package cmd

import (
	"UptimeKumaProbe/helpers"
	"UptimeKumaProbe/utils"
)

func PingTest(address string, timeout string) {
	count, correct := helpers.StrToInt(timeout)
	if !correct {
		helpers.PrintError(true, "Invalid timeout value")
	}

	helpers.PrintInfo("Pinging " + address + " with timeout " + timeout + " ms")
	ret := utils.PingAddress(address, count, true)
	if ret {
		helpers.PrintSuccess("Ping successful")
	} else {
		helpers.PrintWarning("Ping failed")
	}
}

func HttpTest(address string, timeout string) {
	count, correct := helpers.StrToInt(timeout)
	if !correct {
		helpers.PrintError(true, "Invalid timeout value")
	}

	helpers.PrintInfo("Performing HTTP request to " + address + " with timeout " + timeout + " ms")
	ret := utils.CheckHTTP(address, count, true, "", "")
	if ret {
		helpers.PrintSuccess("HTTP request successful")
	} else {
		helpers.PrintWarning("HTTP request failed")
	}
}