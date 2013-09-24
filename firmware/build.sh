#!/bin/sh
set -xe

# remove Arduino IDE build files
rm -rf src/main/build-uno/

VERSION=`git describe --tags --exact-match || git rev-parse --short HEAD`
sed -e "s/__VERSION__/$VERSION/" src/main/version.template > src/main/version.c
ino build -m leonardo
