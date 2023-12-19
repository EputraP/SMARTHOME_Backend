package errs

import "errors"

var (
	EmailAlreadyUsed          = errors.New("email already used")
	UsernameAlreadyUsed       = errors.New("username already used")
	PasswordDoesntMatch       = errors.New("password doesn't match")
	PasswordContainUsername   = errors.New("password must not contain username")
	PasswordSameAsBefore      = errors.New("Password cannot be same as before")
	UsernamePasswordIncorrect = errors.New("username or password incorrect")
)