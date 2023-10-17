#!/bin/bash
set -e

CA_CN="ca.localdomain"
SERVER_CN="127.0.0.1"
CLIENT_CN="127.0.0.1"

rm ca.srl || true
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -sha256 -days 1825 -out ca.crt -subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=${CA_CN}"

openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=${SERVER_CN}"
openssl x509 -req -days 1825 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt

openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/C=US/ST=Denial/L=Springfield/O=Dis/CN=${CLIENT_CN}"
openssl x509 -req -days 1825 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt
