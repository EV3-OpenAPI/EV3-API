#!/bin/bash
# https://www.ev3dev.org/docs/tutorials/setting-up-wifi-using-the-command-line/
# https://gist.github.com/kylemanna/6930087
# hostname.config contains hostname form mac address table
# delete hostname.config file to re-assign host name

# 2021-09-20 @metl

# variables
check_ip=160.85.104.112  # ping test: zhaw.ch IP
host_name="./host.name"  # file to check, if hostname was changed

echo "${dt} startup"

# wifi
echo
echo "1 wait for ${check_ip}"
while true; do sleep 5; ping -c1 ${check_ip} > /dev/null && break; done

# hostname
if [[ ! -f "${host_name}" ]]; then  # first boot only
    echo
    echo "2 get hostname form git"
    touch ${host_name}
    ev3=$(./get_host_name_by_mac.py)
    echo ${ev3} >> ${host_name}
    hostnamectl set-hostname ${ev3}
    echo " > host name set to '${ev3}'"
fi

# update
download_url=$(curl https://api.github.com/repos/PA-arslasel-machitic/EV3-API/releases/latest | jq -r '.assets[] | select(.name == "server") | .browser_download_url')
curl -o server $download_url
chmod +x server

# robots
echo
echo "4 start rpyc robots"
./server