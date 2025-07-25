package observatory

import "github.com/v4fly/v4ray-core/v0/common/errors"

type errorCollector struct {
	errors *errors.Error
}

func (e *errorCollector) SubmitError(err error) {
	if e.errors == nil {
		e.errors = newError("underlying connection error").Base(err)
		return
	}
	e.errors = e.errors.Base(newError("underlying connection error").Base(err))
}

func newErrorCollector() *errorCollector {
	return &errorCollector{}
}

func (e *errorCollector) UnderlyingError() error {
	if e.errors == nil {
		return newError("failed to produce report")
	}
	return e.errors
}
