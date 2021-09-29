## marionette

A tiny package for manipulating the default browser cross platform.

### Practical Example

```go
package main

import (
    "fmt"
    "github.com/LQR471814/marionette"
)

func main() {
    browser, err := marionette.DefaultBrowser()
    if err != nil {
        panic(err)
    }

    switch browser {
    case marionette.CHROME, marionette.EDGE:
        marionette.OpenBrowser(`--app=https://google.com`, `--guest`)
    default:
        fmt.Println("The default browser configured on this computer isn't supported!")
    }
}
```

### Constants

```go
const (
    CHROME BrowserType = iota
    EDGE
    FIREFOX
)
```

### Functions

#### func DefaultBrowser

```go
func DefaultBrowser() (BrowserType, error)
```

#### func GetBrowserPath

```go
func GetBrowserPath() (string, error)
```

#### func OpenBrowser

```go
func OpenBrowser(args ...string) error
```

### Types

#### type BrowserType

```go
type BrowserType int
```

#### type UnknownBrowserType

```go
type UnknownBrowserType struct{}
```

#### func (*UnknownBrowserType) Error() string

```go
func (*UnknownBrowserType) Error() string
```

