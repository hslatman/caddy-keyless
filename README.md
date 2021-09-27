# caddy-keyless

A Caddy module providing Keyless SSL support

## Description

This Caddy module provides Keyless SSL support through a custom `keyless` certificate loader.

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

The certificate, key and CA bundle are required for mTLS between Caddy and the Keyless server.
In this case the Keyless server is configured to run on the same host on port 7000.
The array of certificates contain paths to the certificates that 

TODO: provide some Keyless server configuration?

## TODO

* Add more configuration options and/or smarter defaults
* Provide multiple means for loading the certs (from files, from directories, etc); 
    * Reuse the existing certificate loaders for this?
* Provide an example using Docker?
* See other TODOs in code
* ...
