package gosrv

import "crypto/tls"

type ClientOptions struct {
	ServerHost string
	ServerPort string
	Name       string
	Alias      string
	TlsConfig  *tls.Config
}

func (o ClientOptions) GetHost() string {
	return o.ServerHost
}

func (o ClientOptions) GetPort() string {
	return o.ServerPort
}

func (o ClientOptions) GetName() string {
	return o.Name
}

func (o ClientOptions) GetAlias() string {
	return o.Alias
}

func (o ClientOptions) GetTlsConfig() *tls.Config {
	return o.TlsConfig
}

func (o *ClientOptions) SetTlsConfig(config *tls.Config) {
	o.TlsConfig = config
}
