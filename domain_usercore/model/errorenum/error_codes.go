package errorenum

import "zakariawahyu.com/go-gogen-mytodo/shared/model/apperror"

const (
	SomethingError       apperror.ErrorType = "ER0000 something error"
	NameMustNotEmpty     apperror.ErrorType = "ER0001 name must not empty"
	EmailMustNotEmpty    apperror.ErrorType = "ER0002 email must not empty"
	PasswordMustNotEmpty apperror.ErrorType = "ER0003 password must not empty"
)
