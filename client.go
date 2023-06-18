package gosrv

type ResponseHandler func(data *Response)

type UpdateHandler func(data *Update)
type UpdateSubscriber func(collection MCollection[UpdateHandler])

type Client interface {
}
type pClient struct {
	inRequestState     bool
	responseCollection *pKCollection[*pRequest, *Response]
	updateCollection   *pMCollection[UpdateHandler]
}

func (c *pClient) UpdateHandler(subscribers ...UpdateSubscriber) {
	for _, subscriber := range subscribers {
		subscriber(c.updateCollection)
	}
}

// =====================

func NewClient() Client {
	responseCollection := newKCollection[*pRequest, *Response](func(k *pRequest) (key string) {
		return k.UUID + "::" + k.Method + "::" + k.Channel
	})

	value := &pClient{
		responseCollection: responseCollection,
		updateCollection:   newMCollection[UpdateHandler](),
	}

	return value
}
