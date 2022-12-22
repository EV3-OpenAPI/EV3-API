#!/bin/bash
# https://www.ev3dev.org/docs/tutorials/setting-up-wifi-using-the-command-line/
# https://gist.github.com/kylemanna/6930087
# hostname.config contains hostname form mac address table
# delete hostname.config file to re-assign host name

# 2021-09-20 @metl
# 2022-12-22 @machitic

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
    ev3=$(./ev3api-server -get-hostname)
    echo ${ev3} >> ${host_name}
    hostnamectl set-hostname ${ev3}
    echo " > host name set to '${ev3}'"
fi

# robots
echo
echo "3 start ev3api-server"
./ev3api-server -update