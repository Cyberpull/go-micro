package gosrv

import "cyberpull.com/gotk/objects"

type BaseData interface {
	SetData(v any) (err error)
	ParseData(v any) (err error)
}

// ======================

type pBaseData struct {
	Data []byte `json:"data"`
}

func (b *pBaseData) SetData(v any) (err error) {
	data, err := objects.ToJSON(v)

	if err != nil {
		return
	}

	b.Data = data

	return
}

func (b *pBaseData) ParseData(v any) (err error) {
	return objects.ParseJSON(b.Data, v)
}
