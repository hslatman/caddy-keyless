#!/bin/bash
set -e -v

killall caddy-keyless || true
killall gokeyless || true
sleep 2s

(cd .. && make build) && cp ../build/caddy-keyless .
(cd ../gokeyless && make gokeyless) && cp ../gokeyless/gokeyless .

bash ./prepare-files.sh

tmux new-session -s "demo-caddy-keyless" -d "./caddy-keyless run --config ./caddy.json"
tmux new-session -s "demo-gokeyless" -d "./gokeyless --private_key_files ./keyless-keys/server.key --port 7000 --hostname localhost.localdomain --auth_key ./server-client-comm/server.key --auth_cert ./server-client-comm/server.crt --cloudflare_ca_cert ./server-client-comm/ca.crt"
sleep 1s
curl --cacert ./keyless-browser/ca.crt https://127.0.0.1
