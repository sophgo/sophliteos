#!/usr/bash

cp scrip/sophliteos.service scrip/install.sh  scrip/uninstall.sh scrip/upgrade.sh .

CGO_ENABLED=1 GOOS=linux GOARCH=riscv64  CC=riscv64-linux-gnu-gcc go build -trimpath -ldflags '-s -w'  && tar -zcvf sophliteos-linux_riscv64.tgz sophliteos \
    dist \
    config/sophliteos.yaml \
    database/sophliteos.db \
    sophliteos.service \
    install.sh \
    uninstall.sh \
    release_version.txt \
    upgrade.sh

rm sophliteos.service install.sh  uninstall.sh upgrade.sh 