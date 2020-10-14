#!/usr/bin/env bash

set -eu -o pipefail

cd public
mkdir -p static
cd static

rm -rf feather-font-*
curl -fSsL -o /tmp/feather-font.zip https://github.com/AT-UI/feather-font/archive/master.zip
unzip /tmp/feather-font.zip && rm /tmp/feather-font.zip
rm -rf feather-font
mkdir feather-font
mv feather-font-*/src/{css,fonts} feather-font/
rm -rf feather-font-*
curl -fSsL -O https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css
