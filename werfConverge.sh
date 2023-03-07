#!/bin/env sh
source "$(~/bin/trdl use werf 1.2 stable)"
werf version
werf converge
