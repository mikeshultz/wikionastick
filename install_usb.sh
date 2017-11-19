#!/bin/bash
# install_usb.sh PATH [OS]
#
# This script will install the latest binary from GitHub to a mounted USB drive
# or directory.
#
# PATH - The path to the mounted USB drive
# OS - This argument specifies which build to install (all, darwin64, freebsd64,
#       linux64, or windows64)
################################################################################

VERSION="0.2.3"
WOS_BUILD_DIR="/tmp/wikionastick-$VERSION-build"
CWD=$(pwd)
URL="https://github.com/mikeshultz/wikionastick/releases/download/v$VERSION/wikionastick-$VERSION-all.tar.gz"

[ -z "$1" ] && exit 100
if [ "$2" ]; then
    OS="$2"
else
    OS="all"
fi

echo "WikiOnAStick Installer"
echo "======================"
echo "VERSION: $VERSION"
echo "WOS_BUILD_DIR: $WOS_BUILD_DIR"
echo "URL: $URL"
echo "OS: $OS"
echo "DESTINATION: $1"
echo "----------------------"
echo ""
echo "Creating build environment..."

mkdir -p $WOS_BUILD_DIR
cd $WOS_BUILD_DIR

echo "Downloading wikionastick-$VERSION-$OS.tar.gz"
wget $URL

echo "Extracting files to destination..."
tar -xzf $WOS_BUILD_DIR/wikionastick-$VERSION-$OS.tar.gz

# Return the user to their original CWD
cd $CWD

echo "Cleaning up..."

rm -rf $WOS_BUILD_DIR

echo "Done."