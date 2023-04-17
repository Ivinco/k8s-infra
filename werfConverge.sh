#!/bin/sh
source "$(/root/bin/trdl use werf 1.2 stable)"
werf version
werf build
