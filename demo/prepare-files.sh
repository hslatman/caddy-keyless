#!/bin/bash
set -e

mkdir server-client-comm/ || true
mkdir keyless-keys/ || true
mkdir keyless-certs/ || true
mkdir keyless-browser/ || true

bash ./gen-test-cert-ecc.sh # both rsa and ecc works
cp ca.crt server.key server.crt client.key client.crt server-client-comm/

bash ./gen-test-cert-rsa2048.sh # ecc does not work. I don't know why
cp server.key keyless-keys/
cp server.crt keyless-certs/
cp ca.crt keyless-browser/

rm ca.crt ca.key ca.srl server.csr server.key server.crt client.csr client.key client.crt || true