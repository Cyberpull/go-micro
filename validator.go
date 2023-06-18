package gosrv

import "cyberpull.com/gotk/validator"

const validatorTagName string = "binding"

var pValidator *validator.ValidatorInstance

func init() {
	pValidator = validator.New(validatorTagName)
}
