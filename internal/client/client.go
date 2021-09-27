package client

import (
	"crypto/tls"

	kc "github.com/cloudflare/gokeyless/client"
)

type Client struct {
	c      *kc.Client
	server string
}

func New(server string, certFile string, keyFile string, caFile string) (*Client, error) {

	// TODO: make caFile optional; fallback to system certificate bundle if not provided

	c, err := kc.NewClientFromFile(certFile, keyFile, caFile) // TODO: provide wrapper for the option without file
	if err != nil {
		return nil, err
	}

	// TODO: look into how the multiple remotes are handled by the gokeyless client (remoteCache)

	return &Client{
		c:      c,
		server: server,
	}, nil
}

func (c *Client) DisableVerification() {
	c.c.Config.InsecureSkipVerify = true
}

func (c *Client) LoadTLSCertificate(certFile string) (tls.Certificate, error) {
	return c.c.LoadTLSCertificate(c.server, certFile)
}
