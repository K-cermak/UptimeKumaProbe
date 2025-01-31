package main

import (
	"os"
	"UptimeKumaProbe/helpers"
	"UptimeKumaProbe/cmd"
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
		helpers.PrintInfo("Replacing config file")
		cmd.SetConfig(args[3])
		helpers.PrintSuccess("Config file replaced successfully")
		return	
	}

	//db init
	if (helpers.ArgsMatch(args, []string{"*", "db", "init"})) {
		cmd.InitDatabase()
		return
	}

	//db reset
	if (helpers.ArgsMatch(args, []string{"*", "db", "reset"})) {
		cmd.ResetDatabase()
		return
	}

	helpers.PrintError(true, "Invalid command, rerun with <kprobe help> for help")
}