package errs

type DtoError struct {
	Status  int
	Code    int
	Message interface{}
}

type DataError struct {
	DtoError
}

type TokenError struct {
	DtoError
}

func NewDataError(status, code int, msg interface{}) *DataError {
	return &DataError{
		DtoError{
			Status:  status,
			Code:    code,
			Message: msg,
		},
	}
}

func NewTokenError(status, code int, msg interface{}) *TokenError {
	return &TokenError{
		DtoError{
			Status:  status,
			Code:    code,
			Message: msg,
		},
	}
}
