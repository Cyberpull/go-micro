package gosrv

import "strings"

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
}

func (c *pClient) processUpdate(data string) {
	update, err := parseUpdateValue(data)

	if err != nil {
		return
	}
}

func (c *pClient) processState(data string) {
	name, value := parseStateValue(data)

	switch name {
	case requestState:
		c.inRequestState = value
	}
}

func (c *pClient) getResponse(req *pRequest) (resp *Response) {
	var ok bool

	for {
		resp, ok = c.responseCollection.Get(req)

		if ok {
			break
		}
	}

	return
}
