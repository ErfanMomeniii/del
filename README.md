del is a simple package that handles event listeners and an event dispatcher. it runs listeners' functions in parallel.
# Documentation
## Install
To install this package, run the following command in your project directory.
```bash
go get github.com/ErfanMomeniii/del
```   
Next, include it in your application:
```bash
import "github.com/ErfanMomeniii/del"
``` 
## Quick Start

The following example demonstrates how to create an event and two listeners and attach those listeners to the event and finally dispatch the event.
```go
package main

import (
	"fmt"
	"github.com/ErfanMomeniii/del"
)

func main() {
  // create event
	e := del.NewEvent("first")
  
  // create listener
	l1 := del.NewListener(func() {
		fmt.Println("hello")
	})

  //create listener
	l2 := del.NewListener(func() {
		fmt.Println("goodbye")
	})

  //attach two listener(l1 and l2) to the event
	e.AddListener(l1, l2)

  //dispatch event
	e.Dispatch()
}

//Output:
  //hello
  //goodbye

```
This package run listener's function concurrently. please be careful that it doesn't support priority between listener functions.



## Authors

- [@ErfanMomeniii](https://github.com/ErfanMomeniii/)




