package gosrv

import "crypto/tls"

type Options interface {
	GetHost() string
	GetPort() string
	GetName() string
	GetAlias() string
	GetTlsConfig() *tls.Config
	SetTlsConfig(config *tls.Config)
}
