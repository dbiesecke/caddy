#!/bin/sh

export GOPROXY="https://goproxy.cn,direct"
export CGO_ENABLED=0

set -eu

oss="linux darwin"
archs="amd64"
types="full less basic"

mkdir -p bin
for os in ${oss}; do
    for arch in ${archs}; do
        for tp in ${types}; do
            binName="saddy-${tp}-${os}-${arch}"
            echo "build ${binName}"
            GOOS=${os} GOARCH=${arch} go build -ldflags '-s -w' -v -o ${binName} ./${tp}
            tar zcf ${binName}.tar.gz ${binName}
            mv ${binName}* bin/
        done
    done
done
