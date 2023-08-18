#!/bin/sh
export PATH=/root/bin:/usr/local/git/bin:$PATH
git --version
echo ${WERF_REPO_HARBOR_PASSWORD} |  docker login harbor.sgdctroy.net --username admin --password-stdin
werf version
werf build
