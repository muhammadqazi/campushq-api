// This file is added for purpose of having an example of different path in tests.
package tracerr_test

import "github.com/campushq-official/campushq-api/src/internal/common/tracerr"

func addFrameA(message string) error {
	return addFrameB(message)
}

func addFrameB(message string) error {
	return addFrameC(message)
}

func addFrameC(message string) error {
	return tracerr.New(message)
}
