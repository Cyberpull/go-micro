package gosrv

import "cyberpull.com/gotk/v2/errors"

type ResponseHandler func(data *pResponse)

type UpdateHandler func(data Update)
type UpdateSubscriber func(collection MCollection[UpdateHandler])

// =====================

type Client interface {
	UpdateHandler(subscribers ...UpdateSubscriber)
	On(method, channel string, handler UpdateHandler)
	SendRequest(method, channel string, data any) (resp Response, err error)
	Start(errChan ...chan error)
	Stop() error
}

// =====================

type pClient struct {
	opts               ClientOptions
	server             *pServerConn
	responseCollection *pKCollection[*pRequest, *pResponse]
	updateCollection   *pMCollection[UpdateHandler]
	shouldStop         bool
	inRequestState     bool
}

func (c *pClient) UpdateHandler(subscribers ...UpdateSubscriber) {
	for _, subscriber := range subscribers {
		subscriber(c.updateCollection)
	}
}

func (c *pClient) On(method, channel string, handler UpdateHandler) {
	c.updateCollection.Add(method, channel, handler)
}

func (c *pClient) Start(errChan ...chan error) {
	var err error

	for {
		err = c.connect(errChan...)

		if err == nil || len(errChan) > 0 {
			break
		}
	}

	sendOne(errChan, err)
}

func (c *pClient) Stop() error {
	c.shouldStop = true
	return c.disconnect()
}

func (c *pClient) SendRequest(method, channel string, data any) (resp Response, err error) {
	var req *pRequest

	if !c.inRequestState {
		err = errors.New("Unable to make request. Please try again later.")
		return
	}

	if req, err = newRequest(method, channel); err != nil {
		return
	}

	if err = req.SetContent(data); err != nil {
		return
	}

	if _, err = writeRequest(c.server, req); err != nil {
		return
	}

	resp = c.getResponse(req)

	return
}

// =====================

func NewClient(opts ClientOptions) Client {
	responseCollection := newKCollection[*pRequest, *pResponse](func(k *pRequest) (key string) {
		return k.UUID + "::" + k.Method + "::" + k.Channel
	})

	value := &pClient{
		opts:               opts,
		responseCollection: responseCollection,
		updateCollection:   newMCollection[UpdateHandler](),
	}

	return value
}
