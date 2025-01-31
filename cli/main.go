package main

import (
	"os"
	"UptimeKumaProbe/helpers"
	"UptimeKumaProbe/cmd"
)

func main() {
	args := os.Args

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

	//config verify <path>
	if (helpers.ArgsMatch(args, []string{"*", "config", "verify", "*"})) {
		cmd.VerifyConfig(args[3])
		return
	}

	//config replace <path>
	if (helpers.ArgsMatch(args, []string{"*", "config", "replace", "*"})) {
		cmd.VerifyConfig(args[3])
		cmd.SetConfig(args[3])
		return	
	}

	helpers.PrintError(true, "Invalid command, rerun with <kprobe help> for help")
}