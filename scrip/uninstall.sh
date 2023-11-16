#!/bin/sh

systemctl stop algoliteos
systemctl disable algoliteos
rm -rf  /etc/systemd/system/algoliteos.service  /bin/algoliteos  /data/pictures
rm -rf  /etc/sophliteos/config/event*  /var/lib/sophliteos/db/algoliteos.db /etc/sophliteos/config/algoliteos*

systemctl restart sophliteos.service