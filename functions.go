package gosrv

import (
	"crypto/tls"
	"fmt"
	"net"

	"cyberpull.com/gotk/cert"
)

func address(opts Options) string {
	return fmt.Sprintf(`%s:%s`, opts.getHost(), opts.getPort())
}

func listen(opts Options) (net.Listener, error) {
	addr := address(opts)

	if certOpts := opts.getCertOptions(); certOpts != nil {
		config, err := cert.GetTLSConfig(*certOpts)

		if err != nil {
			return nil, err
		}

		if config != nil {
			return tls.Listen("tcp", addr, config)
		}
	}

	return net.Listen("tcp", addr)
}

func one[T any](def T, args []T) T {
	if len(args) >= 1 {
		return args[0]
	}

	return def
}

// ERROR =============================

func sendOne[T any](out []chan T, data T) {
	if len(out) > 0 {
		send(out[0], data)
	}
}

func send[T any](out chan T, data T) {
	defer recover()

	if out == nil {
		return
	}

	out <- data
}
