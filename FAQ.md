# FaQ Guide
- [Installation](#installation)
- [Editor](#editor)
- [Creating a scan, setup cron](#creating-a-scan-setup-cron)
- [API, Reverse proxy](#api-reverse-proxy)
- [Connect to Uptime Kuma](#connect-to-uptime-kuma)
- [List of All commands](#list-of-all-commands)
- [Uninstallation](#uninstallation)

<br>

## Installation
- Ideally, create a virtual server or specify the physical hardware where the service will run. Recommended requirements are 1 GB RAM and 1 CPU core and 1 GB free space for the database (more if you have a large number of scans or want to keep records for a longer number of days).
- We recommend using Ubuntu Server in the latest LTS release, but it is not a problem to use other distributions - it is just that this project has not been tested on them. The distribution must necessarily support systemctl.

<br>

- <b>First, install Git and clone the repository.</b>

```
sudo apt update
sudo apt install git

git clone https://github.com/K-cermak/UptimeKumaProbe
```

- <b>Open the `UptimeKumaProbe` directory and run the installation script. App will be installed in `/opt/kprobe` directory, API is installed as a service and will start on boot.</b>

```
cd UptimeKumaProbe/scripts
sudo ./install.sh
```

- If the last command fails, try run: `chmod +x install.sh`.

> [!NOTE]
> You can now remove the `UptimeKumaProbe` directory.

- You shoudl now init the database and restart the API service.

```
kprobe db init
sudo kprobe api restart
```

<br>
<br>

## Editor
- You can now open the editor and create a scan. You can find the editor [here](https://github.com/K-cermak/UptimeKumaProbe/blob/main/web-editor/editor.html) (download it and open it in your browser).

> [!NOTE]
> You can also find the editor at `http://YOUR_SERVER_IP/editor` URL.

- The editor is very simple to use, all the information is there. To download the configuration file, click on the "Verify Values" button and then on "Download Config". To reopen the old configuration file, click on "Load config."

<img src="https://cdn.karlosoft.com/cdn-data/ks/img/kprobe/editor.png" width="700" alt="Uptime Kuma Probe Extension">



<br>
<br>

## Creating a scan, setup cron
- After exporting the configuration file, run only this command:
```
kprobe config replace <path_to_config_file>
```

- If you want to verify the configuration file, you can run:
```
kprobe config verify <path_to_config_file>
```

> [!NOTE]
> The configuration file is stored in the database and your file is not used by Probe itself. You can delete it after the configuration is loaded.


<br>

- You can run the scan manually by running:
```
kprobe cron <type>
```

As a type, use:
- `all` - to run all scans
- `all_except:<names>` - to run all scans except the ones specified
- `only:<names>` - to run only the scans specified

For CRON setup, open the CRON editor:
```
crontab -e
```

And add for example the following line:
```
*/5 * * * * /usr/bin/kprobe cron all
```

This will run all scans every 5 minutes.



<br>
<br>


## API, Reverse proxy
- The API is running on port 80 by default. You can access it at `http://YOUR_SERVER_IP/status/<scan_name>`.
- The response is in JSON format and looks like this:
```
{
    "probe_name":"Probe Name", // Name of the probe
    "time":"2025-01-01 01:23:59", // Current time in YYYY-MM-DD HH:MM:SS format
    "scan_name":"scan_name", // Name of the scan
    "check":"2025-01-01T00:20:02Z", // Last check time in ISO format
    "status":"true" // Status of the scan, true for OK, false for error
}
```

- Now you have to somehow set up access to this server from the internet (or from the Uptime Kuma server). 
- For example, this can be by using a reverse proxy or you can use the port forwarding. This depends on the configuration of your local network. 
- If you want to change the port of the API, you can do it:
```
kprobe keys set api_port <port>
sudo kprobe api restart
```

<br>
<br>

## Connect to Uptime Kuma
- Create a new scan (monitor) in Uptime Kuma with the following parameters:
    - <b>Type:</b> HTTP(s) - Keyword
    - <b>URL:</b> `http(s)://YOUR_SERVER_IP/status/<scan_name>`
    - <b>Keyword:</b> `"status":"true"`

- Now you can see the status of the scan in Uptime Kuma.

<br>
<br>

## List of All commands
- You can see the list of all commands also by running: `kprobe help`.

<br>

```
kprobe cron <type>
```
- Start the cron job with the specified type.
    - Use 'all' to start all cron jobs.
    - Use 'all_except:<names>' to start all cron jobs except the specified ones (seperate names with comma without space).
    - Use 'only:<names>' to start only the specified cron jobs (seperate names with comma without space).

<br>

```
kprobe state
```
- View the current state of the scans.

<br>
<br>

```
kprobe history <scan_name> <from> <to>
```
- View the history of the specified scan.
- For <from> and <to> use the format 'YYYY-MM-DD HH:MM:SS'.

<br>
<br>

```
kprobe db init
```
- Initialize the database.

```
kprobe db reset
```
- Reset the database, this will delete all the data!

<br>
<br>

```
kprobe config verify <path>
```
- Verify the configuration file at the specified path.

```
kprobe config replace <path>
```
- Replace the current configuration with the one at the specified path.
- File is copied to the database, so you can delete the original file afterwards.

```
kprobe config view
```
- View the current configuration.

<br>
<br>

```
kprobe keys view all
```
- View all the keys with their values in the database.
- Keys with the * prefix can be changed.

```
kprobe keys view <key>
```
- View the value of the specified key.
- If the key has the * prefix, it can be changed.

```
kprobe keys set <key> <value>
```
- Set the value of the specified key.

<br>
<br>

```
kprobe test ping <address> <timeout_ms>
```
- Test the ping to the specified address with the specified timeout.
- Timeout is in milliseconds.

```
kprobe test http <address> <timeout_ms>
```
- Test the http request to the specified address with the specified timeout.
- Timeout is in milliseconds.

<br>
<br>

```
kprobe api test [service|http]
```
- Test the api service or the http service.
- Use 'service' to test the api service via systemctl.
- Use 'http' to test the api service via http request.

```
kprobe api restart
```
- Restart the api service.
- This command requires sudo privileges.

<br>
<br>


```
kprobe help
```
- Print this help message.

<br>
<br>

## Uninstallation

- If you have removed the `UptimeKumaProbe` directory (cloned Git repository), clone it again or download the [`uninstall.sh`](scripts/uninstall.sh) script.
```
git clone https://github.com/K-cermak/UptimeKumaProbe
cd UptimeKumaProbe/scripts
```


- Run the uninstall script:

```
sudo ./uninstall.sh
```

- If the last command fails, try run: `chmod +x uninstall.sh`.