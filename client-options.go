package gosrv

import (
	"cyberpull.com/gotk/cert"
)

type ClientOptions struct {
	ServerHost  string `binding:"required"`
	ServerPort  string `binding:"required"`
	Info        *Info  `binding:"required"`
	CertOptions *cert.Options
}

func (o ClientOptions) getHost() string {
	return o.ServerHost
}

func (o ClientOptions) getPort() string {
	return o.ServerPort
}

func (o ClientOptions) getInfo() *Info {
	return o.Info
}

func (o ClientOptions) getCertOptions() *cert.Options {
	return o.CertOptions
}
