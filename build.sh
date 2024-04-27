#!/usr/bin/env sh



if [ "$#" -gt 0 ]; then
    BINARY=tfgen
fi

go build -o /opt/bin/${BINARY} .
