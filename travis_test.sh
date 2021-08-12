#!/bin/bash
set -e

go version
go test -v github.com/sgtsquiggs/appengine/...
go test -v -race github.com/sgtsquiggs/appengine/...
if [[ $GOAPP == "true" ]]; then
  export PATH="$PATH:/tmp/sdk/go_appengine"
  export APPENGINE_DEV_APPSERVER=/tmp/sdk/go_appengine/dev_appserver.py
  goapp version
  goapp test -v github.com/sgtsquiggs/appengine/...
fi
