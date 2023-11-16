#!/usr/bash

cp scrip/algoliteos.service scrip/install.sh  scrip/uninstall.sh scrip/upgrade.sh .

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags '-s -w' && tar -zcvf algoliteos-linux_amd64.tgz algoliteos \
    algoliteos.service \
    config/algoliteos.yaml \
    config/events.yaml \
    config/event.json \
    database/algoliteos.db \
    install.sh \
    uninstall.sh \
    upgrade.sh
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -trimpath -ldflags '-s -w' && tar -zcvf algoliteos-linux_arm64.tgz algoliteos \
    algoliteos.service \
    config/algoliteos.yaml \
    config/events.yaml \
    config/event.json \
    database/algoliteos.db \
    install.sh \
    uninstall.sh \
    upgrade.sh

rm algoliteos.service install.sh  uninstall.sh upgrade.sh 