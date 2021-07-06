# Humantouch

[![Go Reference](https://pkg.go.dev/badge/github.com/masv3971/humantouch.svg)](https://pkg.go.dev/github.com/masv3971/humantouch)


## Installation

```
go get -u github.com/masv3971/humantouch
```

## Example
```go
package main

import (
    "github.com/masv3971/humantouch"
)

func main() {
    // Always seed rand!
	rand.Seed(time.Now().Unix())

    person, _ := humantouch.New(nil)

    person.RandomHuman()
    // *Person
}
```