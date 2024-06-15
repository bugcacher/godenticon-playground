package api

import "errors"

// Http errors
var (
	ErrMethodNotAllowed = errors.New("only GET is allowed")
	ErrInterServerError = errors.New("some error occurred")
)

// Validation errors
var (
	ErrRequiredValue       = errors.New("value is required")
	ErrInvalidPixelPattern = errors.New("invalid pixel_pattern value")
	ErrInvalidAlgo         = errors.New("invalid algorithm value")
	ErrInvalidDimension    = errors.New("invalid dimension value")
	ErrInvalidDarkMode     = errors.New("invalid dark mode value")
)
