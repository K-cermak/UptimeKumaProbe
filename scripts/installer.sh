#!/bin/bash
## Tip: If you cannot run this script, try running `chmod +x installer.sh`

set -e

INSTALL_DIR="/opt/kprobe"
CLI_BINARY="../bin/cli"
API_SERVER_BINARY="../bin/api-server"
WEB_EDITOR_FILE="../web-editor/editor.html"
SERVICE_FILE="/etc/systemd/system/kprobe.service"
CURRENT_USER=$(logname)


if [[ $EUID -ne 0 ]]; then
    echo "This script must be run as root" >&2
    exit 1
fi

echo "Are you sure you want to install Uptime Kuma Probe Extension? [y/N]"
read -r CONFIRMATION
if [[ "$CONFIRMATION" != "y" ]]; then
    echo "Installation cancelled"
    exit 1
fi


echo "Installing Uptime Kuma Probe Extension"
if [[ -d "$INSTALL_DIR" ]]; then
    echo "ERROR: Uptime Kuma Probe Extension is already installed. If you want to reinstall, please ruzn uninstaller.sh first"
    exit 1
fi


mkdir -p "$INSTALL_DIR"
chown "$CURRENT_USER":"$CURRENT_USER" "$INSTALL_DIR"
chmod 755 "$INSTALL_DIR"

cp "$CLI_BINARY" "$INSTALL_DIR/cli"
cp "$API_SERVER_BINARY" "$INSTALL_DIR/api-server"
cp "$WEB_EDITOR_FILE" "$INSTALL_DIR/editor.html"

chown "$CURRENT_USER":"$CURRENT_USER" "$INSTALL_DIR/cli"
chmod 755 "$INSTALL_DIR/cli"
chown root:root "$INSTALL_DIR/api-server"
chmod 755 "$INSTALL_DIR/api-server"
chown "$CURRENT_USER":"$CURRENT_USER" "$INSTALL_DIR/editor.html"
chmod 644 "$INSTALL_DIR/editor.html"


echo "Creating symlink for kprobe CLI"
ln -s "$INSTALL_DIR/cli" "/usr/local/bin/kprobe"


echo "Setting up API Server as a service"
cat <<EOF > "$SERVICE_FILE"
[Unit]
Description=KProbe API Server
After=network.target

[Service]
ExecStart=$INSTALL_DIR/api-server
Restart=always
User=root
WorkingDirectory=$INSTALL_DIR

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable kprobe.service
systemctl start kprobe.service

echo "[SUCCESS] Uptime Kuma Probe Extension has been installed successfully, you can now run 'kprobe' command to start the CLI"