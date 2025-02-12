#!/usr/bin/bash
#
# Eli Bendersky [https://eli.thegreenplace.net]
# This code is in the public domain.

set -eux
set -o pipefail

rm -rfv ./plugins
mkdir plugins
cd plugins
go build -buildmode=plugin ../src/plugins/unraid_drives/
