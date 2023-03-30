#!/bin/bash
VERSION=0.1
APP_NAME=cli64
go build -ldflags "-X main.Version=$VERSION" -o $APP_NAME *.go