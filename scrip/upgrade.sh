#!/bin/bash

if [[ ! -f /etc/systemd/system/algoliteos.service ]]; then
  echo "文件不存在，执行安装脚本..."
  /bin/bash install.sh 
  wait $!
  echo "安装脚本执行完毕，退出脚本。"
  exit 0
fi


systemctl stop algoliteos.service

cp algoliteos /bin/
# cp config/event.yaml /etc/sophliteos/config
# cp config/kaola.json /etc/sophliteos/config

systemctl start algoliteos.service