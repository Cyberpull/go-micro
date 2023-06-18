package gosrv

type Update struct {
	pDataWithCode

	Method  string `json:"method" binding:"required"`
	Channel string `json:"channel" binding:"required"`
}

// ==========================

func newUpdate(method, channel string, v any, codes ...int) (data *Update, err error) {
	code := one(200, codes)

	data = &Update{
		pDataWithCode: pDataWithCode{Code: code},
		Method:        method,
		Channel:       channel,
	}

	if err = data.SetContent(v); err != nil {
		data = nil
	}

	return
}
