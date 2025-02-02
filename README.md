# Uptime Kuma Probe Extension

### TODO Text


### Installation

```
sudo apt update
sudo apt install git

git clone https://github.com/K-cermak/UptimeKumaProbe

cd UptimeKumaProbe/scripts
sudo ./install.sh
```

- If this last command fails, try run:
```chmod +x install.sh```

- Tip: You can now remove the `UptimeKumaProbe` directory.
- App is installed in `/opt/kprobe`, API is installed as a service and will start on boot.


<br>

### Uninstallation

- If you have removed the `UptimeKumaProbe` directory (cloned Git repository), clone it again or download the [`uninstall.sh`](scripts/uninstall.sh) script.
```
git clone https://github.com/K-cermak/UptimeKumaProbe
```


- Run the uninstall script:

```
cd UptimeKumaProbe/scripts
sudo ./uninstall.sh
```

- If this last command fails, try run:
```chmod +x uninstall.sh```