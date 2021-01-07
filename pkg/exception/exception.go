package exception

import "errors"

var EntityNotFound = errors.New("not found")

var Conflict = errors.New("conflict")
