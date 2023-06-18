package gosrv

import (
	"reflect"

	"cyberpull.com/gotk/errors"
)

func get[T any](prefix string, io NetIO) (value T, err error) {
	vType := reflect.TypeOf(value)

	if vType.Kind() != reflect.Pointer {
		err = errors.New("Value type must be a pointer")
		return
	}

	var data string

	if data, err = io.ReadString('\n'); err != nil {
		return
	}

	return parse[T](data, prefix)
}

func getError(io NetIO) (value *Data, err error) {
	return get[*Data](errorPrefix, io)
}

func getInfo(io NetIO) (value *Info, err error) {
	return get[*Info](infoPrefix, io)
}

func getRequest(io NetIO) (value *pRequest, err error) {
	return get[*pRequest](requestPrefix, io)
}

func getResponse(io NetIO) (value *Response, err error) {
	return get[*Response](responsePrefix, io)
}

func getUpdate(io NetIO) (value *Update, err error) {
	return get[*Update](updatePrefix, io)
}
