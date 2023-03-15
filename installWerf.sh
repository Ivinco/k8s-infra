#!/bin/env sh
curl -sSLO https://werf.io/install.sh
chmod +x install.sh
./install.sh --ci
source "$(~/bin/trdl use werf 1.2 stable)"
werf version

