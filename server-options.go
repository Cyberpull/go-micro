package gosrv

import (
	"cyberpull.com/gotk/cert"
)

type ServerOptions struct {
	Host        string `binding:"required"`
	Port        string `binding:"required"`
	Info        *Info  `binding:"required"`
	CertOptions *cert.Options
}

func (o ServerOptions) getHost() string {
	return o.Host
}

func (o ServerOptions) getPort() string {
	return o.Port
}

func (o ServerOptions) getInfo() *Info {
	return o.Info
}

func (o ServerOptions) getCertOptions() *cert.Options {
	return o.CertOptions
}
