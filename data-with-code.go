package gosrv

import (
	"cyberpull.com/gotk/errors"
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
