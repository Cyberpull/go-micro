package gosrv

type NetReader interface {
	Read(delim byte) ([]byte, error)
}

// ============================

func (io netIO) Read(b []byte) (n int, err error) {
	return io.reader.Read(b)
}

func (io netIO) ReadByte() (b byte, err error) {
	return io.reader.ReadByte()
}

func (io netIO) ReadBytes(delim byte) (b []byte, err error) {
	return io.reader.ReadBytes(delim)
}

func (io netIO) ReadLine(delim byte) (line []byte, isPrefix bool, err error) {
	return io.reader.ReadLine()
}

func (io netIO) ReadRune() (r rune, size int, err error) {
	return io.reader.ReadRune()
}

func (io netIO) ReadSlice(delim byte) (line []byte, err error) {
	return io.reader.ReadSlice(delim)
}

func (io netIO) ReadString(delim byte) (s string, err error) {
	return io.reader.ReadString(delim)
}
