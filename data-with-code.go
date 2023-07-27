package gosrv

import (
	"cyberpull.com/gotk/v2/errors"
)

type DataWithCode interface {
	Data

	GetError() (err error)
}

// ======================

type pDataWithCode struct {
	pData

	Code int `json:"code"`
}

func (d pDataWithCode) GetError() (err error) {
	if d.Code >= 200 && d.Code < 300 {
		return
	}

	var message string

	if err = d.ParseContent(&message); err != nil {
		return
	}

	err = errors.New(message, d.Code)

	return
}

func (d *pDataWithCode) SetError(data any, code ...int) (err error) {
	var message string

	d.Code = one[int](500, code)

	switch x := data.(type) {
	case Data:
		err = x.ParseContent(&message)

		if err != nil {
			return
		}

	default:
		e := errors.From(data, code...)
		message, d.Code = e.Error(), e.Code()
		err = d.SetContent(e.Error())
	}

	return
}
