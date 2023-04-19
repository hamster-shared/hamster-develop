#!/bin/sh
set -ex

docker build --no-cache -t docker.io/hamstershare/mysten-sui-tools:v0.30.0 .
