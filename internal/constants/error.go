package constants

import "errors"

// Metadata error
var (
	ErrMalformedMetadata = errors.New("[Meta] Malformed repository metadata")
)

// Config error
var (
	ErrMissingToken       = errors.New("[Config] Access token is empty")
	ErrNegativeFileChange = errors.New("[Config] Maximum file change must not be a negative number")
)
