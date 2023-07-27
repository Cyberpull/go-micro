package gosrv

type Data interface {
	// SetContent(v any) (err error)
	ParseContent(v any) (err error)
}

// ======================

type pData struct {
	Content []byte `json:"content" binding:"required"`
}

func (d *pData) SetContent(v any) (err error) {
	data, err := json.Encode(v)

	if err != nil {
		return
	}

	d.Content = data

	return
}

func (d pData) ParseContent(v any) (err error) {
	return json.Decode(d.Content, v)
}
