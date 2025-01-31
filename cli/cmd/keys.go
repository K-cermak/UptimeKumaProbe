package cmd

import (
	"UptimeKumaProbe/db"
	"UptimeKumaProbe/helpers"
	"fmt"
)


func ViewAllKeys() {
	if !db.DatabaseExist() {
		helpers.PrintError(true, "Database does not exist, run <kprobe db init> first")
	}

	fmt.Println("\033[1m*probe_name\033[0m")
	fmt.Println(" -> " + db.GetValue("probe_name"))
	fmt.Println("    \033[3mProbe name is used to identify the probe in the API requests.\033[0m")

	fmt.Println("\033[1mdb_version\033[0m")
	fmt.Println(" -> " + db.GetValue("db_version"))
	fmt.Println("    \033[3mDatabase version.\033[0m")

	fmt.Println("\033[1mdb_init_time\033[0m")
	fmt.Println(" -> " + db.GetValue("db_init_time"))
	fmt.Println("    \033[3mDatabase initialization time.\033[0m")

	fmt.Println("\033[1mconfig_set\033[0m")
	fmt.Println(" -> " + db.GetValue("config_set"))
	fmt.Println("    \033[3mConfig set status, if set to true, you already set the config file.\033[0m")

	//delete_after
	fmt.Println("\033[1m*delete_after\033[0m")
	fmt.Println(" -> " + db.GetValue("delete_after"))
	fmt.Println("    \033[3mDelete after is used to set the time to delete the scan history data after. Updating this value will not affect the existing data.\033[0m")

	fmt.Println("\033[1m*api_port\033[0m")
	fmt.Println(" -> " + db.GetValue("api_port"))
	fmt.Println("    \033[3mAPI port is used to set the port for the API server.\033[0m")

	fmt.Println("\033[1m*editor_endpoint\033[0m")
	fmt.Println(" -> " + db.GetValue("editor_endpoint"))
	fmt.Println("    \033[3mEditor endpoint is used to enable/disable the /editor endpoint.\033[0m")

	fmt.Println("\n(values with * can be changed using <kprobe keys set <key> <value>> command)")


}