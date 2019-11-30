# gotimeout
Go library to execute anything you want with a timeout.

> Example

```golang
package main

import (
	"fmt"
	"github.com/MaximeAubanel/gotimeout"
	"time"
)

func handle(params interface{}) gotimeout.ExecutionResponse {
	return gotimeout.ExecutionResponse{
		Response: params,
		Err:      nil,
	}
}

func handleThatWillTimeOut(params interface{}) gotimeout.ExecutionResponse {
	time.Sleep(10 * time.Second)
	return gotimeout.ExecutionResponse{
		Response: params,
		Err:      nil,
	}
}

func main() {
	resp := gotimeout.Execute(handle, "Hello there", time.Second)
	fmt.Println(resp.Response.(string))

	resp = gotimeout.Execute(handleThatWillTimeOut, "Hello there", time.Second)
	fmt.Println(resp.Err.Error())

}
```