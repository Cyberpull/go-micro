package gosrv

import (
	"cyberpull.com/gotk/errors"
)

type Output interface {
	BaseData

	GetError() (err error)
}

// ======================

type pOutput struct {
	pBaseData

	Method  string `json:"method"`
	Channel string `json:"channel"`
	Code    int    `json:"code"`
}

func (o *pOutput) GetError() (err error) {
	if o.Code >= 200 && o.Code < 300 {
		return
	}

	var message string

	if err = o.ParseData(&message); err != nil {
		return
	}

	err = errors.New(message, o.Code)

	return
}

// ======================

func newOutput() *pOutput {
	return &pOutput{}
}
