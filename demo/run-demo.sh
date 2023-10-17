#!/bin/bash
set -e -v

# check gokeyless and tmux dependencies are installed and available to run
# install gokeyless from a package: https://github.com/cloudflare/gokeyless#installing
# install gokeyless from source: go install github.com/cloudflare/gokeyless/cmd/gokeyless@v1.6.13
command -v gokeyless >/dev/null 2>&1 || { echo >&2 "gokeyless is required, but it's not installed. Aborting."; exit 1; }
command -v tmux >/dev/null 2>&1 || { echo >&2 "tmux is required, but it's not installed. Aborting."; exit 1; }

# cleanup previous runs
killall caddy-keyless || true
killall gokeyless || true # not the nicest way, but it works for demo purposes
sleep 2

# ensure latest caddy-keyless (demo) build
(cd .. && make build) && cp ../build/caddy-keyless .

# prepare certificates and keys in the right locations
bash ./prepare-files.sh

# start the caddy-keyless demo instance and gokeyless 
tmux new-session -s "demo-caddy-keyless" -d "./caddy-keyless run --config ./caddy.json"
tmux new-session -s "demo-gokeyless" -d "gokeyless --private_key_files ./keyless-keys/server.key --port 7123 --hostname localhost.localdomain --auth_key ./server-client-comm/server.key --auth_cert ./server-client-comm/server.crt --cloudflare_ca_cert ./server-client-comm/ca.crt"
sleep 2

# try setting up a HTTPS connection to caddy-keyless
curl --connect-timeout 5 --cacert ./keyless-browser/ca.crt https://127.0.0.1
