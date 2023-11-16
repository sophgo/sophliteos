#!/bin/bash

# 函数用于检查IP地址是否合法
check_ip() {
  local ip="$1"
  local regex="^([0-9]{1,3}\.){3}[0-9]{1,3}$"

  if [[ $ip =~ $regex ]]; then
    IFS='.' read -ra ip_parts <<< "$ip"
    valid=true

    for part in "${ip_parts[@]}"; do
      if ((part < 0 || part > 255)); then
        valid=false
        break
      fi
    done

    if [ "$valid" = true ]; then
      echo "IP地址 $ip。"
      return 0  # 返回0表示真
    else
      echo "IP地址 $ip 不合法，每个段的取值范围应在0-255之间。"
      return 1 
    fi
  else
    echo "IP地址 $ip 不合法，格式不正确。"
    return 1 
  fi
}

# 输入IP地址
read -p "请输入IP地址: " ip

# 调用函数检查IP地址
check_ip "$ip"

result=$?

if [ $result -ne 0 ]; then
  exit 1
fi

sed -i '/clServer/d' /etc/hosts
sed -i '/mediagateway/d' /etc/hosts
echo $ip clServer >> /etc/hosts
echo $ip mediagateway >> /etc/hosts

sudo docker load -i license_service.tar.gz
sudo docker load -i media_gateway.tar.gz
sudo docker load -i device_detector.tar.gz
sudo docker load -i mysql.tar.gz

sudo docker run -d -p 4406:3306 --name media-mysql -e MYSQL_ROOT_PASSWORD='Bitmain_(root123)'  ubuntu/mysql:latest

sudo docker run -itd --name media_gateway --net host media_gateway -bind_ip $ip -db_ip $ip -db_port 4406 -db_user root -db_pass 'Bitmain_(root123)'

sudo docker run -itd  --name device_detector --net host device_detector

sudo docker run -itd --name license_service -p 26085:10085 license_service