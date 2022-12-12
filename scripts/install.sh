#!/bin/bash

# Create directory for ev3api-server
mkdir -p /opt/ev3api-server/

wget https://raw.githubusercontent.com/PA-arslasel-machitic/EV3-API/master/scripts/startup.sh -o /opt/ev3api-server/startup.sh
wget https://raw.githubusercontent.com/PA-arslasel-machitic/EV3-API/master/scripts/get_hostname_by_mac.py -o /opt/ev3api-server/get_hostname_by_mac.py
chmod +x /opt/ev3api-server/startup.sh
chmod +x /opt/ev3api-server/get_hostname_by_mac.py

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