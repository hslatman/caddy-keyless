// Copyright 2021 Herman Slatman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keyless

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddytls"

	"go.uber.org/zap"

	"github.com/hslatman/caddy-keyless/internal/client"
)

func init() {
	caddy.RegisterModule(KeylessLoader{})
}

type KeylessLoader struct {
	CertificateFile string `json:"cert"` // TODO: nest inside some kind of Auth block?
	KeyFile         string `json:"key"`
	CAFile          string `json:"ca"`

	DisableVerification bool `json:"disable_verification"`

	Server           string   `json:"server"`
	CertificateFiles []string `json:"certificates"`

	logger *zap.Logger
	client *client.Client
}

// CaddyModule returns the Caddy module information.
func (KeylessLoader) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "tls.certificates.keyless",
		New: func() caddy.Module { return new(KeylessLoader) },
	}
}

// Provision sets up the handler.
func (m *KeylessLoader) Provision(ctx caddy.Context) error {

	m.logger = ctx.Logger(m)
	defer m.logger.Sync()

	c, err := client.New(m.Server, m.CertificateFile, m.KeyFile, m.CAFile) // TODO: add options?
	if err != nil {
		return err
	}

	m.client = c

	if m.DisableVerification {
		m.client.DisableVerification()
	}

	// TODO: add active validation (i.e. try to contact the server, see if it knows the key, etc)?
	// TODO: add support for Caddyfile (shouldn't be too hard at this time; not that much to configure right now)
	// TODO: provide a means for automatically knowing which certificates to use with keyless?

	return nil
}

func (m *KeylessLoader) LoadCertificates() ([]caddytls.Certificate, error) {
	// TODO: provide a custom layer on top of keyless to get the certificate from
	// the keyless server too? Although it can be done manually, for now ...

	certs := []caddytls.Certificate{}

	// TODO: rewrite into actual loader that can read certs multiple ways (dir, files, pems, remote?, etc)

	for _, certFile := range m.CertificateFiles {
		cert, err := m.loadCertificate(certFile) // TODO: call this with an identifier/file for the cert
		if err != nil {
			return []caddytls.Certificate{}, err
		}

		if cert != nil {
			// fmt.Println("appending the cert") // TODO: do proper logging for this
			// fmt.Println(fmt.Sprintf("%#+v", cert))
			// fmt.Println(fmt.Sprintf("%#+v", cert.Certificate))
			// fmt.Println(fmt.Sprintf("%T", cert.Certificate.PrivateKey))

			// leaf := cert.Certificate.Leaf
			// fmt.Println(fmt.Sprintf("%#+v", leaf))
			// fmt.Println(leaf.Subject)
			// fmt.Println(leaf.DNSNames)
			certs = append(certs, *cert)
		}
	}

	return certs, nil
}

func (m *KeylessLoader) loadCertificate(certFile string) (*caddytls.Certificate, error) {
	tlsCert, err := m.client.LoadTLSCertificate(certFile) // TODO: ensure SNI is also sent to remote, if possible?
	cert := &caddytls.Certificate{
		Certificate: tlsCert,
		Tags:        []string{}, // TODO: add some useful info here (look into how Caddy uses these)
	}
	return cert, err
}

// Interface guards
var (
	_ caddy.Module               = (*KeylessLoader)(nil)
	_ caddy.Provisioner          = (*KeylessLoader)(nil)
	_ caddytls.CertificateLoader = (*KeylessLoader)(nil)
)
