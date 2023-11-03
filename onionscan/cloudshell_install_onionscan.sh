#!/bin/bash

go install github.com/burntonions/onionscan3/onionscan@latest
cd gopath
cp -r `find ~+ -type d -name "templates"` $HOME
cd -
