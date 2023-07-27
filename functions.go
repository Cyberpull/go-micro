package gosrv

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"cyberpull.com/gotk/v2/cert"
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

func connect(opts Options) (net.Conn, error) {
	addr := address(opts)

	timeout := time.Second * 30

	if certOpts := opts.getCertOptions(); certOpts != nil {
		config, err := cert.GetTLSConfig(*certOpts)

		if err != nil {
			return nil, err
		}

		if config != nil {
			dialer := &net.Dialer{Timeout: timeout}
			return tls.DialWithDialer(dialer, "tcp", addr, config)
		}
	}

	return net.DialTimeout("tcp", addr, timeout)
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
