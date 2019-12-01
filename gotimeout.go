// This package provide a function to execute anything you want with a timeout.
package gotimeout

import (
	"errors"
	"time"
)

// Return value of the Execute function.
type ExecutionResponse struct {
	Response interface{}
	Err      error
}

// This function execute the function given in parameter1 with the params given in parameter2.
// If the execution time is taking more time than the timeout given in parameter3 it returns an error.
func Execute(f func(interface{}) ExecutionResponse, params interface{}, timeout time.Duration) ExecutionResponse {

	c1 := make(chan ExecutionResponse, 1)
	defer close(c1)

	go func(func(interface{}) ExecutionResponse, interface{}) {
		c1 <- f(params)
	}(f, params)

	select {
	case res := <-c1:
		return res

	case <-time.After(timeout):
		return ExecutionResponse{
			Response: nil,
			Err:      errors.New("operation timed out"),
		}
	}
}
