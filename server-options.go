package gosrv

import (
	"cyberpull.com/gotk/cert"
)

type ServerOptions struct {
	Host        string
	Port        string
	Name        string
	Alias       string
	CertOptions *cert.Options
}

func (o ServerOptions) getHost() string {
	return o.Host
}

func (o ServerOptions) getPort() string {
	return o.Port
}

func (o ServerOptions) getName() string {
	return o.Name
}

func (o ServerOptions) getAlias() string {
	return o.Alias
}

func (o ServerOptions) getCertOptions() *cert.Options {
	return o.CertOptions
}
