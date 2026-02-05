package todo

import "errors"

var ErrNotFound = errors.New("task not found")
var TaskAlreadyExists = errors.New("task already exists")