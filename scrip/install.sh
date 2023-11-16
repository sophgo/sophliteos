#!/bin/bash

set -e


if [[ -f /etc/systemd/system/sophliteos.service ]]; then
  systemctl stop sophliteos.service
  systemctl disable sophliteos.service
fi
mkdir -p /etc/sophliteos/config /var/log/sophliteos /var/lib/sophliteos/db /data/sophliteos
rm -rf /var/lib/sophliteos/dist

cp -r dist /var/lib/sophliteos/
cp sophliteos /bin
cp config/sophliteos.yaml /etc/sophliteos/config
cp database/sophliteos.db /var/lib/sophliteos/db
cp sophliteos.service /etc/systemd/system/
cp release_version.txt /var/lib/sophliteos

systemctl daemon-reload
systemctl enable sophliteos.service
systemctl start sophliteos.service