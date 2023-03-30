#!/bin/bash

if [ "$EUID" -ne 0 ]
  then echo "Please run with sudo"
  exit
fi

./build.sh
install -Dm755 cli64 /usr/bin/cli64