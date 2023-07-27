package gosrv

type Update interface {
	DataWithCode
}

type pUpdate struct {
	pDataWithCode

	Method  string `json:"method" binding:"required"`
	Channel string `json:"channel" binding:"required"`
}

// ==========================

func newUpdate(method, channel string, v any, codes ...int) (data *pUpdate, err error) {
	code := one(200, codes)

	data = &pUpdate{
		pDataWithCode: pDataWithCode{Code: code},
		Method:        method,
		Channel:       channel,
	}

	if err = data.SetContent(v); err != nil {
		data = nil
	}

	return
}
