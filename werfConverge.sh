#!/bin/env sh
source "$(~/bin/trdl use werf 1.2 stable)"
werf version
export WERF_SET_APF_ENABLED="apf.enabled=true"
werf converge
