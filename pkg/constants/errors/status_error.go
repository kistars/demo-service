package errors

import "net/http"

//go:generate tools gen error StatusError
type StatusError int

func (StatusError) ServiceCode() int {
	return 999 * 1e3
}

const (
	// InternalServerError
	InternalServerError StatusError = http.StatusInternalServerError*1e6 + iota + 1
)

const (
	// @errTalk Unauthorized
	Unauthorized StatusError = http.StatusUnauthorized*1e6 + iota + 1
)

const (
	// @errTalk Conflict
	ConflictError StatusError = http.StatusConflict*1e6 + iota + 1
)

const (
	// @errTalk Forbidden
	ForbiddenError StatusError = http.StatusForbidden*1e6 + iota + 1
)

const (
	// @errTalk Not Acceptable
	NotAcceptable StatusError = http.StatusNotAcceptable*1e6 + iota + 1
)

const (
	// @errTalk NotFound
	NotFoundError StatusError = http.StatusNotFound*1e6 + iota + 1
)

const (
	// @errTalk BadRequest
	BadRequest StatusError = http.StatusBadRequest*1e6 + iota + 1
)
