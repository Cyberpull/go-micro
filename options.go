package gosrv

import (
	"cyberpull.com/gotk/v2/cert"
)

type Options interface {
	getHost() string
	getPort() string
	getCertOptions() *cert.Options
	getInfo() *Info
}
