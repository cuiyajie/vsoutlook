#!/bin/bash
set -ex
registry=cuiyajie
image=${registry}/vsoutlook-server

docker pull ${image}:latest
docker tag ${image}:latest ${image}:stable
docker push ${image}:stable

ssh vso kubectl rollout restart deployment app
