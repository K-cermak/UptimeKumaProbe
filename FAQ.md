# FaQ Guide
- [Installation](#installation)
- [Editor](#editor)
- [Creating a scan, setup cron](#creating-a-scan-setup-cron)
- [API, Reverse proxy](#api-reverse-proxy)
- [Connect to Uptime Kuma](#connect-to-uptime-kuma)
- [All commands](#all-commands)
- [Uninstallation](#uninstallation)

<br>

## Installation
- Ideally, create a virtual server or specify the physical hardware where the service will run. Recommended requirements are 1 GB RAM and 1 CPU core and 1 GB free space for the database (more if you have a large number of scans or want to keep records for a longer number of days).
- We recommend using Ubuntu Server in the latest LTS release, but it is not a problem to use other distributions - it is just that this project has not been tested on them. The distribution must necessarily support systemctl.

<br>

- First, install Git and clone the repository.

```
sudo apt update
sudo apt install git

git clone https://github.com/K-cermak/UptimeKumaProbe
```

- Open the `UptimeKumaProbe` directory and run the installation script. App will be installed in `/opt/kprobe` directory, API is installed as a service and will start on boot.

```
cd UptimeKumaProbe/scripts
sudo ./install.sh
```

- If the last command fails, try run: `chmod +x install.sh`.

> [!NOTE]
> You can now remove the `UptimeKumaProbe` directory.








## Editor

## Creating a scan, setup cron

## API, Reverse proxy

## Connect to Uptime Kuma

## All commands

## Uninstallation

- If you have removed the `UptimeKumaProbe` directory (cloned Git repository), clone it again or download the [`uninstall.sh`](scripts/uninstall.sh) script.
```
git clone https://github.com/K-cermak/UptimeKumaProbe
```


- Run the uninstall script:

```
cd UptimeKumaProbe/scripts
sudo ./uninstall.sh
```

- If the last command fails, try run:
```chmod +x uninstall.sh```
