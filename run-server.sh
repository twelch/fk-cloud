#!/bin/bash

export FIELDKIT_ADDR=:8080
export FIELDKIT_POSTGRES_URL=postgres://fieldkit:password@127.0.0.1/fieldkit?sslmode=disable
export FIELDKIT_DISABLE_MEMORY_LOGGING=true
export FIELDKIT_DISABLE_STARTUP_REFRESH=true
export FIELDKIT_API_DOMAIN=127.0.0.1:8080
export FIELDKIT_API_HOST=http://127.0.0.1:8080
export FIELDKIT_HTTP_SCHEME=http

if [ -f aws.env ]; then
    source aws.env
fi

set -xe

make server

build/server "$@"
