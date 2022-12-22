#!/bin/bash

# Cleanup old files
crontab -u robot -l | grep -v '@reboot /home/robot/startup.sh'  | crontab -u robot -
mkdir -p /home/robot/bak
mv /home/robot/get_host_name_by_mac.py /home/robot/bak/.
mv /home/robot/host.name               /home/robot/bak/.
mv /home/robot/rpyc_robots.py          /home/robot/bak/.
mv /home/robot/rpyc_start.sh           /home/robot/bak/.
mv /home/robot/startup.sh              /home/robot/bak/.

# Create directory for ev3api-server
mkdir -p /opt/ev3api-server/

# Download files
curl -o /opt/ev3api-server/startup.sh https://raw.githubusercontent.com/EV3-OpenAPI/EV3-API/master/scripts/startup.sh
chmod +x /opt/ev3api-server/startup.sh

download_url=$(curl https://api.github.com/repos/EV3-OpenAPI/EV3-API/releases/latest | jq -r '.assets[] | select(.name == "server") | .browser_download_url')
curl -o server $download_url
chmod +x server

chown -R robot:robot /opt/ev3api-server

# Add robot user to sudoers file
echo "robot ALL=NOPASSWD: /usr/bin/hostnamectl set-hostname*, /bin/chvt 2, /bin/chvt 5" > /etc/sudoers.d/ev3api-server

# Create systemd service
cat <<EOT >> /etc/systemd/system/ev3api-server.service
[Unit]
Description=EV3API server
After=network.target

[Service]
Type=simple
ExecStart=/opt/ev3api-server/startup.sh
WorkingDirectory=/opt/ev3api-server/
User=robot

[Install]
WantedBy=multi-user.target
EOT

systemctl daemon-reload

# Enable and start ev3api server service
systemctl enable --now ev3api-server