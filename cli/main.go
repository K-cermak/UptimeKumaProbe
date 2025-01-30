package main

import (
	"UptimeKumaProbe/cmd"
	"UptimeKumaProbe/helpers"
	"os"
)

func main() {
	args := os.Args

	//config verify <path>
	if (helpers.ArgsMatch(args, []string{"*", "config", "verify", "*"})) {
		helpers.PrintInfo("Verifying config file")
		cmd.VerifyConfig(args[3])
		helpers.PrintSuccess("Config file verified successfully")
		return
	}

	//config replace <path>
	if (helpers.ArgsMatch(args, []string{"*", "config", "replace", "*"})) {
		helpers.PrintInfo("Verifying config file")
		cmd.VerifyConfig(args[3])
		helpers.PrintSuccess("Config file verified successfully")

		//TODO
	
	}

	helpers.PrintError(true, "Invalid command, rerun with <kprobe help> for help")
}