package exception

import "encoding/json"

type ErrorCode string

const (
	SERVER_ERROR  ErrorCode = "S001"
	INVALID_ERROR ErrorCode = "E001"
)

type CustomError struct {
	StatusCode int       `json:"statusCode"`
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
}

func (e CustomError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func InvalidError() CustomError {
	return createError(400, INVALID_ERROR)
}

func NotFoundError(errorCode ErrorCode) CustomError {
	return createError(404, errorCode)
}

func InternalServerError() CustomError {
	return createError(500, SERVER_ERROR)
}

func createError(statucCode int, errorCode ErrorCode) CustomError {
	return CustomError{
		StatusCode: statucCode,
		Code:       errorCode,
		Message:    errorCode.getMessage(),
	}
}

func (e ErrorCode) getMessage() string {
	switch e {
	case SERVER_ERROR:
		return "server internal error"
	case INVALID_ERROR:
		return "invalid error"
	default:
		return "Unknown error - Please contact support"
	}
}
