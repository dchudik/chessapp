#!/bin/bash

cd ../

godoc -http=":6060" & pid="$!"

sleep 5

cd resources/

wget \
    --recursive \
    --no-verbose \
    --convert-links \
    --page-requisites \
    --adjust-extension \
    --execute=robots=off \
    --directory-prefix="docs" \
    --exclude-directories="*" \
    --no-host-directories \
    "http://localhost:6060/pkg/github.com/dchudik/chessapp/chesslib" || status=$?

kill -9 "${pid}"

zip -r docs.zip docs

rm -r docs