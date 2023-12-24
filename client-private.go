package gosrv

import (
	"net"
	"strings"

	"cyberpull.com/gotk/v2/log"
)

func (c *pClient) connect(errChan ...chan error) (err error) {
	defer func() {
		if c.shouldStop {
			err = nil
		}

		c.reset()
	}()

	var conn net.Conn

	if conn, err = connect(c.opts); err != nil {
		sendOne(errChan, err)
		return
	}

	c.server = newServerConn(conn)

	if err = c.prepare(); err != nil {
		sendOne(errChan, err)
		return
	}

	log.Successfln("Connected to %s on %s", c.server.info.Name, address(c.opts))

	c.inRequestState = true

	sendOne(errChan, nil)

	for {
		var input string

		if input, err = c.server.ReadString('\n'); err != nil {
			// log.Errorln(err)
			break
		}

		go c.processInput(input)
	}

	return
}

func (c *pClient) prepare() (err error) {
	var input string

	// Get welcome message
	if input, err = c.server.ReadString('\n'); err != nil {
		return
	}

	log.Println(input)

	// Send client info
	if _, err = c.server.WriteInfo(c.opts.Info); err != nil {
		return
	}

	// Get server info
	c.server.info, err = c.server.ReadInfo()

	return
}

func (c *pClient) disconnect() (err error) {
	if c.server != nil {
		err = c.server.Close()
	}

	return
}

func (c *pClient) reset() {
	c.responseCollection.Clear()
	c.updateCollection.Clear()

	c.shouldStop = false
	c.inRequestState = false
	c.server = nil
}

func (c *pClient) processInput(input string) {
	chunks := strings.SplitN(input, separator, 2)

	if len(chunks) != 2 {
		return
	}

	value := chunks[1]

	switch chunks[0] {
	case responsePrefix:
		c.processResponse(value)

	case updatePrefix:
		c.processUpdate(value)

	case statePrefix:
		c.processState(value)
	}
}

func (c *pClient) processResponse(data string) {
	resp, err := parseResponseValue(data)

	if err != nil {
		return
	}

	c.responseCollection.AddOne(resp.Request, resp)
}

func (c *pClient) processUpdate(data string) {
	update, err := parseUpdateValue(data)

	if err != nil {
		return
	}

	handlers := c.updateCollection.GetAll(
		update.Method,
		update.Channel,
	)

	for _, handler := range handlers {
		handler(update)
	}
}

func (c *pClient) processState(data string) {
	name, value := parseStateValue(data)

	switch name {
	case requestState:
		c.inRequestState = value
	}
}

func (c *pClient) getResponse(req *pRequest) (resp *pResponse) {
	var ok bool

	for {
		resp, ok = c.responseCollection.Get(req)

		if ok {
			c.responseCollection.ClearKey(req)
			break
		}
	}

	return
}
