package handler

import "errors"

var (
	ErrorEmptyURL   = errors.New("Empty url param input")
	ErrorEmptyAlias = errors.New("Empty alias param input")
)
