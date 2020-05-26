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

This will emit "values are the same!". However, if you use the go-present library:

```go
package main

import (

"fmt"
"github.com/wojnosystems/go-present"
)

func main() {
    blankableString := present.NewString()
    blankableString.Set("") // no longer blank
    if value, ok := blankableString.ValueWithOK(); ok {
        fmt.Printf("value set! but is \"%s\"!\n", value)
    } else {
        fmt.Println("value not set")
    }
    
    blankableString.Unset() // now blank again
    if value, ok := blankableString.ValueWithOK(); ok {
        fmt.Printf("value set! but is \"%s\"!\n", value)
    } else {
        fmt.Println("value not set")
    }
}
```

This program will return:

```
value set! but is ""
value not set
```
