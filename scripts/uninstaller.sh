#!/bin/bash
## Tip: If you cannot run this script, try running `chmod +x uninstaller.sh`

set -e

INSTALL_DIR="/opt/kprobe"
CLI_BINARY="$INSTALL_DIR/cli"
API_SERVER_BINARY="$INSTALL_DIR/api-server"
WEB_EDITOR_FILE="$INSTALL_DIR/editor.html"
SERVICE_FILE="/etc/systemd/system/kprobe.service"
SYMLINK="/usr/local/bin/kprobe"
CURRENT_USER=$(logname)


if [[ $EUID -ne 0 ]]; then
    echo "This script must be run as root" >&2
    exit 1
fi

echo "Are you sure you want to uninstall Uptime Kuma Probe Extension? [y/N]"
read -r CONFIRMATION
if [[ "$CONFIRMATION" != "y" ]]; then
    echo "Uninstallation cancelled"
    exit 1
fi

echo "Uninstalling Uptime Kuma Probe Extension"

systemctl stop kprobe.service
systemctl disable kprobe.service

rm -f "$SERVICE_FILE"

rm -f "$SYMLINK"

rm -rf "$INSTALL_DIR"

echo "[SUCCESS] Uptime Kuma Probe Extension has been uninstalled successfully"
