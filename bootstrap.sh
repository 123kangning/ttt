#!/bin/bash
go build .
echo "">qin.log
./qin >> qin.log &

