#!/bin/bash
env GOOS=linux GOARCH=amd64 go build -v github.com/arastu/servert
