#!/bin/sh
export PATH=/usr/local/git/bin:$PATH
git --version
source "$(/root/bin/trdl use werf 1.2 stable)"
echo ${WERF_REPO_HARBOR_PASSWORD} |  docker login harbor.sgdctroy.net --username admin --password-stdin
werf version
werf converge
