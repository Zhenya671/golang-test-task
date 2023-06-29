package messages

import "errors"

var (
	AppErrorWithMarshalling  = errors.New("something went wrong")
	AppErrorStatusBadRequest = errors.New("status bad request")

	AppErrorUserNotFound  = errors.New("such user not found")
	AppErrorSuchUserExist = errors.New("such user exist in the system yet")
	AppErrorUnauthorized  = errors.New("unauthorized")
	AppErrorCantPayOff    = errors.New("can't pay off")
)
