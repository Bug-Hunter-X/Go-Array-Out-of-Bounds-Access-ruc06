```go
package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println(a)

    for i := 0; i < len(a); i++ {
        a[i] = i
    }

    fmt.Println(a)
}
```