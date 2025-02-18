#!bash
#
# Eli Bendersky [https://eli.thegreenplace.net]
# This code is in the public domain.

docker run -it --rm \
  -v /Users/briggpw1/Documents/Fanboi:/var/Fanboi \
  -w /var/Fanboi \
  -e CGO_ENABLED=1 \
  --platform linux/amd64 \
    docker.elastic.co/beats-dev/golang-crossbuild:1.24.0-main \
  --build-cmd "make build" \
  -p "linux/amd64" \
 


