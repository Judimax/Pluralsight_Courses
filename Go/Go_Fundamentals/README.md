# Course Overview

## Course Overview

# Getting Up and Running

## Introduction

## Prerequisites and Expected Outcomes

## Getting Up and Running

## Demo: Installing Go

## Setting Up a Code Editor

## Demo: Installing and Configuring VS Code

## Demo: Your First Go Program

## Course Plan

## Summary

# Variables and Simple Data Types

## Introduction

## Simple Data Types
* strings,numbers,booleans,error
* only contain a single value

## The String Type
* 1 or more utf-8 chars
![1686311874648](image/README/1686311874648.png)

![1686311880116](image/README/1686311880116.png)

## Numeric Types
* ![1686311939358](image/README/1686311939358.png)

## Boolean Types'
* true or false

## Error Types
* an error is an error
* it has an ability to report an error
```go
type error interface {
  Error() string
}
```

## Finding Documentation for Built-in Types

## Declaring Variables
```go
var myName string
var myName string = "Mike"

var myNasme = "Mike"
myName := "Mike"  //short declaration syntax
```
## Type Conversions
* go doesnt support implicit conversion
```go
var i int =32
var f float32
f = i
f = float32(i)
```

## Demo: Using Simple Types and Type Conversions
* local varlues must have a usage
Go_Fundamentals\fundamentals-go\03\demos\01_variables\main.go

## Common Arithmetic Operators
```go
a, b := 10, 5

integer division is  /
but floating point is also /
```


