package gotimeout

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type myExampleObject struct {
	name        string
	id          int
	isModerator bool
}

// Do whatever you want in your handler
func handler(param interface{}) ExecutionResponse {
	// Verify that we receive the parameter we expect.
	// Here the name of MyExampleObject
	name, ok := param.(string)
	if !ok {
		return ExecutionResponse{
			Response: nil,
			Err:      errors.New("error while getting parameters"),
		}
	}

	return ExecutionResponse{
		Response: myExampleObject{
			name:        name,
			id:          999,
			isModerator: true,
		},
		Err: nil,
	}
}

// This handle will timeout on purpose for the sake of the example
func exampleOfHandlerThatTimedOut(param interface{}) ExecutionResponse {
	time.Sleep(5 * time.Second)

	// This code will never be reached if the timeout set is lower than 5 seconds
	return ExecutionResponse{
		Response: nil,
		Err:      nil,
	}
}

func TestMyFunction(t *testing.T) {
	executionResponse := Execute(handler, "Patrick", time.Second*1)
	obj, ok := executionResponse.Response.(myExampleObject)

	assert.True(t, ok)
	assert.Equal(t, "Patrick", obj.name)

	executionResponse = Execute(exampleOfHandlerThatTimedOut, "Patrick", time.Second*1)
	assert.Equal(t, errors.New("operation timed out"), executionResponse.Err)
}
