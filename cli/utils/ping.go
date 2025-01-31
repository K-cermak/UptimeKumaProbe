package utils

import (
	"fmt"
	"time"
	"UptimeKumaProbe/helpers"
	"UptimeKumaProbe/db"
	"github.com/prometheus-community/pro-bing"
)

func PingAddress(address string, timeout int, output bool) bool {
	pinger, err := probing.NewPinger(address)
	if err != nil {
		if output {
			helpers.PrintError(false, "Error creating ping command (" + err.Error() + ")")
		}

		return false
	}

	countStr := db.GetValue("ping_retries")
	count, correct := helpers.StrToInt(countStr)
	if !correct {
		if output {
			helpers.PrintError(false, "Error parsing ping retries value")
		}

		return false
	}

	pinger.SetPrivileged(true)
	pinger.Count = count
	pinger.Timeout = time.Duration(timeout) * time.Millisecond

	err = pinger.Run()
	if err != nil {
		if output {
			helpers.PrintError(false, "Error running ping command (" + err.Error() + ")")
		}
		return false
	}

	stats := pinger.Statistics()
	if output {
		fmt.Printf("PING Stats to %s", address)
		fmt.Printf("\n -> Sent: %d", stats.PacketsSent)
		fmt.Printf("\n -> Received: %d", stats.PacketsRecv)
		fmt.Printf("\n -> Lost: %f (%.2f%% loss)", stats.PacketLoss, stats.PacketLoss)
		fmt.Printf("\n -> Min RTT: %s", stats.MinRtt)
		fmt.Printf("\n -> Max RTT: %s", stats.MaxRtt)
		fmt.Printf("\n -> Avg RTT: %s\n", stats.AvgRtt)
	}

	return stats.PacketLoss < 20
}