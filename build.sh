#!bash
#
# Eli Bendersky [https://eli.thegreenplace.net]
# This code is in the public domain.

set -eux
set -o pipefail

rm -rfv ./plugins
rm fanboi



mkdir plugins
cd ./src/plugins
go build -buildmode=plugin ./unraiddrives
mv unraiddrives.so ../../plugins
cd ../..

cd ./src
go build fanboi
mv fanboi ..
cd ..


