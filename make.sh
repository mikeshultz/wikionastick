#!/bin/bash

# Version should be changed on each build!
VERSION="0.1.1"

CWD=$(pwd)
BUILD_DIR="$CWD/build"

# Build directories, filenames, and other per-arch settings

# Windows
WIN64_DIR="$BUILD_DIR/win64"
WIN64="$WIN64_DIR/wiki.exe"
WIN64_TARBALL="wikionastick-$VERSION-windows64.tar.gz"

# Linux
LINUX64_DIR="$BUILD_DIR/linux64"
LINUX64="$LINUX64_DIR/wiki.linux64"
LINUX64_TARBALL="wikionastick-$VERSION-linux64.tar.gz"

# FreeBSD
FBSD64_DIR="$BUILD_DIR/freebsd64"
FBSD64="$FBSD64_DIR/wiki.freebsd64"
FBSD64_TARBALL="wikionastick-$VERSION-freebsd64.tar.gz"

# Darwin
DARWIN_DIR="$BUILD_DIR/darwin64"
DARWIN="$DARWIN_DIR/wiki.app"
DARWIN_TARBALL="wikionastick-$VERSION-darwin64.tar.gz"

# Combined
COMBINED_DIR="$BUILD_DIR/combined"
COMBINED_TARBALL="wikionastick-$VERSION-all.tar.gz"

echo "go getting the dependencies..."

go get -t github.com/mikeshultz/wikionastick

echo "Creating build environment..."

rm -rf $BUILD_DIR/*

# Starting with the templates dir because fewer commands
mkdir -p $LINUX64_DIR/templates
mkdir -p $WIN64_DIR/templates
mkdir -p $FBSD64_DIR/templates
mkdir -p $DARWIN_DIR/templates
mkdir -p COMBINED_DIR

echo "Building Linux amd64..."
GOOS=linux GOARCH=amd64 go build -o $LINUX64

echo "Building Windows amd64..."
GOOS=windows GOARCH=amd64 go build -o $WIN64

echo "Building FreeBSD amd64..."
GOOS=freebsd GOARCH=amd64 go build -o $FBSD64

echo "Building Darwin(OSX) amd64..."
GOOS=darwin GOARCH=amd64 go build -o $DARWIN

echo "Copying templates..."

cp -R templates $WIN64_DIR
cp -R templates $LINUX64_DIR
cp -R templates $FBSD64_DIR
cp -R templates $DARWIN_DIR
cp -R templates $COMBINED_DIR

echo "Generating tarballs..."

cd $WIN64_DIR && tar -czf $BUILD_DIR/$WIN64_TARBALL . && cd $CWD
cd $LINUX64_DIR && tar -czf $BUILD_DIR/$LINUX64_TARBALL . && cd $CWD
cd $FBSD64_DIR && tar -czf $BUILD_DIR/$FBSD64_TARBALL . && cd $CWD
cd $DARWIN_DIR && tar -czf $BUILD_DIR/$DARWIN_TARBALL . && cd $CWD
cd $COMBINED_DIR && cp ../*/wiki.* $COMBINED_DIR/ && tar -czf $BUILD_DIR/$COMBINED_TARBALL . && cd $CWD

echo "Done."