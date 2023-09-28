#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o bin/wordlistidea-linux-amd64
GOOS=windows GOARCH=amd64 go build -o bin/wordlistidea-windows-amd64.exe

scp bin/wordlistidea-linux-amd64 bjones@i3-1:~/wordlistidea-linux-amd64
