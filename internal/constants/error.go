package constants

import "errors"

// Metadata error
var (
	ErrMalformedMetadata = errors.New("error when parsing repository metadata")
)
