#!/bin/bash

set -e

if [[ -f /etc/systemd/system/algoliteos.service ]]; then
  systemctl stop algoliteos.service
  systemctl disable algoliteos.service
fi
rm -rf /data/pictures
mkdir -p /etc/sophliteos/config /var/log/sophliteos /data/pictures

cp algoliteos /bin
cp config/algoliteos.yaml /etc/sophliteos/config
cp config/events.yaml /etc/sophliteos/config
cp config/event.json /etc/sophliteos/config
cp database/algoliteos.db /var/lib/sophliteos/db
cp algoliteos.service /etc/systemd/system/

active_interface=$(ip route get 8.8.8.8 | awk 'NR==1 {print $5}')
if [[ ! -z $active_interface ]];then
    LOCAL_IP=`ip -4 addr show dev "$active_interface" | grep -oP '(?<=inet\s)\d+(\.\d+){3}'`
    sed -i "s/upload: 127.0.0.1:8081/upload: $LOCAL_IP:8081/" /etc/sophliteos/config/algoliteos.yaml
fi

systemctl daemon-reload
systemctl enable algoliteos.service
systemctl start algoliteos.service