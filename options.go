package gosrv

import (
	"cyberpull.com/gotk/cert"
)

type Options interface {
	getHost() string
	getPort() string
	getCertOptions() *cert.Options
	getInfo() *Info
}
