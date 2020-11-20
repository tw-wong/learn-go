# learn-go
Some useful information related with Go.

## Overview
* [Basic data type](#basic-data-type)
* [Type conversion](#type-conversion)
* [Formatting](#formatting)
* [Array](#array)
* [Slice](#slice)
* [Map](#map)
* [New](#new)
* [Make](#make)
* [Type assertions](#type-assertions)
* [Pointer](#pointer)
* [Custom error](#custom-error)
* [Defer, Panic and Recover](#defer-panic-and-recover)
* [Struct](#struct)
* [Functions](#functions)
* [Interface](#interface)
* [Goroutines](#goroutines)

## Basic data type
| Type | Range | Note | 
|------|------|------|
| int  | 32 / 64 bit, depends on system. |  |
| uint  | 32 / 64 bit, depends on system. |  |
|int8 / byte   | -128 to 127 |  |
| int16        | -32,768 to 32,767 |  |
| int32 / rune | -2,147,483,648 to 2,147,483,647 |  |
| int64        | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807   |  |
| uint8        | 0 to 255 |  |
| uint16       | 0 to 65,535 |  |
| uint32       | 0 to 4,294,967,295 |  |
| uint64       | 0 to 18,446,744,073,709,551,615 |  |
| float32      | -3.4E+38 to +3.4E+38 | about 7 decimal digits. |
| float64      | -1.7E+308 to +1.7E+308 | about 16 decimal digits. |

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
```golang
// declares an array with int type and its length with 5.
var a [5]int
fmt.Println(a)
//[0 0 0 0 0]

// declares an array with string type and its length with 2.
var b [2]string
b[0] = "Hello"
b[1] = "World"
fmt.Println(b)
//[Hello World]

c := [3]int{}
fmt.Println(c)
//[0 0 0]

d := [3]int{1, 3, 5}
fmt.Println(d)
//[1 3 5]
```
* An array's length is part of its type, so arrays cannot be resized.

## Slice
```golang
    // Case 01:
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)
    //[3, 5, 7]
    
    // Case 02:
    names := [2]string{
        "John",
        "Paul",
    }
    fmt.Println(names)
    // [John Paul]
    
    // declares a slice variable
    a := names[0:2]
    fmt.Println(a)
    // [John Paul]
    
    names[1] = "Doe"
    fmt.Println(a)
    // [John Doe]

```
* Slice does not store any data, it is a reference.

```golang
package main

import "fmt"

func main() {
    // Case 01:
    s := make([]string, 3)
    fmt.Println("emp:", s) // emp: [  ]    
    
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s) // set: [a b c]
    fmt.Println("get:", s[2]) // get: c	        
    
    printSlice(s) // len=3 cap=3 [a b c]
    
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s) // apd: [a b c d e f]
    
    printSlice(s) // len=6 cap=6 [a b c d e f]
    
    // Case 02:
    c := make([]string, len(s))
    copy(c, s) // copy c from s
    fmt.Println("cpy:", c) // cpy: [a b c d e f]
    
    // this gets a slice of the elements s[2], s[3], and s[4].
    l := s[2:5]
    fmt.Println("sl1:", l) // sl1: [c d e]
    
    // this slices up to (but excluding) s[5].
    l = s[:5]
    fmt.Println("sl2:", l) // sl2: [a b c d e]
    
    // this slices up from (and including) s[2].
    l = s[2:]
    fmt.Println("sl3:", l) // sl3: [c d e f]    

}

func printSlice(s []string) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
* Create an empty slice with non-zero length, use the builtin `make`.
* Use builtin `append`, which returns a slice containing one or more new values.

## Map
```golang
package main

import "fmt"

type Person struct {
    name string
    age int
}

var m map[string]Person

func main() {
    m = make(map[string]Person)
    m["first_person"] = Person{
        name: "Alice", 
        age: 19, 
    }

    fmt.Println(m) //map[first_person:{Alice 19}]
    fmt.Println(m["first_person"].name) //Alice

}
```
* To initialize a map, use the built in `make` function.
* Maps are not safe for concurrent use.
* Iteration order is not specified and is not guaranteed to be the same from one iteration to the next.

## New
```golang
package main

import "fmt"

type Person struct {
    name string
    age int
}

func main() {
    p := new(Person)
    fmt.Println(p) //&{ 0}

    p.name = "Vincent"	
    fmt.Println(p) //&{Vincent 20}
}

```
* Syntax is `new(T)`.
* It returns a pointer to a newly allocated zero value of type T.
* Refs: https://golang.org/doc/effective_go.html#allocation_new

## Make
```golang
package main

import "fmt"

func main() {
    v := make([]int, 10)
    fmt.Println(v) //[0 0 0 0 0 0 0 0 0 0]
}
```
* Syntax is `make(T, args)`.
* It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not pointer).
* Refs: https://golang.org/doc/effective_go.html#allocation_make

## Type assertions
```golang
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println(s) //hello

    s, ok := i.(string)
    fmt.Println(s, ok) //hello true

    f, ok := i.(float64)
    fmt.Println(f, ok) //0 false
}
```

* Syntax `t, ok := i.(T)`. It will not trigger panic if `i` does not hold a `T` (type).
* Syntax `t := i.(T)`. It will trigger panic if `i` does not hold a `T` (type).

## Pointer
```golang
package main

import "fmt"

func zero(xPtr *int) {
    *xPtr = 0
}

func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
}
```

```golang
package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i         // point to i.
    fmt.Println(p)  // read memory address of i. Output: 0xc00002c008
    fmt.Println(*p) // read i through the pointer. Output: 42
    *p = 21         // set i through the pointer. i changed from 42 to 21 now.
    fmt.Println(i)  // see the new value of i. Output: 21

    p = &j         // point to j.
    *p = *p / 37   // divide j through the pointer. 2701 / 37 = 73.
    fmt.Println(j) // see the new value of j. Output: 73.
}
```

```golang
package main

import "fmt"

func main() {
    var intVar int
    var pointerVar *int
    var pointerToPointerVar **int

    intVar = 100
    pointerVar = &intVar
    pointerToPointerVar = &pointerVar

    fmt.Println("Group 01")
    fmt.Println("intVar: ", intVar)                                 //100
    fmt.Println("pointerVar: ", pointerVar)                         //0xc00002c008
    fmt.Println("pointerToPointerVar: ", pointerToPointerVar)       //0xc00000e028

    fmt.Println("Group 02")
    fmt.Println("&intVar: ", &intVar)                               //0xc00002c008
    fmt.Println("&pointerVar: ", &pointerVar)                       //0xc00000e028
    fmt.Println("&pointerToPointerVar: ", &pointerToPointerVar)     //0xc00000e030

    fmt.Println("Group 03")
    fmt.Println("*pointerVar: ", *pointerVar)                       //100
    fmt.Println("*pointerToPointerVar: ", *pointerToPointerVar)     //0xc00002c008
    fmt.Println("**pointerToPointerVar: ", **pointerToPointerVar)   //100
}
```

* `*` operator uses to "dereference" pointer variable. Dereference a pointer gives us to access to the value of the pointer.
* Example: `*xPtr = 0`, it means "store the `int` 0 in the memory location `*xPtr` refers to.
* `&` operator uses to find the memory address of variable.

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
        Err: errors.New("Service unavailable"),
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
```golang
package animals

// Dog represents information about dogs.
type Dog struct {
    Name         string
    BarkStrength int
    age          int
}
```

```golang
package main

import (
    "fmt"
    "test/animals"
)

func main() {
    // Create an object of type Dog from the animals package.
    // This will NOT compile.
    dog := animals.Dog{
        Name: "Bingo",
        BarkStrength: 10,
        age: 5,
    }

    fmt.Printf("Counter: %#v\n", dog) //Output: unknown animal.Dog field ‘age’ in struct literal
}
```
* If a field or method name starts with a capital letter, the member is exported and is accessible outside of the package. 
* If a field or method starts with a lowercase letter, the member is unexported and does not have accessibility outside of the package.

```golang
package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func main() {
    p := Person{
        Name: "John",
        Age: 19,
    }
	
    fmt.Println(p) //Output: {John 19}
    
    p.Name = "John Doe"
    p.Age = 20
    fmt.Println(p) //Output: {John Doe 20}
    
    ptr := &Person {
        Name: "Pointer", 
        Age: 25, 
    }
    ptr.Age = 28
    fmt.Println(ptr) //Output: &{Pointer 28}
    fmt.Println(ptr.Age) //Output: 28
}
```
* To use pointer to a `Struct`, use `&` operator. Ex: `&Person`.
* It allows to access the fields without any dereferencing it explicitly. Ex: `ptr.Age`.
* Golang allows the programmers to access the fields of a structure using the pointers without any dereferencing explicitly. 

## Functions 
```golang
package main

import "fmt"

func plus(a int, b int) int {
    return a + b
}

func main() {
    res := plus(1, 2)
    fmt.Println(res) //Output: 3
}
```
* Call a function with `name(args)`.

```golang
package main

import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {
    a, b := vals()
    fmt.Println(a) //Output: 3
    fmt.Println(b) //Output: 7
}
```
* Multiple return values.

```golang
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    
    return total
}

func main() {
    amount_a := sum(10, 15, 20, 30)
    fmt.Println(amount_a) //Output: 75
    
    args := []int{1,2,3,4,5}
    amount_b := sum(args...)
    fmt.Println(amount_b) //Output: 15
}
```
* `Variadic functions` can be called with any number of trailing arguments. 

```golang
package main

import "fmt"

func update(i *int) {
    *i++    
}

func main() {
    a := 1
    fmt.Println(a) //Output: 1
    
    update(&a)
    fmt.Println(a) //Output: 2
}
```
* Passing a pointer as function's argument (`&a`).

```golang
package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()
    
    fmt.Println(nextInt()) //Output: 1
    fmt.Println(nextInt()) //Output: 2
    fmt.Println(nextInt()) //Output: 3
    
    newInts := intSeq()
    fmt.Println(newInts()) //Output: 1
}
```
* `Anonymous functions` are useful when you want to define a function inline without having to name it.
* Function `intSeq` returns another function, which defined anonymously in the body of `intSeq`. The returned function closes over the variable `i` to form a closure.

## Interface
```golang
package main

import "fmt"

// create an interface for staff
type IStaff interface {
    CalculateSalary() int
}

// create full time staff struct
type FullTimeStaff struct {
    empId int
    basicPay int
    allowance int
}

// create part time staff struct
type PartTimeStaff struct {
    empId int
    hourRate int
    workingHour int
    workingDay int
}

// "implement" interface to full time staff struct
func (f FullTimeStaff) CalculateSalary() int {
    return f.basicPay + f.allowance
}

// "implement" interface to part time staff struct
func (p PartTimeStaff) CalculateSalary() int {
    return p.hourRate * p.workingHour * p.workingDay
}

// calculate total expense
func totalExpense(s []IStaff) int {
    expense := 0
    for _, v := range s {
        expense += v.CalculateSalary()
    }
    
    return expense
}

func main() {
    staff01 := FullTimeStaff{
        empId: 1, 
        basicPay: 2000, 
        allowance: 50,        
    }

    staff02 := PartTimeStaff{
        empId: 2, 
        hourRate: 8, 
        workingHour: 4, 
        workingDay: 20, 
    }

    employees := []IStaff{staff01, staff02}
    total := totalExpense(employees)
    fmt.Printf("Total Expense Per Month is: $%d", total) //Output: Total Expense Per Month is: $2690
}
```
* Pass in param of `totalExpense` is interface `[]IStaff`, easy to extend.

```golang
package main

import "fmt"

type IStaff interface {
    CalculateSalary() float32
    IssueBonus()
}

// create full time staff struct
type FullTimeStaff struct {
    empId int
    basicPay float32
    allowance float32
    bonus float32
}

// create part time staff struct
type PartTimeStaff struct {
    empId int
    hourRate float32
    workingHour int
    workingDay int
    bonus float32
}

// "implement" interface to full time staff struct with pointer receiver
func (f *FullTimeStaff) IssueBonus() {
    bonus := f.basicPay * 0.1
    f.bonus = bonus
}

// "implement" interface to full time staff struct with value receiver
func (f FullTimeStaff) CalculateSalary() float32 {
    return f.basicPay + f.allowance + f.bonus
}

// "implement" interface to full time staff struct with pointer receiver
func (p *PartTimeStaff) IssueBonus() {
     monthlyHour := p.workingHour * p.workingDay
     if monthlyHour >= 50 {
          p.bonus = 10
     } else {
         p.bonus = 5
     }
}

// "implement" interface to full time staff struct with value receiver
func (p PartTimeStaff) CalculateSalary() float32 {
     return (p.hourRate * float32(p.workingHour) * float32(p.workingDay)) + p.bonus
}

// calculate total expense
func totalExpense(s []IStaff) float32 {
    expense := float32(0)
    for _, v := range s {
        expense += v.CalculateSalary()
    }
    
    return expense
}

func main() {
    staff01 := &FullTimeStaff{
        empId: 1, 
        basicPay: 185, 
        allowance: 50,        
    }
    staff01.IssueBonus()
    staff01Salary := staff01.CalculateSalary()
    fmt.Printf("staff01 salary is: $%.2f\n", staff01Salary) //Output: staff01 salary is: $253.50

    staff02 := &PartTimeStaff{
         empId: 2, 
         hourRate: 8, 
         workingHour: 4, 
         workingDay: 20, 
    }
    staff02.IssueBonus()
    staff02Salary := staff02.CalculateSalary()
    fmt.Printf("staff02 salary is: $%.2f\n", staff02Salary) //Output: staff02 salary is: $650.00

    employees := []IStaff{staff01, staff02}
    total := totalExpense(employees)
    fmt.Printf("Total Expense Per Month is: $%.2f", total) //Output: Total Expense Per Month is: $903.50
}
```
* Interface with pointer receiver.  Ex: `func (f *FullTimeStaff) IssueBonus()`.
* When the interface is using a value receiver, it works for both variables type (value or pointer). That's why `staff01.CalculateSalary()` can be executed without any issue when the variable type is pointer (`staff01 := &FullTimeStaff`).
* When interface using pointer receiver, it only works for pointer variables type.

```golang
package main

import "fmt"

// Define an interface as type
type Dog struct {
    Age interface{}
}

//Define a struct
type Cat struct {
}

// pass an empty interface type as a function parameter:
func Guess(t interface{}) {
    switch t.(type) {
	case Dog:
	    fmt.Println("Dog type")
	case Cat:
	    fmt.Println("Cat type")
	default:
	    fmt.Println("Unknown type")
    }
}

func main() {
    dog := Dog{}
    dog.Age = "3"
    fmt.Printf("%#v %T\n", dog.Age, dog.Age) //Output: "3" string

    dog.Age = 3
    fmt.Printf("%#v %T\n", dog.Age, dog.Age) //Output: 3 int

    cat := Cat{}

    Guess(dog) //Output: Dog type
    Guess(cat) //Output: Cat type
    Guess(1)   //Output: Unknown type
}
```
* `interface{}` is empty interface.
* When struct with type `interface{}`, we can assign any type for it. Ex: `dog.Age = '3' (type is string)` or `dog.Age = 3 (type is int)`.

## Goroutines
```golang
package main

import "fmt"

func hello(c chan bool) {
    fmt.Println("Hello world goroutine")
    c <- true //send data to channel
    
}
func main() {
    c := make(chan bool)
    go hello(c)
    <- c //read / receive data from channel, but not use or store the data in any variable is legal.
    fmt.Println("main function")
    
    // Output:
    // Hello world goroutine
    // main function
}
```
* Goroutines allow functions or methods that run concurrently with other functions or methods.
* Use channel to communicate with Goroutines.
* When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel.
* When data is read from a channel, the read is blocked until some Goroutine writes data to that channel.
* Refs: https://golangbot.com/goroutines/

```golang
package main

import "fmt"

func processA(w int, h int, ca chan int) {
    sum := 0
    sum = w * h
    ca <- sum //send process result
}

func processB(a int, b int, cb chan int) {
    sum := 0
    sum = a + b
    cb <- sum
}

func main() {
    ca := make(chan int) //create channel for process a and will return process result.
    cb := make(chan int) //create channel for process b and will return process result.
    go processA(5, 6, ca)
    go processB(5, 5, cb)
    resultA, resultB := <- ca, <- cb
    total := resultA + resultB
    fmt.Println("Total is:", total) //Output: Total is: 40
}
```
* Example of multiple Goroutines run separately.

```golang
package main

import "fmt"

func processA(w int, h int, ch chan <- int) {
    sum := 0
    sum = w * h
    ch <- sum //send process result
}

func main() {
    ch := make(chan int) //create channel for process a and will return process result.
    go processA(5, 6, ch)
    resultA := <- ch //read data from channel.
    fmt.Println("Result A is:", resultA) //Output: Result A is: 30
}
```
* `ch` is defined as bidirectional channel.
* The `processA` function converts this channel (`ch chan <- int`) to a send only channel (unidirectional channel). So the channel is send only in this function.
