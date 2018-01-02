#! /bin/bash
set -e
source ~/.bashrc

BIN_DIR=$(dirname $0)
BIN_DIR=$(cd $BIN_DIR;pwd)
ROOT_DIR=$(cd $BIN_DIR/..;pwd)
cd $ROOT_DIR

export GOPATH=$ROOT_DIR/..
export GOBIN=$GOPATH/bin

if [[ -d /home/compile/makepkg_go/go1.9  ]]; then
    export GOROOT=/home/compile/makepkg_go/go1.9
fi

export PATH=$GOROOT/bin:$PATH

cd ../src
go install qwh/main/main.go

if [ $? -eq 0 ]; then
	echo "[ok] go compile success"
else
	echo "[error] go compile failure"
	exit 1
fi

