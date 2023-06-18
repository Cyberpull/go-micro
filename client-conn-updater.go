package gosrv

type ClientUpdater interface {
	WriteUpdate(data *Update) (n int, err error)
}

// =============================

func (c *pClientConn) WriteUpdate(data *Update) (n int, err error) {
	return writeUpdate(c, data)
}
