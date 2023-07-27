package gosrv

import "cyberpull.com/gotk/v2"

var json gotk.JSONEngine

func init() {
	json = gotk.NewJSON(validatorTagName)
}
