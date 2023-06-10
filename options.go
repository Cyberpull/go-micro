package gosrv

import (
	"cyberpull.com/gotk/cert"
)

type Options interface {
	getHost() string
	getPort() string
	getName() string
	getAlias() string
	getCertOptions() *cert.Options
}
