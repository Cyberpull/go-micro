package gosrv

import "cyberpull.com/gotk/objects"

var pJson objects.JSONEngine

func init() {
	pJson = objects.NewJSON(validatorTagName)
}
