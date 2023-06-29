package messages

import "errors"

var (
	AppErrorWithMarshalling  = errors.New("something went wrong")
	AppErrorStatusBadRequest = errors.New("status bad request")

	AppErrorUserNotFound  = errors.New("such user not found")
	AppErrorSuchUserExist = errors.New("such username exist in the system yet")
)
