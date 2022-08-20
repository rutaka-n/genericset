# Generic Set

`genericset` provides a simple map-based implementation of generic set.

It uses `sync.RWMutex` to keep consistency of data.

## Installation

```sh
go get github.com/rutaka-n/genericset
```

## Usage

```go
import (
    gs "github.com/rutaka-n/genericset"
)

func main() {
    s1 := gs.New[int]()
    s2 := gs.New[int]()

    s1.Add(1,2,3)
    s2.Add(3,4,5)

    s1.Intersection(&s2) // [3]
    s1.Union(&s2) // [1,2,3,4,5]
    s1.IsEmpty() // false
    s1.Size() // 3
    s1.IsSubset(&s2) // false
    // etc
}
```
