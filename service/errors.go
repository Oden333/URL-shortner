package service

import "errors"

var (
	ErrorURLNotFound = errors.New("URL not found")
	ErrorURLExists   = errors.New("URL already exists")
	ErrorEmptyAlias  = errors.New("Empty alias param")
)
