#/bin/bash

DATE=$(date)
GOVERSION=$(go version)
VERSION=$(git describe --tags --abbrev=8 --dirty --always --long)

echo "$VERSION - $GOVERSION - $DATE" > infra/utils/build-info.txt
