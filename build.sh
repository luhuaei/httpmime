#!/usr/bin/sh

CGO_ENABLED=0 go build

./httpmime --help

cp httpmime ~/.local/bin
