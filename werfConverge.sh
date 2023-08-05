#!/bin/sh
set -x
curl -sSL https://werf.io/install.sh | bash -s -- --ci

#export PATH=/root/bin:/usr/local/git/bin:$PATH
#git --version
#trdl add werf https://tuf.werf.io 1 b7ff6bcbe598e072a86d595a3621924c8612c7e6dc6a82e919abe89707d7e3f468e616b5635630680dd1e98fc362ae5051728406700e6274c5ed1ad92bea52a2
#source "$(trdl use werf 1.2 stable)"
source $("/home/go/bin/trdl" use werf "1.2" "stable")
werf version
werf cr login -u admin -p ${WERF_REPO_HARBOR_PASSWORD} harbor.sgdctroy.net
werf render
werf build
sleep 1200
