#!/bin/bash

if [[ -z $1 ]]; then
EXT=""
else
EXT="-${1}"
fi

cp /build/imds-debug /target/imds-debug${EXT}
