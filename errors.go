package imagespy

// ImageSpyError is the base error for all errors specific to this package.
type ImageSpyError struct {
	message string
}

func (e *ImageSpyError) Error() string {
	return e.message
}

// NotFoundError indicates that a resource could not be found.
type NotFoundError struct {
	ImageSpyError
	message string
}
