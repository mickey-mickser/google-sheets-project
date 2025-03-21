package responses

type ResponseError struct {
	Detail string `json:"detail"`
	Code   string `json:"code"`
}

var InternalErrorResponse = ResponseError{
	Detail: "Something bad happened",
	Code:   "000",
}

var ExternalIdIsUsedFailedResponse = ResponseError{
	Detail: "External id is used",
	Code:   "002",
}

var SuchNotExistResponse = ResponseError{
	Detail: "Such someone is not exist",
	Code:   "003",
}

var SuchIsNotExist = ResponseError{
	Detail: "Such someone is not exist",
	Code:   "004",
}

func NewBadRequestError(err error) ResponseError {
	return ResponseError{
		Detail: err.Error(),
		Code:   "001",
	}
}
