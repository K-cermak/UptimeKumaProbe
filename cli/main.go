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

	//config view
	if (helpers.ArgsMatch(args, []string{"*", "config", "view"})) {
		cmd.ViewConfig()
		return
	}

	//keys view all / keys view <key>
	if (helpers.ArgsMatch(args, []string{"*", "keys", "view", "*"})) {
		cmd.ViewKeys(args[3])
		return
	}

	//keys set <key> <value>
	if (helpers.ArgsMatch(args, []string{"*", "keys", "set", "*", "*"})) {
		cmd.SetKeys(args[3], args[4])
		return
	}

	//test ping <address> <timeout_ms>
	if (helpers.ArgsMatch(args, []string{"*", "test", "ping", "*", "*"})) {
		cmd.PingTest(args[3], args[4])
		return
	}

	helpers.PrintError(true, "Invalid command, rerun with <kprobe help> for help")
}