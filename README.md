# Overview

A library for using primitives when you need to know if the value is the default or if the value is just not set. For example, in Go, there's no way to tell the difference between:

```go
package main

import "fmt"

func main() {
    var notSet string
    intentionallyBlank := ""
    if notSet == intentionallyBlank {
        fmt.Println("values are the same!")
    } else {
        fmt.Println("values are not the same")
    }
}
```

This will emit "values are the same!". But just because a string is blank, doesn't mean it's not an intentional value. You can make all of the major Go-primitive values explicitly set with Go-Optional:

```go
package main

import (
  "fmt"
  "github.com/wojnosystems/go-optional"
)

func main() {
    blankableString := optional.StringUnset()
    blankableString.Set("") // no longer blank
    blankableString.IfSet( func(value string) {
		fmt.Printf("value set! is \"%s\"!\n", value)
    })
    
    blankableString.Unset() // now blank again
    blankableString.IfUnset( func() {
			fmt.Println("value not set")
    })
    
    nonBlankString := optional.StringFrom("I have a value!")
	blankableString.IfSet( func(value string) {
		fmt.Printf("value set! is \"%s\"!\n", value)
	})
    nonBlankString.Unset()
	blankableString.IfUnset( func() {
		fmt.Println("I no longer have a value!")
	})
}
```

This program will return:

```
value set! is ""
value not set
value set! is "I have a value!"
I no longer have a value!
```

# Currently supported types

  * bool: optional.Bool
  * int: optional.Int
  * int64: optional.Int64
  * int32: optional.Int32
  * int16: optional.Int16
  * int8: optional.Int8
  * uint: optional.Uint
  * uint64: optional.Uint64
  * uint32: optional.Uint32
  * uint16: optional.Uint16
  * uint8: optional.Uint8
  * byte: optional.Byte
  * rune: optional.Rune
  * string: optional.String
  * float64: optional.Float64
  * float32: optional.Float32
  * uintptr: optional.Uintptr
  * complex128: optional.Complex128
  * complex64: optional.Complex64

# Provided interfaces

## Typeless

The following interfaces are fulfilled by all provided primitives:

* Tester: exposes IsSet and IfUnset so you can query if the optional has a value
* Unsetter: exposes Unset so you can clear the value of the optional
* TestUnsetter: combines both Tester and Unsetter

## Type-specific

The following interfaces are satisfied by all provided primitives, but is dependent on the type of data stored in the Optional:

* TYPETester: Exposes the IfSet callback methods and provides a type-specific callback to avoid compilation errors.

# Why the IntTester/StringTester specific interfaces?

Ideally, the compiler should catch developer mistakes and prevent the developer from making them. V1 of this module fails at that due to exposing Value(). So this situation should be prevented by the compiler:

```go
package main

import (
  "github.com/wojnosystems/go-optional"
)

func main() {
	blankableString := optional.StringUnset()
	if blankableString.IsSet() {
		// do a thing
    }
    dangerZone := blankableString.Value()
    _ = dangerZone // could be unset!
}
```

The compiler can prevent it by only allowing you to access the variable within a block that only exposes it within that block, like so:

```go
package main

import (
  "github.com/wojnosystems/go-optional"
)

func main() {
	blankableString := optional.StringUnset()
	blankableString.IfSet(func(value string) {
		// do a thing
    })
	// If value is removed, I can't do anything with the value as it's only visibile within the block above
}
```

Now, you can smuggle values out of these blocks, I get that, but it's a sure sign of code smell and should help alleviate accidental mis-use of even the Optionally-guarded values.
