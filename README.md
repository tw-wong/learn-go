# learn-go
Some useful information related with Go.
## Basic data type
| Type | Range | Default / Note | 
|------|------|------|------|
| int  | 32 / 64 bit, depends on system. | 1 |
| uint  | 32 / 64 bit, depends on system. | 0 |
|int8 / byte   | -128 to 127 | 0 |
| int16        | -32,768 to 32,767 | 0 |
| int32 / rune | -2,147,483,648 to 2,147,483,647 | 0 |
| int64        | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807   | 0 |
| uint8        | 0 to 255 | 0 |
| uint16       | 0 to 65,535 | 0 |
| uint32       | 0 to 4,294,967,295 | 0 |
| uint64       | 0 to 18,446,744,073,709,551,615 | 0 |
| float32      | -3.4E+38 to +3.4E+38 | 0 / about 7 decimal digits |
| float64      | -1.7E+308 to +1.7E+308 | 0 / about 16 decimal digits |

Refs: https://golang.org/pkg/builtin/

## Type conversion
```golang
package main

import (
	"fmt"
)

func main() {
	var f float64 = 3.9
	c := uint(f)
	fmt.Println(f) //3.9
	fmt.Println(c) //3
}

// Output:
3.9
3
```

## Formatting
```golang
name := "Leslie"
fmt.Printf("My name is %v", name)
// My name is Leslie

age := 34
fmt.Printf("I am %d years old", age)
// I am 34 years old

fmt.Printf("%v is of type %T", name, name)
// Leslie is of type string
```

* `%v` represents the named value in its default format.
* `%d` expects the named value to be an integer type.
* `%f` expects the named value to be a float type.
* `%T` represents the type for the named value.

## Array

## Slice

## Map

## New

## ()

## Pointer

## Custom error
```golang
package main

import (
    "errors"
    "fmt"
    "os"
)

type RequestError struct {
    StatusCode int

    Err error
}

func (r *RequestError) Error() string {
    return fmt.Sprintf("status: %d, error: %v", r.StatusCode, r.Err)
}

func doRequest() error {
    return &RequestError{
        StatusCode: 503,
        Err:        errors.New("Service unavailable"),
    }
}

func main() {
    err := doRequest()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("success!")
}

// Output:
status: 503, error: Service unavailable
```


## Defer, Panic and Recover
```golang
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f. Note: This is return because panic has recovered in f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g. Note: this will not return since panic has occurred.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}

// Output:
 
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f. Note: This is return because panic has recovered in f.
```

## Struct

## Interface
