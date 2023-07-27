package gosrv

import "cyberpull.com/gotk/v2"

const validatorTagName string = "binding"

var validator gotk.Validator

func init() {
	validator = gotk.NewValidator(validatorTagName)
}
