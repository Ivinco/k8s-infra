#!/bin/sh
set -x
curl -sSLO https://werf.io/install.sh && chmod +x install.sh
./install.sh --ci
source "$(~/bin/trdl use werf 1.2 stable)"
echo ${WERF_REPO_HARBOR_PASSWORD} |  docker login harbor.sgdctroy.net --username admin --password-stdin
werf version
werf build
werf render --env dev --namespace admission-controller
