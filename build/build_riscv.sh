#!/bin/sh
set -e

current_directory=${PWD##*/}

# 检查当前目录是否为"build"
if [ "$current_directory" != "build" ]; then
  echo "错误：该脚本必须在build目录中执行。"
  exit 1
fi

sh version.sh "V1.1.2"
mv release_version.txt ../

# docker run --rm -i --name node-build -v `pwd`/../frontend/:/home/node node:16 sh -c 'cd /home/node/sophliteos-frontend && yarn && yarn build'
cp -r ../frontend/sophliteos-frontend/dist ../

cd ..
sh ./scrip/package-riscv.sh

cd build

mv ../sophliteos-linux_riscv64.tgz ../release

rm -rf ../dist tmp ../release_version.txt ../sophliteos 