package gosrv

import (
	"reflect"

	"cyberpull.com/gotk/v2/errors"
)

func SendRequest[T any](client Client, method, channel string, data any) (value T, err error) {
	if client == nil {
		err = errors.New("Invalid Client instance")
		return
	}

	var resp Response

	if resp, err = client.SendRequest(method, channel, data); err != nil {
		return
	}

	var tmpValue T

	vType := reflect.TypeOf(value)

	if vType.Kind() == reflect.Pointer {
		tmpValue = reflect.New(vType.Elem()).Interface().(T)
		err = resp.ParseContent(tmpValue)
	} else {
		err = resp.ParseContent(&tmpValue)
	}

	if err != nil {
		return
	}

	value = tmpValue

	return
}
