#!/bin/bash
# install_usb.sh [WOS_INSTALL_PATH [WOS_OS]]
#
# This script will install the latest binary from GitHub to a mounted USB drive
# or directory.
#
# WOS_INSTALL_PATH - The path to the mounted USB drive
# WOS_OS - This argument specifies which build to install (all, darwin64, 
# freebsd64, linux64, or windows64)
################################################################################

VERSION="0.2.3"
WOS_BUILD_DIR="/tmp/wikionastick-$VERSION-build"
CWD=$(pwd)
URL="https://github.com/mikeshultz/wikionastick/releases/download/v$VERSION/wikionastick-$VERSION-all.tar.gz"

if [ "$1" ] && [ -z "$WOS_INSTALL_PATH" ]; then
    WOS_INSTALL_PATH=$1
else
    echo "No install path provided!"
    exit 100
fi
if [ "$2" && -z "$WOS_OS" ]; then
    WOS_OS="$2"
else
    WOS_OS="all"
fi

echo "WikiOnAStick Installer"
echo "======================"
echo "VERSION: $VERSION"
echo "WOS_BUILD_DIR: $WOS_BUILD_DIR"
echo "URL: $URL"
echo "WOS_OS: $WOS_OS"
echo "DESTINATION: $WOS_INSTALL_PATH"
echo "----------------------"
echo ""
echo "Creating build environment..."

mkdir -p $WOS_BUILD_DIR
cd $WOS_BUILD_DIR

echo "Downloading wikionastick-$VERSION-$WOS_OS.tar.gz"
curl -OL $URL -o "wikionastick-$VERSION-$WOS_OS.tar.gz"

echo "Extracting files to destination..."
cd $WOS_INSTALL_PATH
tar -xzf $WOS_BUILD_DIR/wikionastick-$VERSION-$WOS_OS.tar.gz

# Doubley make sure execute perms are set
chmod 755 ./wiki.*

# Return the user to their original CWD
cd $CWD

echo "Cleaning up..."

rm -rf $WOS_BUILD_DIR

echo "Done."
echo "--------------------------------------------------"
echo "You can now run the wiki with a command like this:"
echo ""
echo "    cd $WOS_INSTALL_PATH && ./wiki.linux64"
echo ""
echo "It will be running at:"
echo ""
echo "    http://localhost:8888/"
echo "--------------------------------------------------"