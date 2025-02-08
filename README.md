# Uptime Kuma Probe Extension

By Karel Cermak | [Karlosoft](https://karlosoft.com).

<img src="https://cdn.karlosoft.com/cdn-data/ks/img/kprobe/github.png" width="700" alt="Uptime Kuma Probe Extension">



## Uptime Kuma Probe Extension
- <b>Uptime Kuma is good at monitoring services and notifications, but it cannot connect through a VPN to the network</b> and find out the status of internal services like Wi-Fi Access Points, Cameras and various other internal things.
- <b>This Probe solves this</b> - it can simply <b>monitor devices inside the network and provide an API to the main Uptime Kuma instance</b>. It is very <b>simple to install</b>, and does not affect the main Kuma Uptime in any way.

<br>

## How it works?
- Create a Linux server (ideally Ubuntu Server) in your internal network. There you will install this extension using [FAQ.md guide](FAQ.md). Then you configure the scans.
- In Uptime Kuma, you then set up a <b>classic HTTP scan</b> - the Probe API will return a certain response if the service works and another if it doesn't. Uptime Kuma will now get up-to-date information about the <b>status of the service even if it does not have access to the network</b>, but only to the API server.

<br>


## What can it do?
- Measure <b>ICMP ping</b> reachability with a timeout limit.
- Make an <b>HTTP request</b> with a timeout limit and check a <b>certain status code</b> and the <b>keyword in the response</b>.
- Simple <b>editor for scan configuration</b> and simple <b>CLI interface</b>.
- <b>API server</b> for the main Uptime Kuma instance.

<br>

## Can I use it without Uptime Kuma?
- Yes, you can use it as a <b>simple monitoring tool</b>. You can use the API to get the status of the services.
- But keep in mind that this app itself can't send notifications, nor will it ever (there is no plan to develop in this direction).

<br>

## How to start?
- You can find the <b>installation guide<b> in the [FAQ.md](FAQ.md) file. You can also find the uninstallation guide there and other useful information.


<br>
<br>

---

#### This project is not affiliated with Uptime Kuma. This is just a simple probe extension for monitoring tool Uptime Kuma. Developed by Karel Cermak (info@karlosoft.com).