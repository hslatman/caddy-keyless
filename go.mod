module github.com/hslatman/caddy-keyless

go 1.16

replace github.com/cloudflare/gokeyless => ./gokeyless

require (
	github.com/caddyserver/caddy/v2 v2.4.5
	github.com/cloudflare/gokeyless v1.6.6
	go.step.sm/crypto v0.11.0 // indirect
	go.uber.org/zap v1.19.0
)
