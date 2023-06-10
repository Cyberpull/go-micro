package gosrv

import (
	"cyberpull.com/gotk/cert"
)

type ClientOptions struct {
	ServerHost  string
	ServerPort  string
	Name        string
	Alias       string
	CertOptions *cert.Options
}

func (o ClientOptions) getHost() string {
	return o.ServerHost
}

func (o ClientOptions) getPort() string {
	return o.ServerPort
}

func (o ClientOptions) getName() string {
	return o.Name
}

func (o ClientOptions) getAlias() string {
	return o.Alias
}

func (o ClientOptions) getCertOptions() *cert.Options {
	return o.CertOptions
}
