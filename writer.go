package gosrv

import (
	"reflect"

	"cyberpull.com/gotk/errors"
)

func write[T any](prefix string, io NetIO, data T) (n int, err error) {
	dValue := reflect.ValueOf(data)
	vType := dValue.Type()

	if vType.Kind() != reflect.Pointer {
		err = errors.New("Data type must be a pointer")
		return
	}

	if dValue.IsNil() {
		err = errors.New("Data type must not be null")
		return
	}

	var b []byte

	if b, err = pJson.Encode(data); err != nil {
		return
	}

	output := prefix + separator + string(b)

	return io.WriteStringLine(output)
}

func writeError(io NetIO, data *Data) (n int, err error) {
	return write[*Data](errorPrefix, io, data)
}

func writeInfo(io NetIO, data *Info) (n int, err error) {
	return write[*Info](infoPrefix, io, data)
}

func writeRequest(io NetIO, data *pRequest) (n int, err error) {
	return write[*pRequest](requestPrefix, io, data)
}

func writeResponse(io NetIO, data *Response) (n int, err error) {
	return write[*Response](responsePrefix, io, data)
}

func writeUpdate(io NetIO, data *Update) (n int, err error) {
	return write[*Update](updatePrefix, io, data)
}

func writeState(io NetIO, name string, value bool) (n int, err error) {
	output := statePrefix + separator + name + "="

	if value {
		output += "YES"
	} else {
		output += "NO"
	}

	return io.WriteStringLine(output)
}

func mustWriteState(io NetIO, name string, value bool) (n int) {
	var err error

	if n, err = writeState(io, name, value); err != nil {
		panic(err)
	}

	return
}
