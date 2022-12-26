package errorenum

import "zakariawahyu.com/go-gogen-mytodo/shared/model/apperror"

const (
	SomethingError             apperror.ErrorType = "ER0000 something error"
	NameMustNotEmpty           apperror.ErrorType = "ER0001 name must not empty"
	EmailMustNotEmpty          apperror.ErrorType = "ER0002 email must not empty"
	PasswordMustNotEmpty       apperror.ErrorType = "ER0003 password must not empty"
	UserAlreadyExist           apperror.ErrorType = "ER0004 user already exist"
	WrongEmailOrPassword       apperror.ErrorType = "ER0005 wrong email or password"
	UserIsNotActive            apperror.ErrorType = "ER0006 user is not active"
	UserAlreadyActivated       apperror.ErrorType = "ER0007 user already activated"
	UserActivatedTokenNotMatch apperror.ErrorType = "ER0008 user activated token not match"
)
