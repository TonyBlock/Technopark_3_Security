#!/bin/sh

mkdir certs/
#cd certs
openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/CN=golang proxy CA"
openssl genrsa -out cert.key 2048
sudo cp ca.crt /usr/local/share/ca-certificates/
sudo update-ca-certificates

#!/bin/sh
#
#openssl genrsa -out ca.key 2048
#openssl req -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/CN=yngwie proxy CA"
#openssl genrsa -out cert.key 2048
#mkdir certs/
