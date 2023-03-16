#!/bin/sh
source "$(/root/bin/trdl use werf 1.2 stable --no-self-update)"
werf version
werf converge
