#!/bin/bash
set -ex
bash infra/utils/update-build-info.sh

mkdir -p .cache
cd .cache
upx_version=4.0.1
curl -C - -LO https://github.com/upx/upx/releases/download/v${upx_version}/upx-${upx_version}-amd64_linux.tar.xz
tar xvJf upx-${upx_version}-amd64_linux.tar.xz --strip-components=1 upx-${upx_version}-amd64_linux/upx
chmod +x upx

cd ..
docker build -t app-server -f Dockerfile .
registry=dockerproxy.com/cuiyajie/vsoutlook

docker tag app-server $registry:latest
docker push $registry:latest

commit=$(git rev-parse HEAD | cut -c1-8)
docker tag app-server $registry:$commit
docker push $registry:$commit
