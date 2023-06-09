package gosrv

import "crypto/tls"

type ServerOptions struct {
	Host      string
	Port      string
	Name      string
	Alias     string
	TlsConfig *tls.Config
}

func (o ServerOptions) GetHost() string {
	return o.Host
}

func (o ServerOptions) GetPort() string {
	return o.Port
}

func (o ServerOptions) GetName() string {
	return o.Name
}

func (o ServerOptions) GetAlias() string {
	return o.Alias
}

func (o ServerOptions) GetTlsConfig() *tls.Config {
	return o.TlsConfig
}

func (o *ServerOptions) SetTlsConfig(config *tls.Config) {
	o.TlsConfig = config
}
