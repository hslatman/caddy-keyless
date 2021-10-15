# caddy-keyless

A Caddy module providing [Keyless SSL](https://www.cloudflare.com/ssl/keyless-ssl/) support

## Description

This Caddy module provides Keyless SSL support, bringing this Cloudflare technology into self-hosted environments.

It is based on a custom `keyless` certificate loader that offloads the TLS handshake to a Keyless SSL server.

*This is an early POC; things will change!*

## Configuration

Currently only the Caddy JSON configuration format is supported.
An example TLS configuration is shown below:

```json
    "tls": {
        "certificates": {
            "keyless": {
                "cert": "/path/to/client/cert.pem",
                "key": "/path/to/client/key.pem",
                "ca": "/path/to/cacert.pem",
                "disable_verification": false,
                "server": "127.0.0.1:7000",
                "certificates": [
                    "/path/to/keyless/certificate.crt"
                ]
            }
        }
    }
```

The cert, key and CA bundle are required for mTLS between Caddy and the Keyless server.
It is possible to disable TLS certificate validation, for example when the Keyless server uses a self-signed certificate that is not trusted, but this must not be used in production.
The Keyless server to contact is running on the same host on port 7000.
The certificates array contains paths to the certificates that are loaded by the `keyless` loader.
TLS handshakes destined for hostnames that are in one of those certificates will be performed by the Keyless SSL server.

```shell
$ gokeyless --private-key-dirs "/path/to/private/keys"
```

This module (currently) does not offer a method to automatically retrieve the certificates to serve.
This means that certificates for which the Keyless server manages keys should be made available to the Caddy instance using other means.

## TODO

* Add more configuration options and/or smarter defaults
* Provide multiple means for loading the certs (from files, from directories, from remote, etc); 
    * Reuse the existing certificate loaders for this?
* Implement an CertMagic issuer backed by Keyless SSL?
    * Likely requires a layer on top of the plain Gokeyless server
* Provide an example using Docker?
* Caddyfile support
* See other TODOs in code
* ...
