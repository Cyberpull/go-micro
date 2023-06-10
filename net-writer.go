package gosrv

type NetWriter interface {
	Write(b []byte) (n int, err error)
	WriteLine(b []byte) (n int, err error)
	WriteString(s string) (n int, err error)
	WriteStringLine(s string) (n int, err error)
}

// ============================

func (io *netIO) Write(b []byte) (n int, err error) {
	return io.conn.Write(b)
}

func (io *netIO) WriteLine(b []byte) (n int, err error) {
	return io.Write(append(b, '\n'))
}

func (io *netIO) WriteString(s string) (n int, err error) {
	return io.Write([]byte(s))
}

func (io *netIO) WriteStringLine(s string) (n int, err error) {
	return io.WriteLine([]byte(s))
}