## Common Comparison Operators
* comparison is only supportted for number
[languae spec](go.dev/ref/spec#Arithmetic_operators)


## Demo: Using Arithmetic and Comparison Operators

## Constants, Constant Expressions, and Iota
* to make a constant
```go
const a = 42
const b string ="hello, world"
const (
  d = true,
  e = 3.14
)
const (
  d = "foo",
  e  // foo
)
const e = someFunction( ) //mo cant be allocated at run time
const (
  b = iota // 0
  c // 1
  d =3 * iota // 6
)
const (
  e = iota // 0 resets at ver block
)
```
* go is going to treat a untyped constant as a literal meaning it can be applied to strings, int
Go_Fundamentals\fundamentals-go\03\demos\03_constants\main.go
* iota in action
```go
type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```
## Demo: Constants, Constant Expressions, and Iota
Go_Fundamentals\fundamentals-go\03\demos\03_constants\main.go


## Pointers
```go
a := 42
b := a
a = 27 // b = 42

a :=42
b := &a //& address operator
a = 27 // *b 27  not holding value holding address to another variable

c = new(int) // points to anomyous memory

// use copies but you code gets buggy during concurrent programs
```

## Demo: Creating and Using Pointers
Go_Fundamentals\fundamentals-go\03\demos\04_pointers\main.go

## Summary

# Creating and Debugging Programs

## Introduction

## Concept: Command-line Interfaces

## CLI Tools
os
  * stdin, stdout, stderr

fmt
  Scan*,
  Print*,
  bufio //gather groups of chars in to meaningful units


## Demo: Building a CLI Application
* ![](Go_Fundamentals\fundamentals-go\04\demos\01_cli_program\main.go)

## Concept: Web Services
* go web services
![1686314351390](image/README/1686314351390.png)
front cotroller is like a proxy, back controllers is the server

## Demo: Building a Web Service
Go_Fundamentals\fundamentals-go\04\demos\02_web_service\main.go

## Demo: Debugging a Program
Go_Fundamentals\fundamentals-go\04\demos\03_debugging\main.go


## Summary

# Aggregate Data Types
*

## Introduction

## Concept: Array Types


## Creating and Using Arrays
* when we assign arrays we get a copy
```go
arr := [3]string("foo","bar","baz")
```
* to compare array must have same size and type
* checks each index

## Demo: Arrays
Go_Fundamentals\fundamentals-go\05\demos\01_arrays\main.go
* cant grow or shrink

## Concept: Slice Types
* an array that can grow or shrink
* first index of a slice point to another value in an array
* slice does not hold data its poiting
  * we will change the underlying array


## Creating and Using Slices
* make a slice
```go
var s []int //slinces of int
fmt.Println(s) // [] (nil)
s = []int{1, 2 ,3}

slices.Delete(s,1,3)
```

## Demo: Slices
Go_Fundamentals\fundamentals-go\05\demos\02_slices\main.go
* to add dependecies
```go
go get golang.org/x/exp/slices
```
* makes go.sum a file that the dep we retreived is the dep we through


## Concept: Map Types
* an an object, but you  can customize the key type

## Creating and Using Maps
* delete a key for a map
```go
	delete(m, "tea")
```
* if you delete you get the default value for the key type
```go
v,ok  := map[k] //whether 0 comes from the map or map
```
* maps are not comparable
## Demo: Maps
Go_Fundamentals\fundamentals-go\05\demos\03_maps\main.go
* be careful when treating uninitalized as initalized
* this is mulltiple assignment,
* we need line separation with commans here
* map has no order, go is usiung internal data structure we dont have access to
```go
	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}
```
## Concept: Struct Types
* fixed sized, but values can be differnt types
Go_Fundamentals\fundamentals-go\05\demos\04_structs\main.go

## Creating and Using Structs
Go_Fundamentals\fundamentals-go\05\demos\04_structs\main.go

* create a struct
```go
	type menuItem struct {
		name   string
		prices map[string]float64
	}
```

* init a struct
```go
s := myStruct{
  name: "Aruthur",
  id: 42
}

* structs are comparable, and they copy by value
```
## Demo: Structs
Go_Fundamentals\fundamentals-go\05\demos\04_structs\main.go
## Summary

# Looping

## Introduction
* infinite look,
* loop till condiiton
* conuter based loops
* looping of sequence

## Concept: Looping

## Basic Loops
![1686322134136](image/README/1686322134136.png)

## Demo: Looping
Go_Fundamentals\fundamentals-go\06\demos\01_basic_loops\main.go

## Looping through Collections
* use range to tell go you are looking through a collection
![1686322652708](image/README/1686322652708.png)
* look for format verbs for the fmt package


## Demo: Looping through Collections
Go_Fundamentals\fundamentals-go\06\demos\02_looping_over_collections\02_after\main.go

## Summary

# Branching

## Introduction

## If Statements
* if stmt
```go
	if i := 5; i < 5 {
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is greater than or equal to 10")
	}
```


## Demo: If Statements
Go_Fundamentals\fundamentals-go\07\demos\01_if_statements\main.go


## Concept: Switch Statements
* you can have something like this
```go
case "2","3":

case 2+3,2*i+3:
```

## Switch Statements


## Logical Switches
* can create manually, but you wont really see this, just leave the value true out
```go
switch i :=8;true
```

## Demo: Switch Statements
Go_Fundamentals\fundamentals-go\07\demos\02_switches\02_after\main.go
* to make an empty map
```go
prices: make(map[string]float64)
```
*  label, to label your statement, in case for loop for the break to break out of the code
```go
loop:
```

## Concept: Deferred Functions
* go will defer till main function exits

## Deferred Functions
* defer function calls in a stack

## Demo: Deferred Functions
Go_Fundamentals\fundamentals-go\07\demos\03_deferred_functions\main.go
* to work with a sql database connection
* you want to release resouces in the way that you acquire them
```go
db, _ := sql.Open("driverName", "connection string")
defer db.Close()
```
## Concept: Panic and Recover
*

## Panic and Recover
* just like try and catch
* as soon a function panics it returns control to the calling fn
![1686323891158](image/README/1686323891158.png)
* the reconver function needs to be in an IIFE
```go
	defer func() {

	}()
```

## Demo: Panic and Recover
Go_Fundamentals\fundamentals-go\07\demos\04_panic_and_recover\02_after\main.go

## Goto Statements
* goto statement can leave a block
* can jump to a containing block
* cannot jump after variable declaration
* cannot jump to  another block

## Summary

# Organizing Programs

## Introduction


## Function Signatures
* to make a function
```go
func fuctionName (parameters)(return values){
  function body
}
```

## Function Parameters
* if you have 2+ params with same type
```go
(name1,name2 string)
```
* variadic parameters
  * treate as slice
  * cant have multiple it has to be the last one
```go
greet(name, ...string)
```



## Passing Values and Pointers as Parameters
* name pass to a function name will not change
* pointer passed to a function pointer will chnage
* use pointers to share memory
```go
func main(){
  name, otherName := "Name","Other"
  fmt.Println(name)
  fmt.Println(otherName)
  myFunc(name, &otherName)
   fmt.Println(name)      // "Name"
  fmt.Println(otherName) //"Other new Name"
}

func myFunc(name string, otherName *string){
  name = "New Name"
  *otherName = "Other new name"
}
```
## Returning Data from Functions
![1686324745535](image/README/1686324745535.png)

* naked returns can be confusing
![1686324768972](image/README/1686324768972.png)

## Demo: Functions
* Go_Fundamentals\fundamentals-go\08\demos\01_functions\02_after\main.go

## Concept: Packages
![1686324995473](image/README/1686324995473.png)

## Package Members and Scoping Rules
variable,constanct ,function are package level members
![1686325098954](image/README/1686325098954.png)


## Demo: Packages
Go_Fundamentals\fundamentals-go\08\demos\02_packages\02_after\main.go
* no private scope in go, the lowest scope is the package level, so if the package is imported everything can see one another
* all capital members are public, package level are lowercase

## Documenting Code
* documentation if for users
* first sentence, why the function is there in the first place


## Examples of Documentation in Standard Library

## Summary

# Object Orientation and Polymorphism

## Introduction

## Defining Methods
* methods dont need to be near the variable
* to create a method
  * will execute in the context of myInt
```go
type myInt int
// i myInt is a method receiver
func (i myInt) isEven() bool{
  function mody
}
// to call
var mi myInt
mi.isEven()
```


## Method Receivers
* very good to use pointers in update methods
```go
type user struct{
  id int
  username string
}

func (u user ) String() string{
  return fmt.Sprintf("%v (%v)\n", u.username,u.id)
}

func( u *user) UpdateName(n name){
  u.username = name
}
```

## Demo: Methods

## Concept: Interfaces
* ![1686326258310](image/README/1686326258310.png)
* we just care if we can read from a file or TCP connection, not how it workds


## Defining and Implementing Interfaces
* 	Print() is the interface
```go
type printer interface {
	Print() string
}
```

* we dont get File AND TCP, features
![1686326362644](image/README/1686326362644.png)

## Type Assertions
* this is how we get back  from assertion to conecrete type,
* use ok syntax,  if it fails f2 is a noop
* you can use type switches if you are dealing with several types
![1686326454677](image/README/1686326454677.png)
![1686326524653](image/README/1686326524653.png)
* Type switch
```go
	switch v := p.(type) {
	case user:
		fmt.Println("Found a user!", v)
	case menuItem:
		fmt.Println("Found a menuItem!", v)
	default:
		fmt.Println("I'm not sure what this is...")
	}
```
## Demo: Interfaces
Go_Fundamentals\fundamentals-go\09\demos\02_interfaces\main.go

## Concept: Generic Programming
* we may want our types to work polymorhipcs to work for a specific period of time
## Demo: Creating Generic Functions
* when the varible goes in it turns into  the interacem, when it leaves it turns back into the concrete
## Demo: Generics with the Comparable Constraint
Go_Fundamentals\fundamentals-go\09\demos\03_basic_generics\main.go
* V is the generic
to define the generic
```go
[V any]
```
```go
func clone[V any](s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}

	return result
}
```
## Demo: Creating Custom Type Constraints
Go_Fundamentals\fundamentals-go\09\demos\04_generic_constraints\main.go
* for maps keys needs to be comporable we can use geneic to clone
```go
[K comparable, V any]
```
Go_Fundamentals\fundamentals-go\09\demos\05_type_interfaces\main.go
```go
type Addable interface {
	int | float64 | string
}
func add[V Addable](s []V) V {
...
}
```
## Review: Generic Programming
* comparable is anything that can be compared with '=='
* constrants
golong.org/x/exp/constraints,slices,maps


## Summary

# Error Management


## Introduction

## Errors in Go
* errors are values not exceptions
* handle errors immediately
## Concept: Error Handling
* builtin,
* simplest way to create an err
```go
erro := errors.New("this is an error")
fmt.Println(errpr)

err2 := fmt.Errorf("this error the first one : %w",err)
```
* we wrap our errors because, we have business logic that called the low level logic

fmt.Errorf

## Demo: Creating Error Objects
Go_Fundamentals\fundamentals-go\10\demos\02_comma_error_pattern\02_after\main.go


## Demo: Error Handling

## Concept: Errors vs. Panics
* errors is easy to find,
* erros used frequently

## Demo: Converting Panics to Errors
* Go_Fundamentals\fundamentals-go\10\demos\03_panic_vs_error\main.go
* a function that returns an error
```go
func divide1(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor: must not be zero")
	}
	return l / r, nil

}
* we can still turn panics to errors
```go
func divide3(l, r int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}

```
## Summary

# Concurrent Programming

## Introduction

## Concept: Concurrency
![1686331277940](image/README/1686331277940.png)
## Concept: CSP (Communicating Sequential Processes)
* workers, seperate execution work that communicates in a channel
* there is fan input and fan output
## Concept: Goroutines
* function executing concurrently with others in the same address ( same program)
* any fn can be invoked by the gorotinue
* lightweight


## WaitGroups
![1686331738428](image/README/1686331738428.png)
* turn a fn to a gorotinue add the go rontinue in front of it
* use waitgroup to wait until gorontinue s completed
* start to run the gorontinue
```go
	wg.Add(1) //important
	go func() {
		fmt.Println("This happens asynchronously!")
		wg.Done() //when called counter gets back down to 0
	}()

//
wg.Wait() scheduler runs the gorontinue
```


## Demo: Goroutines and WaitGroups
Go_Fundamentals\fundamentals-go\11\demos\01_goroutines_and_waitgroups\main.go
## Demo: Goroutines and WaitGroups

## Demo: Channels
* allow gorotinues to communicate with each other
![1686332305379](image/README/1686332305379.png)
Go_Fundamentals\fundamentals-go\11\demos\02_channels\main.go
Go_Fundamentals\fundamentals-go\11\demos\02_channels\main.go
## Select Statements
![1686333916936](image/README/1686333916936.png)
## Demo: Select Statements
Go_Fundamentals\fundamentals-go\11\demos\03_select_statements\main.go

## Looping with Channels
* close methods, means the channel is closed
* need this to know that all channels are closed and the application can close Go_Fundamentals\fundamentals-go\11\demos\04_ranging_over_channels\main.go
## Demo: Looping with Channels

## Summary

# Testing

## Introduction

## Why Write Tests?
*

## What to Test?
* correctness and performance

## Test Support in Go
* fuzz, test
* ai to generate millions of test cases
* example tests, test are correct and go puts it in your documentation

* benchmark
  *

* profiling
  * tracing - also how were resources used over time

## Demo: Writing a Test
Go_Fundamentals\fundamentals-go\12\demos\main_test.go
* not included in prod builds
* valid test
```go
Test[Xxx]
```
* run test with the terminal
```go
go test ./... // runs all tests in go
```
## Summary
