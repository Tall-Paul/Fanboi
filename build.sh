#!bash
#
# Eli Bendersky [https://eli.thegreenplace.net]
# This code is in the public domain.


#set -eux
#set -o pipefail

echo "cleaning build"
rm -rfv ./plugins > /dev/null 2>&1
rm fanboi  > /dev/null 2>&1



mkdir plugins
cd ./src/plugins

echo " "
echo "building plugins"
for f in *; do
    if [ -d "$f" ]; then
        echo "             $f"
        go build -buildmode=plugin ./$f
         mv $f.so ../../plugins
    fi
done


cd ../..
echo " "
echo "building main"
cd ./src
go build fanboi
mv fanboi ..
cd ..
echo " "
echo "done"


